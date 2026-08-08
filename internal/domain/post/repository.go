package post

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type PostRepository interface {
	Create(post *entity.Post) error
	// FindByID(postID entity.PostID) (*entity.Post, error)
	FindByID(postID entity.PostID, userID *entity.UserID) (*PostDetail, error)
	Update(post *entity.Post) (*PostDetail, error)
	DeleteByID(postID entity.PostID, userID entity.UserID) error
	FindAll() ([]*PostList, error)
	// FindPostDetail(postID entity.PostID, userID *entity.UserID) ([]*PostDetail, error)
}
