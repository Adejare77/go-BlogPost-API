package post

type PostCreateRequest struct {
	Title string `json:"title" binding:"required,max=100"`
	Content string `json:"content" binding:"required"`
	IsPublished bool `json:"is_published"`
}

type PostUpdateRequest struct{
	Title string `json:"title" binding:"required,max=100"`
	Content string `json:"content" binding:"required"`
	IsPublished bool `json:"is_published"`
}
