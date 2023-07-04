package repository

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/begenov/register-service/internal/domain"
	"github.com/begenov/register-service/pkg/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var db *sql.DB
var err error
var driver = "postgres"
var dsn = "postgresql://root:secret@localhost:5432/register?sslmode=disable"

func init() {
	db, err = sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func TestUserRepo_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx  context.Context
		user domain.Register
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:  context.Background(),
				user: getUser(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepo{
				db: tt.fields.db,
			}
			err := r.Create(tt.args.ctx, tt.args.user)
			require.NoError(t, err)
		})
	}
}

func TestUserRepo_GetByEmail(t *testing.T) {
	arg := getUser()

	ctx := context.Background()

	repo := UserRepo{
		db: db,
	}

	err := repo.Create(ctx, arg)
	require.NoError(t, err)

	user, err := repo.GetByEmail(ctx, arg.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Address, user.Address)
}

func getUser() domain.Register {
	arg := domain.Register{
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
		Phone:    util.RandomString(8),
		Address:  util.RandomString(5),
	}
	return arg
}

func TestUserRepo_GetUserById(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Register
		wantErr bool
	}{
		{
			fields: fields{db: db},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepo{
				db: tt.fields.db,
			}
			got, err := r.GetUserById(tt.args.ctx, tt.args.id)
			require.NoError(t, err)

			require.NotEmpty(t, got)
		})
	}
}

func TestUserRepo_SetSession(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		userID    int
		token     string
		expiredAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{
				db: db,
			},
			args: args{
				ctx:       context.Background(),
				userID:    1,
				token:     util.RandomString(10),
				expiredAt: time.Now().Add(10 * time.Minute),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepo{
				db: tt.fields.db,
			}
			err := r.SetSession(tt.args.ctx, tt.args.userID, tt.args.token, tt.args.expiredAt)
			require.NoError(t, err)
		})
	}
}

func TestUserRepo_GetUserByRefreshToken(t *testing.T) {
	refresh_token := util.RandomString(10)
	expiresAt := time.Now()

	repo := UserRepo{
		db: db,
	}

	err := repo.SetSession(context.Background(), 1, refresh_token, expiresAt)
	require.NoError(t, err)

	user, err := repo.GetUserByRefreshToken(context.Background(), refresh_token)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.RefreshToken, refresh_token)
}

func TestNewUserRepo(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *UserRepo
	}{
		// TODO: Add test cases.
		{
			name: "ok",
			args: args{
				db: db,
			},
			want: &UserRepo{db: db},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepo(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}
