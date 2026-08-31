package auth

import (
	"errors"
	"net/http"

	"github.com/Adejare77/go-BlogPost-API/internal/config"
	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *usecase.AuthService
}

func NewAuthHandler(authservice *usecase.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authservice,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req AuthRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, req, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"invalid_request",
			"invalid request",
			err,
		)
		return
	}

	auth, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	response := AuthTokenResponse{
		AccessToken: auth.AccessToken,
		UserID: auth.UserID,
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
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		httperrors.HandleRequestError(
			ctx,
			http.StatusUnauthorized,
			"unauthorized",
			"unauthorized",
			err,
		)
		return
	}

	if err := h.authService.Logout(refreshToken); err != nil {
		httperrors.HandleError(ctx, err)
		return
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
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, req, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"invalid_request",
			"invalid request",
			err,
		)
		return
	}

	user := entity.User{
		FullName: req.FullName,
		Email: req.Email,
		Password: &req.Password,
	}

	if err := h.authService.Register(&user); err != nil {
		httperrors.HandleError(ctx, err)
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
		httperrors.HandleError(ctx, err)
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

func (h *AuthHandler) RefreshToken(ctx *gin.Context){
	token, err := ctx.Cookie("refresh_token")
	if err != nil {
		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"refresh_token",
			"invalid refresh token",
			err,
		)
		return
	}

	auth, err := h.authService.RefreshToken(token)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name: "refresh_token",
		Value: auth.RefreshToken,
		MaxAge: int(config.Current.App.RefreshTokenTTL.Seconds()),
		Path: "/",
		SameSite: http.SameSiteLaxMode,
		Secure: false,
		HttpOnly: true,
	})

	response := AuthTokenResponse {
		AccessToken: auth.AccessToken,
		UserID: auth.UserID,
	}

	ctx.JSON(http.StatusOK, response)
}
