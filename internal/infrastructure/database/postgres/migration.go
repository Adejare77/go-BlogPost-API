package postgres

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Comment{},
		&entity.Like{},
		&entity.RefreshToken{},
	)
}
