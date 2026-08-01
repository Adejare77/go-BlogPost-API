package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthorSummary struct {
	ID entity.UserID
	FullName string
}

type CommentDetail struct {
	ID entity.CommentID
	Author AuthorSummary
	Content string
	Likes int
	Liked bool
	ReplyCount int
	TopReplies []ReplySummary
	CreatedAt time.Time
}

type ReplySummary struct {
	ID entity.CommentID
	Author AuthorSummary
	ParentID entity.CommentID
	Excerpt string
	Likes int
	Liked bool
	ReplyCount int
	CreatedAt time.Time
}
