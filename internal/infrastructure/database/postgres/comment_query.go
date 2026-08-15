package postgres

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type CommentDetailRow struct {
	ID entity.CommentID
	AuthorID entity.UserID
	FullName string
	PostID entity.PostID
	Content string
	Likes int
	Liked bool
	ReplyCount int
	TopReplies []ReplySummaryRow
	CreatedAt time.Time
}

type ReplySummaryRow struct {
	ID entity.CommentID
	AuthorID entity.UserID
	FullName string
	PostID entity.PostID
	ParentID entity.CommentID
	Excerpt string
	Likes int
	Liked bool
	ReplyCount int
	CreatedAt time.Time
}

type CommentListRow struct {
	ID entity.CommentID
	AuthorID entity.UserID
	FullName string
	PostID entity.PostID
	Excerpt string
	Likes int
	Liked bool
	ReplyCount int
	CreatedAt time.Time
}
