package like

type LikeRequestPostPath struct {
	PostID string `uri:"post_id" binding:"required,uuid"`
}

type LikeRequestCommentPath struct {
	CommentID string `uri:"comment_id" binding:"required,uuid"`
}
