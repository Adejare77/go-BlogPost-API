package postgres

import (
	"fmt"

	"github.com/Adejare77/go-BlogPost-API/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSL,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	SQLDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	SQLDB.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	SQLDB.SetMaxOpenConns(cfg.MaxOpenConns)
	SQLDB.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}
