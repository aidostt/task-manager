package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	DB struct {
		User     string
		Password string
		Name     string
		Host     string
		Port     string
	}
	HTTP struct {
		Host string
		Port string
	}
	JWT struct {
		Secret          string
		AccessTokenTTL  string
		RefreshTokenTTL string
	}
)
type Config struct {
	DB   DB
	HTTP HTTP
	Jwt  JWT
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: DB{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
		HTTP: HTTP{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Jwt: JWT{
			Secret:          os.Getenv("JWT_SECRET"),
			AccessTokenTTL:  os.Getenv("ACCESS_TOKEN_TTL"),
			RefreshTokenTTL: os.Getenv("REFRESH_TOKEN_TTL"),
		},
	}, nil
}
