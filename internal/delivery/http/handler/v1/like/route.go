package like

import "github.com/gin-gonic/gin"


var LikeRoute = func(r *gin.RouterGroup, h LikeHandler) {
	r.POST("", h.CreateLikeComment)
	r.POST("", h.CreateLikePost)
	r.DELETE("", h.DeleteLikedComment)
	r.DELETE("", h.DeleteLikedPost)
}
