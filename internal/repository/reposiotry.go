package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/begenov/register-service/internal/domain"
)

//go:generate mockgen -source=reposiotry.go -destination=mocks/mock.go
type Users interface {
	Create(ctx context.Context, user domain.Register) error
	GetByEmail(ctx context.Context, email string) (domain.Register, error)
	SetSession(ctx context.Context, userID int, token string, expiresAt time.Time) error
	GetUserById(ctx context.Context, id int) (domain.Register, error)
	GetUserByRefreshToken(ctx context.Context, token string) (domain.Register, error)
}

type Courier interface {
	Create(ctx context.Context, courier domain.Register) error
	GetCourierByEmail(ctx context.Context, email string) (domain.Register, error)
	SetSession(ctx context.Context, token string, expiredAt time.Time, courierId int) error
	GetCourierByID(ctx context.Context, id int) (domain.Register, error)
	GetCourierByRefreshToken(ctx context.Context, token string) (domain.Register, error)
}

type Restaurant interface {
	Create(ctx context.Context, user domain.Register) error
	GetRestaurantByEmail(ctx context.Context, email string) (domain.Register, error)
	SetSession(ctx context.Context, token string, expiredAt time.Time, restaurantId int) error
	GetRestaurantByID(ctx context.Context, id int) (domain.Register, error)
	GetRestaurantByRefreshToken(ctx context.Context, token string) (domain.Register, error)
}

type Repository struct {
	Users      Users
	Courier    Courier
	Restaurant Restaurant
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Users:      NewUserRepo(db),
		Courier:    NewCourierRepo(db),
		Restaurant: NewRestaurantRepo(db),
	}
}
