package service

import (
	"context"
	"time"

	"github.com/begenov/register-service/internal/domain"
	"github.com/begenov/register-service/internal/repository"
	"github.com/begenov/register-service/pkg/auth"
	"github.com/begenov/register-service/pkg/hash"
)

type Register interface {
	SignUp(ctx context.Context, register domain.Register) error
	SignIn(ctx context.Context, req domain.RequestSignIn) (domain.Token, error)
	RefreshToken(ctx context.Context, refreshToken string, role string) (domain.Token, error)
}

type Service struct {
	Register Register
}

func NewService(repo *repository.Repository, hash hash.PasswordHasher, auth auth.TokenManager, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *Service {
	return &Service{
		Register: NewRegisterService(repo, hash, auth, accessTokenTTL, refreshTokenTTL),
	}
}
