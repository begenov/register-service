package domain

import "time"

type Register struct {
	ID       int       `json:"id"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Role     string    `json:"role"`
	Address  string    `json:"address"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
