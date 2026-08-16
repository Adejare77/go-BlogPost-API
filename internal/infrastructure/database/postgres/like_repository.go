package postgres

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"gorm.io/gorm"
)

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{
		db: db,
	}
}

func (repo *LikeRepository) Create(like *entity.Like) error {
	return repo.db.Create(like).Error
}

func (repo *LikeRepository) DeleteByUserAndPost(userID entity.UserID, postID entity.LikeID) error {
	return repo.db.
	Where("id = ? AND likeable_type = 'post' AND author_id = ?", postID, userID).
	Delete(&entity.Like{}).Error
}

func (repo *LikeRepository) DeleteByUserAndComment(userID entity.UserID, commentID entity.LikeID) error {
	return repo.db.
	Where("id = ? AND likeable_type = 'comment' AND author_id = ?", commentID, userID).
	Delete(&entity.Like{}).Error
}
