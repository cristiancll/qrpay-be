package repository

import (
	"context"
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WhatsAppRepository interface {
	Migrater
	Updater[model.WhatsApp]
	Deleter[model.WhatsApp]
	TCreater[model.WhatsApp]
	TGetterByUUID[model.WhatsApp]
	TGetterAll[model.WhatsApp]
	DeleteByQR(ctx context.Context, code string) error
	CreateFromQR(ctx context.Context, code string) (*model.WhatsApp, error)
}

func NewWhatsAppRepository(db *pgxpool.Pool) WhatsAppRepository {
	return &whatsAppRepository{db: db}
}

const (
	createWhatsappTableQuery = `CREATE TABLE IF NOT EXISTS whatsapp (
									id SERIAL PRIMARY KEY, 
									uuid VARCHAR(255) NOT NULL, 
									qr VARCHAR(255) NOT NULL, 
									phone VARCHAR(255), 
									active BOOLEAN DEFAULT FALSE, 
									banned BOOLEAN DEFAULT FALSE, 
									created_at TIMESTAMP NOT NULL, 
									updated_at TIMESTAMP NOT NULL
								)`

	createWhatsappQuery = `INSERT INTO whatsapp (uuid, qr, created_at, updated_at) VALUES ($1, $2, now(), now()) RETURNING id, created_at, updated_at`

	updateWhatsappQuery  = `UPDATE whatsapp SET qr = $2, phone = $3, active = $4, banned = $5, updated_at = now() WHERE id = $1`
	toggleActiveByQRCode = `UPDATE whatsapp SET active = $2, updated_at = now() WHERE qr = $1`
	toggleActiveByPhone  = `UPDATE whatsapp SET active = $2, updated_at = now() WHERE phone = $1`
	toggleBannedByQRCode = `UPDATE whatsapp SET banned = $2, updated_at = now() WHERE qr = $1`
	toggleBannedByPhone  = `UPDATE whatsapp SET banned = $2, updated_at = now() WHERE phone = $1`

	deleteWhatsappQuery         = `DELETE FROM whatsapp WHERE id = $1`
	deleteWhatsappByUUIDQuery   = `DELETE FROM whatsapp WHERE uuid = $1`
	deleteWhatsappByQRCodeQuery = `DELETE FROM whatsapp WHERE qr = $1`
	deleteWhatsappByPhoneQuery  = `DELETE FROM whatsapp WHERE phone = $1`

	getWhatsappByIDQuery     = `SELECT id, uuid, qr, phone, active, banned, created_at, updated_at FROM whatsapp WHERE id = $1`
	getWhatsappByUUIDQuery   = `SELECT id, uuid, qr, phone, active, banned, created_at, updated_at FROM whatsapp WHERE uuid = $1`
	getWhatsappByQRCodeQuery = `SELECT id, uuid, qr, phone, active, banned, created_at, updated_at FROM whatsapp WHERE qr = $1`
	getWhatsappByPhoneQuery  = `SELECT id, uuid, qr, phone, active, banned, created_at, updated_at FROM whatsapp WHERE phone = $1`
	getAllWhatsappQuery      = `SELECT id, uuid, qr, phone, active, banned, created_at, updated_at FROM whatsapp`

	countWhatsappByQRCodeQuery = `SELECT COUNT(*) FROM whatsapp WHERE qr = $1`
)

type whatsAppRepository struct {
	db *pgxpool.Pool
}

func (r *whatsAppRepository) TCreate(ctx context.Context, tx pgx.Tx, whats *model.WhatsApp) error {
	whats.UUID = uuid.New().String()
	row := tx.QueryRow(ctx, createWhatsappQuery, whats.UUID, whats.QR)
	err := row.Scan(&whats.ID, &whats.CreatedAt, &whats.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error scanning whatsapp: %w", err)
	}
	return nil
}

func (r *whatsAppRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.WhatsApp, error) {
	row := tx.QueryRow(ctx, getWhatsappByUUIDQuery, uuid)
	whats := &model.WhatsApp{}
	err := row.Scan(&whats.ID, &whats.UUID, &whats.QR, &whats.Phone, &whats.Active, &whats.Banned, &whats.CreatedAt, &whats.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error scanning whatsapp: %w", err)
	}
	return whats, nil
}

func (r *whatsAppRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.WhatsApp, error) {
	rows, err := tx.Query(ctx, getAllWhatsappQuery)
	if err != nil {
		return nil, fmt.Errorf("error getting all whatsapp: %w", err)
	}
	defer rows.Close()

	whatsList := make([]*model.WhatsApp, 0)
	for rows.Next() {
		w := &model.WhatsApp{}
		err = rows.Scan(&w.ID, &w.UUID, &w.QR, &w.Phone, &w.Active, &w.Banned, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning whatsapp: %w", err)
		}
		whatsList = append(whatsList, w)
	}
	return whatsList, nil
}

func (r *whatsAppRepository) Update(ctx context.Context, whats *model.WhatsApp) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, updateWhatsappQuery, whats.ID, whats.QR, whats.Phone, whats.Active, whats.Banned)
	if err != nil {
		return fmt.Errorf("error updating whatsapp: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return status.Error(codes.Unimplemented, "method GetAll not implemented")
}

func (r *whatsAppRepository) Delete(ctx context.Context, whats *model.WhatsApp) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, deleteWhatsappQuery, whats.ID)
	if err != nil {
		return fmt.Errorf("error deleting whatsapp: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return status.Error(codes.Unimplemented, "method GetAll not implemented")
}

func (r *whatsAppRepository) DeleteByQR(ctx context.Context, qrCode string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, deleteWhatsappByQRCodeQuery, qrCode)
	if err != nil {
		return fmt.Errorf("error deleting whatsapp by qr code: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return status.Error(codes.Unimplemented, "method GetAll not implemented")
}

func (r *whatsAppRepository) CreateFromQR(ctx context.Context, qrCode string) (*model.WhatsApp, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	count := 0
	row := tx.QueryRow(ctx, countWhatsappByQRCodeQuery, qrCode)
	err = row.Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error getting whatsapp by qr code: %w", err)
	}
	if count > 0 {
		return nil, fmt.Errorf("error whatsapp already exists: %w", err)
	}

	whats := &model.WhatsApp{
		QR: qrCode,
	}

	err = r.TCreate(ctx, tx, whats)
	if err != nil {
		return nil, fmt.Errorf("error creating whatsapp: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}
	return whats, nil
}

func (r *whatsAppRepository) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createWhatsappTableQuery)
	if err != nil {
		return fmt.Errorf("error migrating whatsapp table: %w", err)
	}
	return nil
}
