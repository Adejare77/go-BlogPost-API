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

type PostDetail struct {
	ID entity.PostID
	Author AuthorSummary
	Title string
	Content string
	Likes int
	Liked bool
	IsPublished bool
	CommentCount int
	TopComments []comment.CommentList
	CreatedAt time.Time
}
