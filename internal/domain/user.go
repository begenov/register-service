package domain

type Session struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type Token struct {
	RefreshToken string `json:"refreshtoken"`
}

type Role struct {
	UserID int    `json:"userid"`
	Roles  string `json:"roles"`
}
