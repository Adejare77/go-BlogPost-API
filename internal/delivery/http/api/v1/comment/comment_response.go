package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type CommentAuthorInfo struct {
	ID entity.UserID `json:"id"`
	FullName string `json:"full_name"`
}

type CommentListResponse struct {
	ID entity.CommentID `json:"id"`
	Author CommentAuthorInfo `json:"author"`
	PostID entity.PostID `json:"post_id"`
	Excerpt string `json:"excerpt"`
	Likes uint `json:"likes"`
	Liked bool `json:"liked"`
	Reply_count uint `json:"reply_count"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentDetailResponse struct {
	ID entity.CommentID `json:"id"`
	Author CommentAuthorInfo `json:"author"`
	PostID entity.PostID `json:"post_id"`
	Excerpt string `json:"excerpt"`
	Likes uint `json:"likes"`
	Liked bool `json:"liked"`
	Reply_count uint `json:"reply_count"`
	TopReplies []CommentListResponse `json:"top_replies"`
	CreatedAt time.Time `json:"created_at"`
}
