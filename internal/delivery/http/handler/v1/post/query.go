package post

type PostQueryRequest struct {
	Status string `form:"status" binding:"oneof=all draft published"`
	Author string `form:"author"`
}
