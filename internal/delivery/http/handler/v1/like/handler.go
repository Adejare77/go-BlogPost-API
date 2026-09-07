package like

import (
	"errors"
	"net/http"

	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase/like"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type LikeHandler struct {
	likeService *like.LikeService
}

func NewLikeHandler(likeService *like.LikeService) *LikeHandler {
	return &LikeHandler{
		likeService: likeService,
	}
}

func (h *LikeHandler) CreateLikePost(ctx *gin.Context) {
	var path LikeRequestPostPath

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, validationErrs) {
			httperrors.Validator(ctx, path, validationErrs)
			return
		}

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
	postID := entity.LikeID(uuid.MustParse(path.PostID))

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
	var path LikeRequestCommentPath

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, validationErrs) {
			httperrors.Validator(ctx, path, validationErrs)
			return
		}

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
	commentID := entity.LikeID(uuid.MustParse(path.CommentID))

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
	var path LikeRequestPostPath

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, validationErrs) {
			httperrors.Validator(ctx, path, validationErrs)
			return
		}

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
	postID := entity.LikeID(uuid.MustParse(path.PostID))

	if err := h.likeService.DeleteByUserAndPost(userID, postID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *LikeHandler) DeleteLikedComment(ctx *gin.Context) {
	var path LikeRequestCommentPath

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, validationErrs) {
			httperrors.Validator(ctx, path, validationErrs)
			return
		}

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
	commentID := entity.LikeID(uuid.MustParse(path.CommentID))

	if err := h.likeService.DeleteByUserAndPost(userID, commentID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
