package postgres

import (
	"strings"
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
)

type PostDetailRow struct {
	ID entity.PostID
	UserID entity.UserID
	FullName string
	Title string
	Content string
	Likes int
	Liked bool
	IsPublished bool
	CommentCount int
	TopComments []post.CommentSummary
	CreatedAt time.Time
}


func (detail *PostDetailRow) ToPostDetail() *post.PostDetail {
	return &post.PostDetail{
		ID: detail.ID,
		Author: post.AuthorSummary{
			ID: detail.UserID,
			FullName: detail.FullName,
		},
		Title: detail.Title,
		Content: detail.Content,
		Likes: detail.Likes,
		IsPublished: detail.IsPublished,
		CommentCount: detail.CommentCount,
		TopComments: detail.TopComments,
		CreatedAt: detail.CreatedAt,
	}
}

type PostListRow struct {
	ID entity.PostID
	AuthorID entity.UserID
	FullName string
	Title string
	Content string
	Likes int
	Liked bool
	IsPublished bool
	CommentCount int
	CreatedAt time.Time
}

func (p PostListRow) ToPostList() post.PostList {
	return post.PostList{
		ID: p.ID,
		Author: post.AuthorSummary{
			ID: p.AuthorID,
			FullName: p.FullName,
		},
		Title: p.Title,
		Excerpt: ToExercept(p.Content),
		Likes: p.Likes,
		Liked: p.Liked,
		IsPublished: p.IsPublished,
		CommentCount: p.CommentCount,
		CreatedAt: p.CreatedAt,
	}
}

func ToExercept(content string) string {
	words := strings.Fields(content)

	if len(words) > 30 {
		words = words[:30]
	}

	return strings.Join(words, " ")
}
