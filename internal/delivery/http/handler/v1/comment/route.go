package comment

import "github.com/gin-gonic/gin"

var CommentProtectedRoute = func(r *gin.RouterGroup, h *CommentHandler) {
	r.POST("", h.CreateComment)
	r.POST("", h.CreateReply)
	r.PATCH("", h.Update)
	r.DELETE("", h.DeleteByID)
}

var CommentPublicRoute = func(r *gin.RouterGroup, h *CommentHandler) {
	r.GET("", h.FindByID)
}
