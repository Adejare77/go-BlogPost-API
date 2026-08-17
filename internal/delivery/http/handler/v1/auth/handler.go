package auth

import (
	"net/http"

	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *usecase.AuthService
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req AuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Take care of errors
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	auth, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		// Take care of errors
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := AuthResponse{
		AccessToken: auth.AccessToken,
		UserID: auth.UserID,
		Email: auth.Email,
	}

	ctx.SetCookie(
		"refresh_token",
		auth.RefreshToken,
		15*60,
		"/",
		"lax",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, response)
}
