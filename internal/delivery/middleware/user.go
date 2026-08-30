package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isStaff := ctx.MustGet("isStaff").(bool)
		if !isStaff {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "forbidden",
				"message": "forbidden",
			})
			return
		}

		ctx.Next()
	}
}
