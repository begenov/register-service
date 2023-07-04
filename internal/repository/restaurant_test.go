package repository

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/begenov/register-service/internal/domain"
	"github.com/begenov/register-service/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestRestaurantRepo_Create(t *testing.T) {
	restaurant := getRestaurant()
	repo := RestaurantRepo{
		db: db,
	}
	err := repo.Create(context.Background(), restaurant)
	require.NoError(t, err)
}

func TestRestaurantRepo_GetRestaurantByEmail(t *testing.T) {
	restaurant := getRestaurant()
	repo := RestaurantRepo{
		db: db,
	}
	err := repo.Create(context.Background(), restaurant)
	require.NoError(t, err)

	r, err := repo.GetRestaurantByEmail(context.Background(), restaurant.Email)
	require.NoError(t, err)
	require.NotEmpty(t, r)

	require.Equal(t, restaurant.Email, r.Email)
	require.Equal(t, restaurant.Phone, r.Phone)
	require.Equal(t, restaurant.Password, r.Password)
	require.Equal(t, restaurant.Address, r.Address)
}

func TestRestaurantRepo_SetSession(t *testing.T) {

	repo := RestaurantRepo{
		db: db,
	}
	token := util.RandomString(10)
	expiresAt := time.Now()
	id := 1
	err := repo.SetSession(context.Background(), token, expiresAt, id)
	require.NoError(t, err)
}

func TestRestaurantRepo_GetRestaurantByID(t *testing.T) {
	repo := RestaurantRepo{
		db: db,
	}
	id := 1

	restaurant, err := repo.GetRestaurantByID(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, restaurant)
}

func TestNewRestaurantRepo(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *RestaurantRepo
	}{
		{
			name: "ok",
			args: args{
				db: db,
			},
			want: &RestaurantRepo{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRestaurantRepo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRestaurantRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestaurantRepo_GetRestaurantByRefreshToken(t *testing.T) {
	refresh_token := util.RandomString(10)
	expiresAt := time.Now()
	id := 1
	repo := RestaurantRepo{
		db: db,
	}

	err := repo.SetSession(context.Background(), refresh_token, expiresAt, id)
	require.NoError(t, err)

	r, err := repo.GetRestaurantByRefreshToken(context.Background(), refresh_token)
	require.NoError(t, err)
	require.NotEmpty(t, r)

	require.Equal(t, r.ID, id)
}

func getRestaurant() domain.Register {
	return domain.Register{
		Email:    util.RandomEmail(),
		Phone:    util.RandomString(8),
		Address:  util.RandomString(6),
		Password: util.RandomString(8),
	}
}
