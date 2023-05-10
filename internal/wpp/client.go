package wpp

import (
	"context"
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/configs"
	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"

	//_ "github.com/glebarez/go-sqlite"
	//_ "github.com/jackc/pgx/v5/pgxpool"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"

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
	//dbLog := waLog.Stdout("Database", "DEBUG", true)
	wc := configs.Get().WhatsApp
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

	return &whatsAppClient{
		repo:      repo,
		container: container,
		client:    client,
		device:    device,
		ctx:       context.Background(),
	}, nil
}

func (c *whatsAppClient) qrCodeRoutine(ctx context.Context, qrChan <-chan whatsmeow.QRChannelItem) {
	var previousCode string
	for evt := range qrChan {
		if evt.Event != "code" {
			break
		}
		newQRCode := evt.Code
		if previousCode == newQRCode {
			continue
		}
		if previousCode != "" {
			c.repo.DeleteByQR(ctx, previousCode)
		}
		c.repo.CreateFromQR(ctx, newQRCode)
		fmt.Println(newQRCode)
		previousCode = newQRCode
	}
}

func (c *whatsAppClient) Start() error {
	c.client.AddEventHandler(c.eventHandler)
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
	// TODO: Implement
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

func (c *whatsAppClient) SendImage(user *model.User) error {
	var qrCodeBytes []byte
	res, err := c.client.Upload(context.Background(), qrCodeBytes, whatsmeow.MediaImage)
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
