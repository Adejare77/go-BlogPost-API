package post

type PostQueryRequest struct {
	Status string `form:"status" binding:"omitempty,oneof=all draft published"`
	Author string `form:"author,omitempty"`
}
