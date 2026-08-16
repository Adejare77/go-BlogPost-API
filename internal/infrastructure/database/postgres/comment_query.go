package postgres

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
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
	TopReplies []comment.ReplySummary
	CreatedAt time.Time
}

func (c *CommentDetailRow) ToCommentDetail() *comment.CommentDetail {
	return &comment.CommentDetail{
		ID: c.ID,
		Author: comment.AuthorSummary{
			ID: c.AuthorID,
			FullName: c.FullName,
		},
		PostID: c.PostID,
		Content: c.Content,
		Likes: c.Likes,
		Liked: c.Liked,
		ReplyCount: c.ReplyCount,
		TopReplies: c.TopReplies,
		CreatedAt: c.CreatedAt,
	}
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

func (r *ReplySummaryRow) ToReplySummary() comment.ReplySummary {
	return comment.ReplySummary{
		ID: r.ID,
		Author: comment.AuthorSummary{
			ID: r.AuthorID,
			FullName: r.FullName,
		},
		PostID: r.PostID,
		ParentID: r.ParentID,
		Excerpt: r.Excerpt,
		Likes: r.Likes,
		Liked: r.Liked,
		ReplyCount: r.ReplyCount,
		CreatedAt: r.CreatedAt,
	}
}

type CommentListRow struct {
	ID entity.CommentID
	AuthorID entity.UserID
	FullName string
	PostID entity.PostID
	Content string
	Likes int
	Liked bool
	ReplyCount int
	CreatedAt time.Time
}

func (c *CommentListRow) ToCommentList() comment.CommentList {
	return comment.CommentList{
		ID: c.ID,
		Author: comment.AuthorSummary{
			ID: c.AuthorID,
			FullName: c.FullName,
		},
		PostID: c.PostID,
		Excerpt: ToExcerpt(c.Content),
		Likes: c.Likes,
		Liked: c.Liked,
		ReplyCount: c.ReplyCount,
		CreatedAt: c.CreatedAt,
	}
}
