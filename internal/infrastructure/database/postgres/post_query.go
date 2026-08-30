package postgres

import (
	"strings"
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"gorm.io/gorm"
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
	TopComments []comment.CommentList
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
		Excerpt: ToExcerpt(p.Content),
		Likes: p.Likes,
		Liked: p.Liked,
		IsPublished: p.IsPublished,
		CommentCount: p.CommentCount,
		CreatedAt: p.CreatedAt,
	}
}

func ToExcerpt(content string) string {
	words := strings.Fields(content)

	if len(words) > 30 {
		words = words[:30]
	}

	return strings.Join(words, " ")
}


func PostQueryScope(userID entity.UserID, query post.PostQuery) func (*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// IsPublished filter
		switch query.Status {
		case "published":
			db = db.Where("is_published = ?", true)

		case "draft":
			db = db.Where("is_published = ?", false)

		default:
			db = db
		}

		// Author filter
		switch query.Author {
		case "":
			return db

		case "me":
			return db.Where("author_id = ?", userID)

		default:
			return db.Where("full_name = ?",  query.Author)
		}
	}
}
