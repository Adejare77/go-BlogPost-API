package comment

import (
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/middleware"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

var commentProtectedRoute = func(r *gin.RouterGroup, h *CommentHandler) {
	r.POST("/posts/:post_id/comments", h.CreateComment)
	r.POST("/comments/:comment_id/replies", h.CreateReply)
	r.PATCH("/comments/:comment_id", h.Update)
	r.DELETE("/comments/:comment_id", h.DeleteByID)
}

var commentPublicRoute = func(r *gin.RouterGroup, h *CommentHandler) {
	r.GET("/comments", h.FindByID)
}


func RegisterRoutes(
	r *gin.RouterGroup,
	h *CommentHandler,
	tokenService auth.TokenService,
	) {
		commentPublicRoute(r, h)

		protected := r.Group("")
		protected.Use(middleware.AuthMiddleware(tokenService))
		commentProtectedRoute(protected, h)
}
