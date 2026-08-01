package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type CommentList struct {
	ID entity.CommentID
	Author AuthorSummary
	PostID entity.PostID
	Excerpt string
	Likes int
	Liked bool
	ReplyCount int
	CreatedAt time.Time
}
