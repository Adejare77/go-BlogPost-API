package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type PostCreateResponse struct {
	ID entity.PostID
	Author AuthorSummary
	Title string
	Content string
	IsPublished bool
	CreatedAt time.Time
}
