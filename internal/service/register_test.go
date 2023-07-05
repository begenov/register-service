package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/begenov/register-service/internal/domain"
	"github.com/begenov/register-service/internal/repository"
	mocks_repo "github.com/begenov/register-service/internal/repository/mocks"
	"github.com/begenov/register-service/pkg/auth"
	"github.com/begenov/register-service/pkg/hash"
	"github.com/begenov/register-service/pkg/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var h hash.PasswordHasher

type eqCreateUserParamsMatcher struct {
	arg      string
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(string)
	if !ok {
		return false
	}
	err := h.CompareHashAndPassword(arg, e.password)
	if err != nil {
		return false
	}
	e.arg = arg
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg string, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg: arg, password: password}
}

func TestRegisterService_SignUp(t *testing.T) {

	type args struct {
		ctx      context.Context
		register domain.Register
	}
	tests := []struct {
		name          string
		buildStubs    func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant)
		args          args
		checkResponse func(t *testing.T, register *RegisterService, args domain.Register)
	}{
		{
			name: "USER OK",
			args: args{
				ctx: context.Background(),
				register: domain.Register{
					Role:     "user",
					Email:    util.RandomEmail(),
					Phone:    util.RandomString(7),
					Address:  util.RandomString(6),
					Password: util.RandomString(6),
				},
			},
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
				user.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, register *RegisterService, args domain.Register) {
				err := register.SignUp(context.Background(), args)
				require.NoError(t, err)
			},
		},
		{
			name: "COURIER OK",
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
				courier.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				register: domain.Register{
					Role:     "courier",
					Email:    util.RandomEmail(),
					Phone:    util.RandomString(7),
					Address:  util.RandomString(6),
					Password: util.RandomString(6),
				},
			},
			checkResponse: func(t *testing.T, register *RegisterService, args domain.Register) {
				err := register.SignUp(context.Background(), args)
				require.NoError(t, err)
			},
		},
		{
			name: "RESTAURNAT OK",
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
				restaurant.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				register: domain.Register{
					Role:     "restaurant",
					Email:    util.RandomEmail(),
					Phone:    util.RandomString(7),
					Address:  util.RandomString(6),
					Password: util.RandomString(6),
				},
			},
			checkResponse: func(t *testing.T, register *RegisterService, args domain.Register) {
				err := register.SignUp(context.Background(), args)
				require.NoError(t, err)
			},
		},
		{
			name: "role emtpy",
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
			},
			args: args{
				ctx: context.Background(),
				register: domain.Register{
					Role:     "",
					Email:    util.RandomEmail(),
					Phone:    util.RandomString(7),
					Address:  util.RandomString(6),
					Password: util.RandomString(6),
				},
			},
			checkResponse: func(t *testing.T, register *RegisterService, args domain.Register) {
				err := register.SignUp(context.Background(), args)
				require.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userRepo := mocks_repo.NewMockUsers(ctrl)
			courierRepo := mocks_repo.NewMockCourier(ctrl)
			restaurantRepo := mocks_repo.NewMockRestaurant(ctrl)

			signInKey := "qwerty"

			auth, err := auth.NewManager(signInKey)
			require.NoError(t, err)
			require.NotEmpty(t, auth)

			h = hash.NewHash(10)

			repo := &repository.Repository{
				Users:      userRepo,
				Courier:    courierRepo,
				Restaurant: restaurantRepo,
			}

			accessTokenTTL := 5 * time.Minute
			refreshTokenTTL := 5 * time.Minute

			tt.buildStubs(userRepo, courierRepo, restaurantRepo)
			register := NewRegisterService(repo, h, auth, accessTokenTTL, refreshTokenTTL)
			tt.checkResponse(t, register, tt.args.register)
		})
	}
}

func TestRegisterService_SignIn(t *testing.T) {
	req_user := domain.RequestSignIn{
		Role:     "user",
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
	}
	req_courier := domain.RequestSignIn{
		Role:     "courier",
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
	}
	req_restaurant := domain.RequestSignIn{
		Role:     "restaurant",
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
	}
	req_role_emty := domain.RequestSignIn{
		Role:     "",
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
	}
	type args struct {
		ctx context.Context
		req domain.RequestSignIn
	}
	tests := []struct {
		name          string
		buildStubs    func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant)
		args          args
		checkResponse func(t *testing.T, register *RegisterService, req domain.RequestSignIn)
	}{
		// TODO: Add test cases.
		{
			name: "USER OK",
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
				res := domain.Register{
					Email:    req_user.Email,
					Password: req_user.Password,
					Address:  util.RandomString(6),
					Phone:    util.RandomString(6),
					ID:       int(util.RandomInt(1, 100)),
				}
				user.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Times(1).Return(res, nil)
				user.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				req: req_user,
			},
			checkResponse: func(t *testing.T, register *RegisterService, req domain.RequestSignIn) {
				token, err := register.SignIn(context.Background(), req)
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
		{
			name: "COURIER OK",
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
				res := domain.Register{
					Email:    req_user.Email,
					Password: req_user.Password,
					Phone:    util.RandomString(6),
				}

				courier.EXPECT().GetCourierByEmail(gomock.Any(), gomock.Eq(req_courier.Email)).Times(1).Return(res, nil)
				courier.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				req: req_courier,
			},
			checkResponse: func(t *testing.T, register *RegisterService, req domain.RequestSignIn) {
				token, err := register.SignIn(context.Background(), req)
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
		{
			name: "RESTAURANT OK",
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
				res := domain.Register{
					Email:    req_user.Email,
					Password: req_user.Password,
					Phone:    util.RandomString(6),
					Address:  util.RandomString(6),
				}
				restaurant.EXPECT().GetRestaurantByEmail(gomock.Any(), gomock.Eq(req_restaurant.Email)).Times(1).Return(res, nil)
				restaurant.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				req: req_restaurant,
			},
			checkResponse: func(t *testing.T, register *RegisterService, req domain.RequestSignIn) {
				token, err := register.SignIn(context.Background(), req)
				require.NoError(t, err)
				require.NotEmpty(t, token)
			},
		},
		{
			name: "Error role emty",
			buildStubs: func(user *mocks_repo.MockUsers, courier *mocks_repo.MockCourier, restaurant *mocks_repo.MockRestaurant) {
				// res := domain.Register{
				// 	Email:    req_user.Email,
				// 	Password: req_user.Password,
				// 	Phone:    util.RandomString(6),
				// 	Address:  util.RandomString(6),
				// }
				// restaurant.EXPECT().GetRestaurantByEmail(gomock.Any(), gomock.Eq(req_restaurant.Email)).Times(1).Return(res, nil)
				// restaurant.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				req: req_role_emty,
			},
			checkResponse: func(t *testing.T, register *RegisterService, req domain.RequestSignIn) {
				token, err := register.SignIn(context.Background(), req)
				require.Error(t, err)
				require.Empty(t, token)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userRepo := mocks_repo.NewMockUsers(ctrl)
			courierRepo := mocks_repo.NewMockCourier(ctrl)
			restaurantRepo := mocks_repo.NewMockRestaurant(ctrl)

			signInKey := "qwerty"

			auth, err := auth.NewManager(signInKey)
			require.NoError(t, err)
			require.NotEmpty(t, auth)

			h = hash.NewHash(10)

			repo := &repository.Repository{
				Users:      userRepo,
				Courier:    courierRepo,
				Restaurant: restaurantRepo,
			}

			accessTokenTTL := 5 * time.Minute
			refreshTokenTTL := 5 * time.Minute

			tt.buildStubs(userRepo, courierRepo, restaurantRepo)
			register := NewRegisterService(repo, h, auth, accessTokenTTL, refreshTokenTTL)
			tt.checkResponse(t, register, tt.args.req)
		})
	}
}
