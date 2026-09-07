package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)


type CommentDetailRow struct {
	ID entity.CommentID
	AuthorID entity.UserID
	FullName string
	PostID entity.PostID
	ParentID *entity.CommentID
	Content string
	Likes int
	Liked bool
	ReplyCount int
	TopReplies []CommentListRow `gorm:"-"`
	CreatedAt time.Time
}
