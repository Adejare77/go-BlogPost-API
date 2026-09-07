package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthorSummary struct {
	ID entity.UserID
	FullName string
}

type PostDetailRow struct {
	ID entity.PostID
	AuthorID entity.UserID
	FullName string
	Title string
	Content string
	Likes int
	Liked bool
	IsPublished bool
	CommentCount int
	TopComments []comment.CommentListRow `gorm:"-"`
	CreatedAt time.Time
}
