package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/begenov/register-service/internal/domain"
)

type Users interface {
	Create(ctx context.Context, user domain.Register) error
	GetByEmail(ctx context.Context, email string) (domain.Register, error)
	UpdateUser(ctx context.Context, email string, password string) (domain.Register, error)
	SetSession(ctx context.Context, userID int, token string, expiresAt time.Time) error
}

type Repository struct {
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
