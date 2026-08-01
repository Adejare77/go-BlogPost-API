package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type PostAuthorInfo struct {
	ID entity.UserID `json:"id"`
	FullName string `json:"full_name"`
}

type PostListResponse struct {
	ID entity.UserID `json:"id"`
	Author PostAuthorInfo `json:"author"`
	Title string `json:"title"`
	Excerpt string `json:"excerpt"`
	Likes int `json:"likes"`
	Liked bool `json:"liked"`
	CommentCount int `json:"comment_count"`
	IsPublished bool `json:"is_published"`
	CreatedAt time.Time `json:"created_at"`
}

type PostDetailResponse struct {
	ID entity.UserID `json:"id"`
	Author PostAuthorInfo `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
	Likes int `json:"likes"`
	Liked bool `json:"liked"`
	CommentCount int `json:"comment_count"`
	TopComments []PostListResponse `json:"top_comments"`
	IsPublished bool `json:"is_published"`
	CreatedAt time.Time `json:"created_at"`
}
