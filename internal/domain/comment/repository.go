package comment

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type CommentRepository interface {
	Create(comment *entity.Comment) error
	// FindByID(commentID entity.CommentID) (*entity.Comment, error)
	FindByID(commentID entity.CommentID, userID *entity.UserID) (*CommentDetail, error)
	Update(comment *entity.Comment) (*CommentDetail, error)
	DeleteByID(commentID entity.CommentID, userID entity.UserID) error
	// FindByCommentID(commentID entity.CommentID) ([]*CommentDetail, error)
}
