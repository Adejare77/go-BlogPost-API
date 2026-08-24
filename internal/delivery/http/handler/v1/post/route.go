package post

import (
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/middleware"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

var postProtectedRoute = func(r *gin.RouterGroup, h *PostHandler) {
	r.POST("/posts", h.Create)
	r.DELETE("/posts/:post_id", h.DeleteByID)
	r.PATCH("/Posts/:post_id", h.Update)
}

var postPublicRoute = func(r *gin.RouterGroup, h *PostHandler) {
	r.GET("/posts", h.FindAll)
	r.GET("/posts/:post_id", h.FindByID)
}


func RegisterRoutes(
	r *gin.RouterGroup,
	h *PostHandler,
	tokenService auth.TokenService,
	) {
	posts := r.Group("/posts")

	postPublicRoute(posts, h)

	protected := posts.Group("")
	protected.Use(middleware.AuthMiddleware(tokenService))

	postProtectedRoute(protected, h)
}
