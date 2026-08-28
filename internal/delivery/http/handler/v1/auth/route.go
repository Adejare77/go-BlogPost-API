package auth

import (
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/middleware"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

var authPublicRoute = func(r *gin.RouterGroup, h *AuthHandler) {
	r.POST("/login", h.Login)
	r.POST("/register", h.Create)
	r.POST("/logout", h.Logout)
}

var authPrivateRoute = func(r *gin.RouterGroup, h *AuthHandler) {
	r.GET("/me", h.Me)
}

func RegisterRoutes(
	r *gin.RouterGroup,
	h *AuthHandler,
	tokenService auth.TokenService,
) {
	auth := r.Group("/auth")

	authPublicRoute(auth, h)

	protected := auth.Group("")
	protected.Use(middleware.AuthMiddleware(tokenService))
	authPrivateRoute(protected, h)
}
