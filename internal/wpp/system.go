package wpp

import (
	"context"
	"fmt"
	"github.com/cristiancll/qrpay-be/configs"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/common"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"

	"github.com/skip2/go-qrcode"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

type WhatsAppSystem interface {
	Start() error
	Stop()
	SendText(user *model.User, msg string)
	SendImage(user *model.User, msg string)
}

type whatsAppSystem struct {
	db *pgxpool.Pool

	repo     repository.WhatsApp
	authRepo repository.Auth
	userRepo repository.User

	container *sqlstore.Container
	client    *whatsmeow.Client
	device    *store.Device

	whatsapp *model.WhatsApp
	ctx      context.Context

	verifiedCache []string
}

func New(db *pgxpool.Pool, repo repository.WhatsApp, userRepo repository.User, authRepo repository.Auth, verifiedCache []string) (WhatsAppSystem, error) {
	//dbLog := waLog.Stdout("Database", "DEBUG", true)
	wc := configs.Get().Database
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", wc.Username, wc.Password, wc.Host, wc.Port, wc.Name)
	container, err := sqlstore.New("postgres", url, nil)
	if err != nil {
		return nil, err
	}

	device, err := container.GetFirstDevice()
	if err != nil {
		return nil, err
	}
	//clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(device, nil)
	instance := &whatsAppSystem{
		db:            db,
		repo:          repo,
		container:     container,
		client:        client,
		device:        device,
		ctx:           context.Background(),
		whatsapp:      &model.WhatsApp{},
		verifiedCache: verifiedCache,
		userRepo:      userRepo,
		authRepo:      authRepo,
	}
	instance.client.AddEventHandler(instance.eventHandler)
	return instance, nil
}

func (s *whatsAppSystem) qrCodeRoutine(ctx context.Context, qrChan <-chan whatsmeow.QRChannelItem) {
	var previousCode string
	defer s.repo.ClearUnusedWhatsApp(ctx)
	timeout := false
	for evt := range qrChan {
		var err error
		if evt.Event != "code" {
			if evt.Event == "timeout" {
				timeout = true
			}
			break
		}
		newQRCode := evt.Code
		if previousCode == newQRCode {
			continue
		}
		if previousCode != "" {
			err = s.repo.DeleteByQR(ctx, previousCode)
			if err != nil {
				// TODO: log error
			}
		}

		s.whatsapp, err = s.repo.CreateFromQR(ctx, newQRCode)
		if err != nil {
			// TODO: log error
		}
		previousCode = newQRCode
	}
	if timeout {
		err := s.Start()
		if err != nil {
			// TODO: log error
		}
	}
}

func (s *whatsAppSystem) Start() error {
	err := s.repo.ClearUnusedWhatsApp(s.ctx)
	if err != nil {
		// TODO: log error
	}
	var qrChan <-chan whatsmeow.QRChannelItem
	if s.client.Store.ID != nil {
		err := s.client.Connect()
		if err != nil {
			return err
		}
		whats, err := s.repo.GetByPhone(s.ctx, s.client.Store.ID.User)
		if err != nil {
			return err
		}
		s.whatsapp = whats
		return nil
	}
	qrChan, _ = s.client.GetQRChannel(s.ctx)
	err = s.client.Connect()
	if err != nil {
		return err
	}
	go s.qrCodeRoutine(s.ctx, qrChan)
	return nil
}

func (s *whatsAppSystem) eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.PairSuccess:
		s.whatsapp.Scanned = true
		err := s.repo.Update(s.ctx, s.whatsapp)
		if err != nil {
			// TODO: log error
		}
	case *events.PairError:
		err := s.repo.Delete(s.ctx, s.whatsapp)
		if err != nil {
			// TODO: log error
		}
	case *events.Connected:
		s.whatsapp.Phone = &s.device.ID.User
		s.whatsapp.Active = true
		err := s.repo.Update(s.ctx, s.whatsapp)
		if err != nil {
			// TODO: log error
		}
	case *events.Disconnected:
		err := s.repo.Delete(s.ctx, s.whatsapp)
		if err != nil {
			// TODO: log error
		}
		s.restart()
	case *events.TemporaryBan:
		s.whatsapp.Banned = true
		s.whatsapp.Active = false
		err := s.repo.Update(s.ctx, s.whatsapp)
		if err != nil {
			// TODO: log error
		}
		s.restart()
	case *events.LoggedOut:
		err := s.repo.Delete(s.ctx, s.whatsapp)
		if err != nil {
			// TODO: log error
		}
		s.restart()
	case *events.Message:
		phone := v.Info.MessageSource.Sender.User
		msg := *v.Message.Conversation
		valid := []string{"OK", "K", "CONFIRM", "CONFIRMA", "CONFIRMO", "SIM", "S", "BLZ", "YES", "Y", "CERTO"}
		for _, v := range valid {
			if strings.ToUpper(msg) == v {
				s.handleUserVerification(phone)
				continue
			}
		}
	}
}

func (s *whatsAppSystem) handleUserVerification(phone string) {
	for _, v := range s.verifiedCache {
		if phone == v {
			return
		}
	}
	tx, err := s.db.Begin(s.ctx)
	if err != nil {
		// TODO: log error
		return
	}
	defer tx.Rollback(s.ctx)

	user, err := s.userRepo.GetUserByPhone(s.ctx, tx, phone)
	if err != nil {
		// TODO: log error
		return
	}

	err = s.authRepo.VerifyUser(s.ctx, tx, user)
	if err != nil {
		// TODO : log error
		return
	}

	err = tx.Commit(s.ctx)
	if err != nil {
		// TODO : log error
		return
	}

	s.verifiedCache = append(s.verifiedCache, phone)
	go s.SendText(user, user.ConfirmMessage())
}

func (s *whatsAppSystem) restart() {
	s.Stop()
	s.client = whatsmeow.NewClient(s.device, nil)
	s.client.AddEventHandler(s.eventHandler)
	err := s.Start()
	if err != nil {
		// TODO: log error
	}
}

func (s *whatsAppSystem) Stop() {
	err := s.repo.DisableAll(s.ctx)
	if err != nil {
		// TODO: log error
	}
	s.client.Disconnect()
}

func (s *whatsAppSystem) SendText(user *model.User, msg string) {
	sanitizedPhone := common.SanitizePhone(user.Phone)
	to := types.NewJID(sanitizedPhone, types.DefaultUserServer)
	message := &waProto.Message{
		Conversation: proto.String(msg),
	}

	_, err := s.client.SendMessage(s.ctx, to, message)
	if err != nil {
		// TODO: log error
	}
}

func GenerateQR(data string) ([]byte, error) {
	var qrBytes []byte
	qrBytes, err := qrcode.Encode(data, qrcode.Highest, 256)
	if err != nil {
		return nil, err
	}
	return qrBytes, nil
}

func (s *whatsAppSystem) SendImage(user *model.User, msg string) {
	qrBytes, err := GenerateQR(user.UUID)
	if err != nil {
		// TODO: log error
	}
	res, err := s.client.Upload(context.Background(), qrBytes, whatsmeow.MediaImage)
	if err != nil {
		// TODO: log error
	}
	imageMsg := &waProto.ImageMessage{
		Caption:       proto.String(msg),
		Mimetype:      proto.String("image/png"),
		Url:           &res.URL,
		DirectPath:    &res.DirectPath,
		MediaKey:      res.MediaKey,
		FileEncSha256: res.FileEncSHA256,
		FileSha256:    res.FileSHA256,
		FileLength:    &res.FileLength,
	}
	sanitizedPhone := common.SanitizePhone(user.Phone)
	to := types.NewJID(sanitizedPhone, types.DefaultUserServer)
	message := &waProto.Message{
		ImageMessage: imageMsg,
	}
	_, err = s.client.SendMessage(context.Background(), to, message)
	if err != nil {
		// TODO: log error
	}
}
