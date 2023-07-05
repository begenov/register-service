package domain

import "time"

type Register struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Role         string    `json:"role"`
	Address      string    `json:"address"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"create_at"`
}

type RequestSignIn struct {
	Email    string
	Password string
	Role     string
}
