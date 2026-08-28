package comment

import (
	"errors"
	"net/http"

	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)


type CommentHandler struct {
	commentService *usecase.CommentService
}

func NewCommentHandler(commentService *usecase.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) CreateComment(ctx *gin.Context) {
	var req CommentRequest
	userID := ctx.MustGet("userID").(entity.UserID)

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

	postID := entity.PostID(id)

	comment := entity.Comment{
		AuthorID: userID,
		PostID: postID,
		Content: req.Content,
	}

	if err := h.commentService.Create(&comment); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	response := CommentCreateResponse {
		ID: comment.ID,
		Author: AuthorSummary{
			ID: comment.Author.ID,
			FullName: comment.Author.FullName,
		},
		PostID: comment.PostID,
		Content: comment.Content,
		CreatedAt: comment.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *CommentHandler) CreateReply(ctx *gin.Context) {
	var req CommentRequest
	userID := ctx.MustGet("userID").(entity.UserID)

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

	parentID := entity.CommentID(id)

	comment := entity.Comment{
		AuthorID: userID,
		ParentID: &parentID,
		Content: req.Content,
	}

	if err := h.commentService.Create(&comment); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	response := CommentCreateResponse {
		ID: comment.ID,
		Author: AuthorSummary{
			ID: comment.Author.ID,
			FullName: comment.Author.FullName,
		},
		Content: comment.Content,
		CreatedAt: comment.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *CommentHandler) FindByID(ctx *gin.Context) {
	var userID entity.UserID

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

	commentID := entity.CommentID(id)

	if v, ok := ctx.Get("userID"); ok {
		userID = v.(entity.UserID)
	}

	comment, err := h.commentService.FindByID(commentID, userID)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	topReplies := make([]CommentListResponse, len(comment.TopReplies) )
	for i, v := range comment.TopReplies {
		topReplies[i] = CommentListResponse{
			ID: v.ID,
			Author: AuthorSummary{
				v.Author.ID,
				v.Author.FullName,
			},
			PostID: v.PostID,
			Excerpt: v.Excerpt,
			Likes: v.Likes,
			ReplyCount: v.ReplyCount,
			CreatedAt: v.CreatedAt,
		}
	}

	response := CommentDetailResponse{
		ID: comment.ID,
		Author: AuthorSummary{
			ID: comment.Author.ID,
			FullName: comment.Author.FullName,
		},
		Content: comment.Content,
		Likes: comment.Likes,
		Liked: comment.Liked,
		ReplyCount: comment.ReplyCount,
		TopReplies: topReplies,
		CreatedAt: comment.CreatedAt,
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

	userID := ctx.MustGet("userID").(entity.UserID)

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

	commentID := entity.CommentID(id)

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

	commentID := entity.CommentID(id)
	userID := ctx.MustGet("userID").(entity.UserID)

	if err := h.commentService.DeleteByID(commentID, userID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
