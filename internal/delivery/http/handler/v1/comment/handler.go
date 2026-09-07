package comment

import (
	"errors"
	"net/http"

	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/http/handler/v1/post"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase/comment"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)


type CommentHandler struct {
	commentService *comment.CommentService
}

func NewCommentHandler(commentService *comment.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) CreateComment(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(entity.UserID)

	var path post.PostPathRequest
	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
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

	postID := entity.PostID(uuid.MustParse(path.PostID))

	var req CommentRequest
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
			"invalid_request",
			err,
		)
		return
	}

	comment := entity.Comment{
		AuthorID: userID,
		PostID: postID,
		Content: req.Content,
	}

	response, err := h.commentService.Create(&comment)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *CommentHandler) CreateReply(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(entity.UserID)

	var req CommentRequest
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

	var path CommentPathRequest

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
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

	parentID := entity.CommentID(uuid.MustParse(path.CommentID))

	comment := entity.Comment{
		AuthorID: userID,
		ParentID: &parentID,
		Content: req.Content,
	}

	response, err := h.commentService.Create(&comment)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *CommentHandler) FindByID(ctx *gin.Context) {
	var userID entity.UserID
	var path CommentPathRequest

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
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

	commentID := entity.CommentID(uuid.MustParse(path.CommentID))

	if value, exist := ctx.Get("userID"); exist {
		userID = value.(entity.UserID)
	}

	response, err := h.commentService.FindByID(commentID, userID)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": response,
	})
}

func (h *CommentHandler) Update(ctx *gin.Context) {
	var req CommentRequest

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

	var path CommentPathRequest

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
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
	commentID := entity.CommentID(uuid.MustParse(path.CommentID))

	comment := entity.Comment {
		ID: commentID,
		AuthorID: userID,
		Content: req.Content,
	}

	response, err := h.commentService.Update(&comment)

	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


func (h *CommentHandler) DeleteByID(ctx *gin.Context) {
	var path CommentPathRequest

	if err := ctx.ShouldBindUri(&path); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
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

	commentID := entity.CommentID(uuid.MustParse(path.CommentID))
	userID := ctx.MustGet("userID").(entity.UserID)

	if err := h.commentService.DeleteByID(commentID, userID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *CommentHandler) FindByPostID(ctx *gin.Context) {
	var path post.PostPathRequest

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

	var userID entity.UserID

	postID := entity.PostID(uuid.MustParse(path.PostID))
	if user, exists := ctx.Get("userID"); exists {
		userID = user.(entity.UserID)
	}


	response, err := h.commentService.FindByPostID(postID, userID)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
