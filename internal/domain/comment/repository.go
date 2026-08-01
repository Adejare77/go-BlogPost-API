package comment

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type CommentRepository interface {
	Create(comment *entity.Comment) error
	FindByID(commentID entity.CommentID) (*entity.Comment, error)
	Update(comment *entity.Comment) error
	DeleteByID(commentID entity.CommentID) error
	FindByPostID(postID entity.PostID) ([]*CommentDetail, error)
}
