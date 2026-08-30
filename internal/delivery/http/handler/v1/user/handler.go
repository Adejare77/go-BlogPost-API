package user

import (
	"errors"
	"net/http"

	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type UserHandler struct {
	userService *usecase.UserService
}

func NewUserHandler(userSerivce *usecase.UserService) *UserHandler {
	return &UserHandler{
		userService: userSerivce,
	}
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	response := make([]UserResponse, len(users))

	for i, user := range users {
		response[i] = UserResponse{
			ID: user.ID,
			FullName: user.FullName,
			Email: user.Email,
			CreatedAt: user.CreatedAt,
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": response,
	})
}

func (h *UserHandler) FindByID(ctx *gin.Context) {
	var reqPath UserPathRequest

	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, reqPath, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"user_id",
			"must be a valid integer",
			err,
		)
		return
	}

	targetUserID := entity.UserID(reqPath.UserID)

	user, err := h.userService.FindByID(targetUserID)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	response := UserResponse{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteByID(ctx *gin.Context) {
	var reqPath UserPathRequest

	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, reqPath, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"user_id",
			"must be a valid integer",
			err,
		)
		return
	}

	targetUserID := entity.UserID(reqPath.UserID)

	if err := h.userService.DeleteByID(targetUserID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	var req UserUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, req, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"user_id",
			"must be a valid integer",
			err,
		)
		return
	}

	var reqPath UserPathRequest

	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, reqPath, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"user_id",
			"must be a valid integer",
			err,
		)
		return
	}

	targetUserID := entity.UserID(reqPath.UserID)

	user := entity.User{
		ID: targetUserID,
		FullName: req.FullName,
		Password: &req.Password,
	}

	result, err := h.userService.Update(&user);
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	response := UserResponse{
		ID: result.ID,
		FullName: result.FullName,
		Email: result.Email,
		CreatedAt: result.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *UserHandler) EnableByID(ctx *gin.Context) {
	var reqPath UserPathRequest

	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, reqPath, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"user_id",
			"must be a valid integer",
			err,
		)
		return
	}

	targetUserID := entity.UserID(reqPath.UserID)

	if err := h.userService.EnableByID(targetUserID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) DisableByID(ctx *gin.Context) {
	var reqPath UserPathRequest

	if err := ctx.ShouldBindUri(&reqPath); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, reqPath, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"user_id",
			"must be a valid integer",
			err,
		)
		return
	}

	targetUserID := entity.UserID(reqPath.UserID)

	if err := h.userService.DisableByID(targetUserID); err != nil {
		httperrors.HandleError(ctx, err)
	}

	ctx.Status(http.StatusNoContent)
}
