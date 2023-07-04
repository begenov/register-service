package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/begenov/register-service/internal/domain"
	"github.com/begenov/register-service/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestCourierRepo_Create(t *testing.T) {

	courier := getCourier()
	ctx := context.Background()
	repo := CourierRepo{db: db}
	err := repo.Create(ctx, courier)
	require.NoError(t, err)
}

func getCourier() domain.Register {
	return domain.Register{
		Email:    util.RandomEmail(),
		Phone:    util.RandomString(8),
		Password: util.RandomString(5),
	}
}

func TestCourierRepo_GetCourierByEmail(t *testing.T) {
	courier := getCourier()
	ctx := context.Background()

	repo := CourierRepo{
		db: db,
	}

	err := repo.Create(ctx, courier)
	require.NoError(t, err)

	c, err := repo.GetCourierByEmail(ctx, courier.Email)
	require.NoError(t, err)
	require.NotEmpty(t, courier)

	require.Equal(t, courier.Email, c.Email)
	require.Equal(t, courier.Phone, c.Phone)
	require.Equal(t, courier.Password, c.Password)
}

func TestCourierRepo_SetSession(t *testing.T) {
	ctx := context.Background()
	id := 1
	expiresAt := time.Now()
	refreshToken := util.RandomString(10)

	repo := CourierRepo{db: db}
	err := repo.SetSession(ctx, refreshToken, expiresAt, id)
	require.NoError(t, err)
}

func TestCourierRepo_GetCourierByID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CourierRepo{
				db: tt.fields.db,
			}
			got, err := r.GetCourierByID(tt.args.ctx, tt.args.id)
			require.NoError(t, err)
			require.NotEmpty(t, got)
		})
	}
}
