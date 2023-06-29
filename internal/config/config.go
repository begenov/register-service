package config

import (
	"fmt"
	"os"
	"time"

	"github.com/subosito/gotenv"
)

const (
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
)

type Config struct {
	JWT    JWTConfig
	DB     DatabaseConfig
	Server ServerConfig
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	SigningKey      string
}

type DatabaseConfig struct {
	Driver string
	DSN    string
}

type ServerConfig struct {
	Port               string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

func NewConfig() (*Config, error) {
	err := gotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables from file: %v", err)
	}

	driver := os.Getenv("DRIVER")
	dsn := os.Getenv("DSN_STUDENTS")
	jwtKey := os.Getenv("SIGNIN_KEY")

	return &Config{
		JWT: JWTConfig{
			AccessTokenTTL:  defaultAccessTokenTTL,
			RefreshTokenTTL: defaultRefreshTokenTTL,
			SigningKey:      jwtKey,
		},
		DB: DatabaseConfig{
			Driver: driver,
			DSN:    dsn,
		},
		Server: ServerConfig{
			Port:               defaultHTTPPort,
			WriteTimeout:       defaultHTTPRWTimeout,
			ReadTimeout:        defaultHTTPRWTimeout,
			MaxHeaderMegabytes: defaultHTTPMaxHeaderMegabytes,
		},
	}, nil
}
