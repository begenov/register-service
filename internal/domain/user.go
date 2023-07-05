package domain

import "time"

type Token struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type Session struct {
	RefreshToken string    `json:"refreshtoken"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Role struct {
	UserID int    `json:"userid"`
	Roles  string `json:"roles"`
}

var (
	User       = "user"
	Courier    = "courier"
	Restaurant = "restaurant"
)
