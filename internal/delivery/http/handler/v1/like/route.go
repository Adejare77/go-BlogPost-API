package like

import (
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/middleware"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.RouterGroup,
	h *LikeHandler,
	tokenService auth.TokenService) {
	protected := r.Group("")
	protected.Use(middleware.AuthMiddleware(tokenService))

	protected.POST("/comments/:comment_id/likes", h.CreateLikeComment)
	protected.DELETE("/comments/:comment_id/likes", h.DeleteLikedComment)

	protected.POST("/posts/:post_id/likes", h.CreateLikePost)
	protected.DELETE("/posts/:post_id/likes", h.DeleteLikedPost)
}
