package comment

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type CommentRepository interface {
	Create(comment *entity.Comment) error
	FindByID(commentID entity.CommentID, userID entity.UserID) (*CommentDetailRow, error)
	Update(comment *entity.Comment) (*CommentDetailRow, error)
	DeleteByID(commentID entity.CommentID, userID entity.UserID) error
	FindByPostID(postID entity.PostID, userID entity.UserID) ([]CommentListRow, error)
}
