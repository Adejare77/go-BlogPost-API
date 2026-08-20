package user

import "github.com/gin-gonic/gin"

var UserRoute = func(r *gin.RouterGroup, h *UserHandler) {
	r.DELETE("", h.DeleteByID)
	r.PATCH("", h.Update)
}
