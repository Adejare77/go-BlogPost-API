package post

import "github.com/gin-gonic/gin"

var PostProtectedRoute = func(r *gin.RouterGroup, h *PostHandler) {
	r.POST("", h.Create)
	r.DELETE("", h.DeleteByID)
	r.PATCH("", h.Update)
}

var PostPublicRoute = func(r *gin.RouterGroup, h *PostHandler) {
	r.GET("", h.FindAll)
	r.GET("", h.FindByID)
}
