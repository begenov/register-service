package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/begenov/register-service/internal/domain"
)

type RestaurantRepo struct {
	db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{
		db: db,
	}
}

func (r *RestaurantRepo) Create(ctx context.Context, restaurant domain.Register) error {
	stmt := `INSERT INTO restaurant (email, phone, home_address, password_hash) VALUES ($1, $2, $3, $4)`
	if _, err := r.db.ExecContext(ctx, stmt, restaurant.Email, restaurant.Phone, restaurant.Address, restaurant.Password); err != nil {
		return err
	}

	return nil
}

func (r *RestaurantRepo) GetRestaurantByEmail(ctx context.Context, email string) (domain.Register, error) {
	stmt := `SELECT email, phone, home_address, password_hash FROM restaurant WHERE email = $1`
	row := r.db.QueryRowContext(ctx, stmt, email)
	var restaurant domain.Register
	if err := row.Scan(&restaurant.Email, &restaurant.Phone, &restaurant.Address, &restaurant.Password); err != nil {
		return restaurant, err
	}

	return restaurant, nil
}

func (r *RestaurantRepo) SetSession(ctx context.Context, token string, expiredAt time.Time, restaurantId int) error {
	stmt := `UPDATE restaurant SET refresh_token=$1, expired_at=$2 WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, stmt, token, expiredAt, restaurantId); err != nil {
		return err
	}

	return nil
}

func (r *RestaurantRepo) GetRestaurantByID(ctx context.Context, id int) (domain.Register, error) {
	stmt := `SELECT email, phone, password_hash FROM restaurant WHERE id = $1`
	row := r.db.QueryRowContext(ctx, stmt, id)
	var restaurant domain.Register
	if err := row.Scan(&restaurant.Email, &restaurant.Phone, &restaurant.Password); err != nil {
		return restaurant, err
	}

	return restaurant, nil
}

func (r *RestaurantRepo) GetRestaurantByRefreshToken(ctx context.Context, token string) (domain.Register, error) {
	stmt := `SELECT id, email, phone, home_address, password_hash, refresh_token, expired_at FROM "restaurant" WHERE refresh_token = $1`
	row := r.db.QueryRowContext(ctx, stmt, token)

	var restaurant domain.Register
	if err := row.Scan(&restaurant.ID, &restaurant.Email, &restaurant.Phone, &restaurant.Address, &restaurant.Password, &restaurant.RefreshToken, &restaurant.ExpiresAt); err != nil {
		return restaurant, err
	}

	return restaurant, nil
}
