package post

import (
	"errors"
	"net/http"

	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainPost "github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	usecasePost "github.com/Adejare77/go-BlogPost-API/internal/usecase/post"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)


type PostHandler struct {
	postService *usecasePost.PostService
}

func NewPostHandler(postService *usecasePost.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) Create(ctx *gin.Context) {
	var req PostCreateRequest

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

	postRequest := entity.Post{
		AuthorID: userID,
		Title: req.Title,
		Content: req.Content,
		IsPublished: req.IsPublished,
	}

	response, err := h.postService.Create(&postRequest)
	if  err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *PostHandler) FindByID(ctx *gin.Context) {
	var path PostPathRequest

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
			"invalid postId",
			err,
		)
		return
	}

	postID := entity.PostID(uuid.MustParse(path.PostID))

	var userID entity.UserID
	isStaff := false

	value, exist := ctx.Get("UserID")
	if exist {
		userID = value.(entity.UserID)
		isStaff = ctx.MustGet("isStaff").(bool)
	}

	response, err := h.postService.FindByID(postID, userID, isStaff)

	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


func (h *PostHandler) Update(ctx *gin.Context) {
	var req PostUpdateRequest

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

	var path PostPathRequest

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
			"invalid path UUID",
			err,
		)
		return
	}

	postID := entity.PostID(uuid.MustParse(path.PostID))

	userID := ctx.MustGet("userID").(entity.UserID)

	post := entity.Post{
		ID: postID,
		AuthorID: userID,
		Title: req.Title,
		Content: req.Content,
		IsPublished: req.IsPublished,
	}

	response, err := h.postService.Update(&post)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": response,
	})
}

func (h *PostHandler) DeleteByID(ctx *gin.Context) {
	var path PostPathRequest

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
			"invalid path UUID",
			err,
		)
		return
	}

	postID := entity.PostID(uuid.MustParse(path.PostID))
	userID := ctx.MustGet("userID").(entity.UserID)

	if err := h.postService.DeleteByID(postID, userID); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)

}

func (h *PostHandler) FindAll(ctx *gin.Context) {
	var userID entity.UserID
	var reqQuery PostQueryRequest

	if err := ctx.ShouldBindUri(&reqQuery); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, reqQuery, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"invalid_queries",
			"invalid queries",
			err,
		)
		return
	}

	isStaff := false
	value, exist := ctx.Get("UserID")
	if exist {
		userID = value.(entity.UserID)
		isStaff = ctx.MustGet("isStaff").(bool)
	}

	query := domainPost.PostQuery{
		Status: reqQuery.Status,
		Author: reqQuery.Author,
	}

	response, err := h.postService.FindAll(userID, query, isStaff)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"results": response,
	})
}
