package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthorSummary struct {
	ID entity.UserID `json:"id"`
	FullName string `json:"full_name"`
}

type PostListResponse struct {
	ID entity.PostID `json:"id"`
	Author AuthorSummary `json:"author"`
	Title string `json:"title"`
	Excerpt string `json:"excerpt"`
	Likes int `json:"likes"`
	Liked bool `json:"liked"`
	CommentCount int `json:"comment_count"`
	IsPublished bool `json:"is_published"`
	CreatedAt time.Time `json:"created_at"`
}

type PostDetailResponse struct {
	ID entity.PostID `json:"id"`
	Author AuthorSummary `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	Likes int `json:"likes"`
	Liked bool `json:"liked"`
	IsPublished bool `json:"is_published"`
	CommentCount int `json:"comment_count"`
	TopComments []CommentSummary `json:"top_comments"`
	CreatedAt time.Time `json:"created_at"`
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

type PostCreateResponse struct {
	ID entity.PostID `json:"id"`
	Author AuthorSummary `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	IsPublished bool `json:"is_published"`
	CreatedAt time.Time `json:"created_at"`
}
