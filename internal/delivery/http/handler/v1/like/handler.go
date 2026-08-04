package like

import (
	"net/http"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LikeHandler struct {
	likeService *usecase.LikeService
}

func (h *LikeHandler) CreateLikePost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(entity.UserID)
	postID := entity.LikeID(id)

	like := entity.Like{
		UserID: userID,
		LikeableID: postID,
		LikeableType: "post",
	}

	if err := h.likeService.Create(&like); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *LikeHandler) CreateLikeComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(entity.UserID)
	commentID := entity.LikeID(id)

	like := entity.Like{
		UserID: userID,
		LikeableID: commentID,
		LikeableType: "comment",
	}

	if err := h.likeService.Create(&like); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *LikeHandler) DeleteLikedPost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(entity.UserID)
	postID := entity.LikeID(id)

	if err := h.likeService.DeleteByUserAndPost(userID, postID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *LikeHandler) DeleteLikedComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(entity.UserID)
	commentID := entity.LikeID(id)

	if err := h.likeService.DeleteByUserAndPost(userID, commentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
