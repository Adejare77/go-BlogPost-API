package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthorSummary struct {
	ID entity.UserID
	FullName string
}

type CommentSummary struct {
	ID entity.CommentID
	Author AuthorSummary
	PostID entity.PostID
	Excerpt string
	Likes int
	Liked bool
	ReplyCount int
	CreatedAt time.Time
}

type PostDetail struct {
	ID entity.PostID
	Author AuthorSummary
	Title string
	Content string
	Likes int
	Liked bool
	IsPublished bool
	CommentCount int
	TopComments []CommentSummary
	CreatedAt time.Time
}
