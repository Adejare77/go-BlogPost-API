package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

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
