package post

type PostPathRequest struct {
	PostID string `uri:"post_id" binding:"required,uuid"`
}
