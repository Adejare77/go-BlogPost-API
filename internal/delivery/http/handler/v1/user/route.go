package user

import (
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/middleware"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

var userRoute = func(r *gin.RouterGroup, h *UserHandler) {
	r.GET("", h.FindAll)
	r.GET("/:user_id", h.FindByID)
	r.DELETE("/:user_id", h.DeleteByID)
	r.PATCH("/:user_id", h.Update)
}


func RegisterRoutes(
	r *gin.RouterGroup,
	h *UserHandler,
	tokenService auth.TokenService,
	) {
	protected := r.Group("/users")
	protected.Use(middleware.AuthMiddleware(tokenService))
	protected.Use(middleware.UserMiddleware())

	userRoute(protected, h)
}
