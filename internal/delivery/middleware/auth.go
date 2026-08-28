package middleware

import (
	"net/http"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService auth.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		claimsAny, err := tokenService.Validate(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}
		claims := claimsAny.(*Claims)

		ctx.Set("userID", claims.UserID)
		ctx.Set("isStaff", claims.IsStaff)

		ctx.Next()
	}
}
