package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type CommentCreateResponse struct {
	ID entity.CommentID
	Author AuthorSummary
	PostID entity.PostID
	ParentID *entity.CommentID
	Content string
	CreatedAt time.Time
}
