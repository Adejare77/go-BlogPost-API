package post

import "github.com/google/uuid"

type PostPathRequest struct {
	PostID uuid.UUID `uri:"post_id" binding:"required,uuid"`
}
