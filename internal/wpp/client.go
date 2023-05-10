package wpp

import (
	"context"
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/configs"
	_ "github.com/lib/pq"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
	"os"

	//_ "github.com/glebarez/go-sqlite"
	//_ "github.com/jackc/pgx/v5/pgxpool"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"

	qrcode "github.com/skip2/go-qrcode"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

type WhatsAppClient interface {
	Start() error
	Stop()
	SendText(user *model.User, msg string) error
	SendImage(user *model.User) error
}

type whatsAppClient struct {
	repo repository.WhatsAppRepository

	container *sqlstore.Container
	client    *whatsmeow.Client
	device    *store.Device

	whatsapp *model.WhatsApp
	ctx      context.Context
}

func New(repo repository.WhatsAppRepository) (WhatsAppClient, error) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	wc := configs.Get().WhatsApp
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", wc.Username, wc.Password, wc.Host, wc.Port, wc.Name)
	container, err := sqlstore.New("postgres", url, dbLog)
	if err != nil {
		return nil, err
	}

	device, err := container.GetFirstDevice()
	if err != nil {
		return nil, err
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(device, clientLog)
	instance := &whatsAppClient{
		repo:      repo,
		container: container,
		client:    client,
		device:    device,
		ctx:       context.Background(),
		whatsapp:  &model.WhatsApp{},
	}
	instance.client.AddEventHandler(instance.eventHandler)
	return instance, nil
}

func (c *whatsAppClient) qrCodeRoutine(ctx context.Context, qrChan <-chan whatsmeow.QRChannelItem) {
	var previousCode string
	for evt := range qrChan {
		var err error
		if evt.Event != "code" {
			break
		}
		newQRCode := evt.Code
		if previousCode == newQRCode {
			continue
		}
		if previousCode != "" {
			err = c.repo.DeleteByQR(ctx, previousCode)
			if err != nil {
				// TODO: log error
			}
		}

		c.whatsapp, err = c.repo.CreateFromQR(ctx, newQRCode)
		if err != nil {
			// TODO: log error
		}
		qrterminal.GenerateHalfBlock(evt.Code, qrterminal.H, os.Stdout)
		//qrterminal.GenerateHalfBlock(evt.Code, qrterminal.H, os.Stdout)
		previousCode = newQRCode
	}
}

func (c *whatsAppClient) Start() error {
	var qrChan <-chan whatsmeow.QRChannelItem
	if c.client.Store.ID != nil {
		err := c.client.Connect()
		if err != nil {
			return err
		}
		return nil
	}
	qrChan, _ = c.client.GetQRChannel(c.ctx)
	err := c.client.Connect()
	if err != nil {
		return err
	}
	go c.qrCodeRoutine(c.ctx, qrChan)
	return nil
}

func (c *whatsAppClient) eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.PairError:
		c.repo.Delete(c.ctx, c.whatsapp)

	case *events.Connected:
		c.whatsapp.Phone = c.device.ID.User
		c.whatsapp.Active = true
		c.repo.Update(c.ctx, c.whatsapp)

	case *events.Disconnected:
		c.repo.Delete(c.ctx, c.whatsapp)
		c.restart()

	case *events.TemporaryBan:
		c.whatsapp.Banned = true
		c.whatsapp.Active = false
		c.repo.Update(c.ctx, c.whatsapp)
		c.restart()

	case *events.LoggedOut:
		c.repo.Delete(c.ctx, c.whatsapp)
		c.restart()

	default:
		fmt.Println("Unknown event!", v)
	}
}

func (c *whatsAppClient) restart() {
	c.client.Disconnect()
	err := c.Start()
	if err != nil {
		// TODO: log error
	}
}

func (c *whatsAppClient) Stop() {
	c.client.Disconnect()
}

func (c *whatsAppClient) SendText(user *model.User, msg string) error {

	to := types.NewJID(user.Phone, types.DefaultUserServer)
	message := &waProto.Message{
		Conversation: proto.String(msg),
	}

	_, err := c.client.SendMessage(c.ctx, to, message)
	if err != nil {
		return err
	}
	return nil
}

func GenerateQR(data string) ([]byte, error) {
	var qrBytes []byte
	qrBytes, err := qrcode.Encode(data, qrcode.Highest, 256)
	if err != nil {
		return nil, err
	}
	return qrBytes, nil
}

func (c *whatsAppClient) SendImage(user *model.User) error {
	qrBytes, err := GenerateQR(user.UUID)
	if err != nil {
		return err
	}
	res, err := c.client.Upload(context.Background(), qrBytes, whatsmeow.MediaImage)
	if err != nil {
		return err
	}
	imageMsg := &waProto.ImageMessage{
		Caption:       proto.String("Hello, world!"),
		Mimetype:      proto.String("image/png"),
		Url:           &res.URL,
		DirectPath:    &res.DirectPath,
		MediaKey:      res.MediaKey,
		FileEncSha256: res.FileEncSHA256,
		FileSha256:    res.FileSHA256,
		FileLength:    &res.FileLength,
	}

	to := types.NewJID(user.Phone, types.DefaultUserServer)
	msg := &waProto.Message{
		ImageMessage: imageMsg,
	}

	_, err = c.client.SendMessage(context.Background(), to, msg)
	if err != nil {
		return err
	}
	return nil
}
