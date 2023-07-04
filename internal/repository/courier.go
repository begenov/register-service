package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/begenov/register-service/internal/domain"
)

type CourierRepo struct {
	db *sql.DB
}

func NewCourierRepo(db *sql.DB) *CourierRepo {
	return &CourierRepo{
		db: db,
	}
}

func (r *CourierRepo) Create(ctx context.Context, courier domain.Register) error {

	stmt := `INSERT INTO "courier" (email, phone, password_hash) VALUES ($1, $2, $3)`
	if _, err := r.db.ExecContext(ctx, stmt, courier.Email, courier.Phone, courier.Password); err != nil {
		return err
	}

	return nil
}

func (r *CourierRepo) GetCourierByEmail(ctx context.Context, email string) (domain.Register, error) {
	stmt := `SELECT email, phone, password_hash FROM courier WHERE email = $1`
	row := r.db.QueryRowContext(ctx, stmt, email)
	var courier domain.Register
	if err := row.Scan(&courier.Email, &courier.Phone, &courier.Password); err != nil {
		return courier, err
	}

	return courier, nil
}

func (r *CourierRepo) SetSession(ctx context.Context, token string, expiredAt time.Time, courierId int) error {
	stmt := `UPDATE courier SET refresh_token = $1, expired_at = $2 WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, stmt, token, expiredAt, courierId); err != nil {
		return err
	}

	return nil
}

func (r *CourierRepo) GetCourierByID(ctx context.Context, id int) (domain.Register, error) {
	stmt := `SELECT email, phone, password_hash FROM courier WHERE id = $1`
	row := r.db.QueryRowContext(ctx, stmt, id)
	var courier domain.Register
	if err := row.Scan(&courier.Email, &courier.Phone, &courier.Password); err != nil {
		return courier, err
	}

	return courier, nil
}

func (r *CourierRepo) GetCourierByRefreshToken(ctx context.Context, token string) (domain.Register, error) {
	stmt := `SELECT id, email, phone, password_hash, refresh_token, expired_at FROM courier WHERE refresh_token = $1`
	row := r.db.QueryRowContext(ctx, stmt, token)

	var user domain.Register
	if err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.Password, &user.RefreshToken, &user.ExpiresAt); err != nil {
		return user, err
	}

	return user, nil
}
