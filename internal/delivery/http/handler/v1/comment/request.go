package comment


type CommentRequest struct {
	Content string `json:"content" binding:"required"`
}
