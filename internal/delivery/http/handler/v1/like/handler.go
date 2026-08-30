package like

import (
	"net/http"

	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LikeHandler struct {
	likeService *usecase.LikeService
}

func NewLikeHandler(likeService *usecase.LikeService) *LikeHandler {
	return &LikeHandler{
		likeService: likeService,
	}
}

func (h *LikeHandler) CreateLikePost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"post_id",
			"must be a valid UUID",
			err,
		)
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
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *LikeHandler) CreateLikeComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"comment_id",
			"must be a valid UUID",
			err,
		)
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
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *LikeHandler) DeleteLikedPost(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"post_id",
			"must be a valid UUID",
			err,
		)
		return
	}

	userID := ctx.MustGet("userID").(entity.UserID)
	postID := entity.LikeID(id)

	if err := h.likeService.DeleteByUserAndPost(userID, postID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *LikeHandler) DeleteLikedComment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"comment_id",
			"must be a valid UUID",
			err,
		)
		return
	}

	userID := ctx.MustGet("userID").(entity.UserID)
	commentID := entity.LikeID(id)

	if err := h.likeService.DeleteByUserAndPost(userID, commentID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
