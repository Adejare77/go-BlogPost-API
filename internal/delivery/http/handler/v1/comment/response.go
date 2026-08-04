package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type AuthorSummary struct {
	ID entity.UserID `json:"id"`
	FullName string `json:"full_name"`
}

type CommentListResponse struct {
	ID entity.CommentID `json:"id"`
	Author AuthorSummary `json:"author"`
	PostID entity.PostID `json:"post_id"`
	Excerpt string `json:"excerpt"`
	Likes int `json:"likes"`
	Liked bool `json:"liked"`
	ReplyCount int `json:"reply_count"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentDetailResponse struct {
	ID entity.CommentID `json:"id"`
	Author AuthorSummary `json:"author"`
	Content string `json:"comment"`
	Likes int `json:"likes"`
	Liked bool `json:"liked"`
	ReplyCount int `json:"reply_count"`
	TopReplies []CommentListResponse `json:"top_replies"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentCreateResponse struct {
	ID entity.CommentID `json:"id"`
	Author AuthorSummary `json:"author"`
	PostID entity.PostID `json:"post_id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
