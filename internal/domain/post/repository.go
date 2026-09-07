package post

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type PostRepository interface {
	Create(post *entity.Post) error
	FindByID(postID entity.PostID, userID entity.UserID) (*PostDetailRow, error)
	Update(post *entity.Post) (*PostDetailRow, error)
	DeleteByID(postID entity.PostID, userID entity.UserID) error
	FindAll(userID entity.UserID, query PostQuery) ([]PostListRow, error)
}
