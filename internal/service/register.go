package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/begenov/register-service/internal/domain"
	"github.com/begenov/register-service/internal/repository"
	"github.com/begenov/register-service/pkg/auth"
	"github.com/begenov/register-service/pkg/hash"
)

type RegisterService struct {
	repo            *repository.Repository
	manager         auth.TokenManager
	hash            hash.PasswordHasher
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewRegisterService(repo *repository.Repository, hash hash.PasswordHasher, auth auth.TokenManager, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *RegisterService {
	return &RegisterService{
		repo:            repo,
		manager:         auth,
		hash:            hash,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *RegisterService) SignUp(ctx context.Context, register domain.Register) error {
	var err error
	register.Password, err = s.hash.GenerateFromPassword(register.Password)
	if err != nil {
		return err
	}

	switch register.Role {
	case domain.User:
		return s.createUser(ctx, register)
	case domain.Courier:
		return s.createCourier(ctx, register)
	case domain.Restaurant:
		return s.createRestaurant(ctx, register)
	default:
		return errors.New("error role empty")
	}
}

func (s *RegisterService) SignIn(ctx context.Context, req domain.RequestSignIn) (domain.Token, error) {
	user, err := s.getByEmailRole(ctx, req.Email, req.Role)
	if err != nil {
		return domain.Token{}, err
	}
	// if err = s.hash.CompareHashAndPassword(user.Password, req.Password); err != nil {
	// 	return domain.Token{}, err
	// }

	return s.createSession(ctx, user.ID, req.Role)
}

func (s *RegisterService) GetUserById(ctx context.Context, id int) (domain.Register, error) {
	return s.repo.Users.GetUserById(ctx, id)
}

func (s *RegisterService) GetUserByRefreshToken(ctx context.Context, token string, role string) (domain.Token, error) {
	user, err := s.getByRefreshTokenRole(ctx, token, role)
	if err != nil {
		return domain.Token{}, err
	}

	return s.createSession(ctx, user.ID, role)
}
func (s *RegisterService) RefreshToken(ctx context.Context, refreshToken string, role string) (domain.Token, error) {

	user, err := s.getByRefreshTokenRole(ctx, refreshToken, role)
	if err != nil {
		return domain.Token{}, err
	}
	if time.Now().Before(user.ExpiresAt) {
		return domain.Token{}, errors.New("error expires at")
	}

	return s.createSession(ctx, user.ID, role)
}

func (s *RegisterService) createUser(ctx context.Context, register domain.Register) error {
	err := s.repo.Users.Create(ctx, register)
	return err
}

func (s *RegisterService) createCourier(ctx context.Context, register domain.Register) error {
	return s.repo.Courier.Create(ctx, register)
}

func (s *RegisterService) createRestaurant(ctx context.Context, register domain.Register) error {
	return s.repo.Restaurant.Create(ctx, register)
}

func (s *RegisterService) createSession(ctx context.Context, id int, role string) (domain.Token, error) {
	var (
		res domain.Token
		err error
	)
	res.AccessToken, err = s.manager.NewJWT(strconv.Itoa(id), s.accessTokenTTL)
	if err != nil {
		return res, err
	}
	res.RefreshToken, err = s.manager.NewRefreshToken()
	if err != nil {
		return res, err
	}
	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}
	err = s.setSessionRole(ctx, id, res.RefreshToken, session.ExpiresAt, role)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *RegisterService) setSessionRole(ctx context.Context, id int, token string, expiresAt time.Time, role string) error {
	switch role {
	case domain.User:
		return s.repo.Users.SetSession(ctx, id, token, expiresAt)
	case domain.Courier:
		return s.repo.Courier.SetSession(ctx, token, expiresAt, id)
	case domain.Restaurant:
		return s.repo.Restaurant.SetSession(ctx, token, expiresAt, id)
	default:
		return errors.New("error role empty")
	}
}

func (s *RegisterService) getByEmailRole(ctx context.Context, email string, role string) (domain.Register, error) {
	switch role {
	case domain.User:
		return s.repo.Users.GetByEmail(ctx, email)
	case domain.Courier:
		return s.repo.Courier.GetCourierByEmail(ctx, email)
	case domain.Restaurant:
		return s.repo.Restaurant.GetRestaurantByEmail(ctx, email)
	default:
		return domain.Register{}, errors.New("error role empty")
	}
}

func (s *RegisterService) getByRefreshTokenRole(ctx context.Context, token string, role string) (domain.Register, error) {
	switch role {
	case domain.User:
		return s.repo.Users.GetUserByRefreshToken(ctx, token)
	case domain.Courier:
		return s.repo.Courier.GetCourierByRefreshToken(ctx, token)
	case domain.Restaurant:
		return s.repo.Restaurant.GetRestaurantByRefreshToken(ctx, token)
	default:
		return domain.Register{}, errors.New("error role empty")
	}
}
