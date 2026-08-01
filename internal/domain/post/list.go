package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type PostList struct {
	ID entity.PostID
	Author AuthorSummary
	Title string
	Excerpt string
	Likes int
	Liked bool
	IsPublished bool
	CommentCount int
	CreatedAt time.Time
}
