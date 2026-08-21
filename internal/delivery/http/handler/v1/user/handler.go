package user

import (
	"net/http"
	"strconv"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	userService *usecase.UserService
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		// work on error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error. Try again later",
		})
		return
	}

	response := make([]UserResponse, 0, len(users))

	for _, user := range users {
		response = append(response, UserResponse{
			ID: user.ID,
			FullName: user.FullName,
			Email: user.FullName,
			CreatedAt: user.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": response,
	})
}

func (h *UserHandler) FindByID(ctx *gin.Context) {
	// userID := ctx.MustGet("userID").(entity.UserID)
	// This should only be accessible by staff
	idStr := ctx.Param("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// work on error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := entity.UserID(id)
	user, err := h.userService.FindByID(userID)

	if err != nil {
		// work on error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error. Try again later",
		})
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
	userID := ctx.MustGet("userID").(entity.UserID)

	if err := h.userService.DeleteByID(userID); err != nil {
		// work on error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error. Try again later",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	var req UserUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// work on error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid inputs",
		})
		return
	}

	user := entity.User{
		FullName: req.FullName,
		Password: &req.Password,
	}

	result, err := h.userService.Update(&user);
	if err != nil {
		// work on error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error. Try again later",
		})
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
