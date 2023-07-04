package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/begenov/register-service/internal/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user domain.Register) error {
	stmt := `INSERT INTO "user" (email, phone, home_address, password_hash) VALUES ($1, $2, $3, $4)`
	if _, err := r.db.ExecContext(ctx, stmt, user.Email, user.Phone, user.Address, user.Password); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (domain.Register, error) {
	stmt := `SELECT id, email, phone, home_address, password_hash FROM "user" WHERE email = $1`
	row := r.db.QueryRowContext(ctx, stmt, email)
	var user domain.Register
	if err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.Address, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetUserById(ctx context.Context, id int) (domain.Register, error) {
	stmt := `SELECT id, email, phone, home_address, password_hash FROM "user" WHERE id = $1`
	row := r.db.QueryRowContext(ctx, stmt, id)
	var user domain.Register
	if err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.Address, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) SetSession(ctx context.Context, userID int, token string, expiredAt time.Time) error {
	stmt := `UPDATE "user" SET refresh_token = $1, expired_at = $2 WHERE id = $3`

	if _, err := r.db.ExecContext(ctx, stmt, token, expiredAt, userID); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *UserRepo) GetUserByRefreshToken(ctx context.Context, token string) (domain.Register, error) {
	stmt := `SELECT id, email, phone, home_address, password_hash, refresh_token, expired_at FROM "user" WHERE refresh_token = $1`
	row := r.db.QueryRowContext(ctx, stmt, token)

	var user domain.Register
	if err := row.Scan(&user.ID, &user.Email, &user.Phone, &user.Address, &user.Password, &user.RefreshToken, &user.ExpiresAt); err != nil {
		return user, err
	}

	return user, nil
}
