package comment

type CommentPathRequest struct {
	CommentID string `uri:"comment_id" binding:"required,uuid"`
}
