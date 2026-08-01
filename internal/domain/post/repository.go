package post

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type PostRepository interface {
	Create(post *entity.Post) error
	FindByID(postID entity.PostID) (*entity.Post, error)
	Update(post *entity.Post) error
	DeleteByID(postID entity.PostID) error
	FindAll() ([]*PostList, error)
	FindPostDetail(postID entity.PostID, userID entity.UserID) ([]*PostDetail, error)
}
