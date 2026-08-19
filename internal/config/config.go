package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var Current Config

type Config struct {
	App AppConfig
	DB DBConfig
}

type AppConfig struct {
	Env string
	JWT string
	JWTIssuer string

	LogLevel string
	Port string

	AccessTokenTTL time.Duration
	RefreshTokenTTL time.Duration
}

type DBConfig struct {
	Host string
	Name string
	Password string
	Port string
	SSL string
	User string

	MaxIdleConns int
	MaxOpenConns int
	ConnMaxLifeTime time.Duration
}

func Load() (*Config, error) {
	if err := godotenv.Overload(); err != nil {
		log.Fatalf("Could not read .env file: %s", err.Error())
	}

	if err := validateRequired(
		"DB_NAME",
		"DB_USER",
		"JWT_SECRET",
	); err != nil {
		return nil, err
	}

	cfg := &Config{
		App: AppConfig{
			Env: optional("APP_ENV", "development"),
			JWT: os.Getenv("JWT_SECRET"),
			JWTIssuer: optional("JWT_ISSUER", "Rashisky"),
			LogLevel: optional("LOG_LEVEL", "0.9"),
			Port: optional("APP_PORT", "8080"),
			AccessTokenTTL: getEnvDuration("ACCESS_TOKEN_TTL", "15m"),
			RefreshTokenTTL: getEnvDuration("REFRESH_TOKEN_TTL", "7d"),
		},
		DB: DBConfig{
			Host: optional("DB_HOST", "localhost"),
			Name: os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			SSL: optional("SSL", "disable"),
			MaxIdleConns: getEnvInt("MAX_IDLE_CONNS", 10),
			MaxOpenConns: getEnvInt("MAX_OPEN_CONNS", 20),
			ConnMaxLifeTime: getEnvDuration("CONN_MAX_LIFE_TIME", "30m"),
		},
	}

	return cfg, nil
}
