package auth

import (
	"net/http"

	"github.com/Adejare77/go-BlogPost-API/internal/config"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
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

	response := AuthTokenResponse{
		AccessToken: auth.AccessToken,
		UserID: auth.UserID,
		Email: auth.Email,
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "refresh_token",
		Value: auth.RefreshToken,
		Path: "/",
		MaxAge: int(config.Current.App.RefreshTokenTTL.Seconds()),
		Secure: false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})

	ctx.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(entity.UserID)

	if err := h.authService.Logout(userID); err != nil {
		// Take care of error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "refresh_token",
		Value: "",
		Path: "/",
		MaxAge: -1,
	})

	ctx.Status(http.StatusOK)

}

func (h *AuthHandler) Create(ctx *gin.Context) {
	var req AuthRegister

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// log the error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := entity.User{
		FullName: req.FullName,
		Email: req.Email,
		Password: &req.Password,
	}

	if err := h.authService.Register(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, AuthResponse{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	})
}

func (h *AuthHandler) Me(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(entity.UserID)

	user, err := h.authService.FindByID(userID)

	if err != nil {
		// work on error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error. Try again later",
		})
		return
	}

	response := AuthResponse{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
