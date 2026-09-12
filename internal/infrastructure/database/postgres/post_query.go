package postgres

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"gorm.io/gorm"
)


func PostQueryScope(userID entity.UserID, query post.PostQuery) func (*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// IsPublished filter
		switch query.Status {
		case "published":
			db = db.Where("is_published = ?", true)

		case "draft":
			db = db.Where("is_published = ?", false)

		default:
			db = db
		}

		// Author filter
		switch {
		case query.Author == "" && userID == entity.UserID(0):
			return db

		case query.Author == "" || query.Author == "me":
			return db.Where("author_id = ?", userID)

		default:
			return db.Where("full_name = ?",  query.Author)
		}
	}
}
