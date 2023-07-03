package postgres

import (
	"context"
	"database/sql"
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

	return nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (domain.Register, error) {
	return domain.Register{}, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, email string, password string) (domain.Register, error) {
	return domain.Register{}, nil
}

func (r *UserRepo) SetSession(ctx context.Context, userID int, token string, expiresAt time.Time) error {
	return nil
}
