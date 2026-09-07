package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase/comment"
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


func ToPostDetail(post *post.PostDetailRow) *PostDetail {
	TopComments := comment.ToCommentList(post.TopComments)

	return &PostDetail{
		ID: post.ID,
		Author: AuthorSummary{
			ID: post.AuthorID,
			FullName: post.FullName,
		},
		Title: post.Title,
		Content: post.Content,
		Likes: post.Likes,
		Liked: post.Liked,
		IsPublished: post.IsPublished,
		CommentCount: post.CommentCount,
		TopComments: TopComments,
		CreatedAt: post.CreatedAt,
	}
}
