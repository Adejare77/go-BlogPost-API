package post

import (
	"errors"
	"net/http"

	httperrors "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type PostHandler struct {
	postService *usecase.PostService
}

func NewPostHandler(postService *usecase.PostService) *PostHandler {
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

	post := entity.Post{
		AuthorID: userID,
		Title: req.Title,
		Content: req.Content,
		IsPublished: req.IsPublished,
	}

	if err := h.postService.Create(&post); err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	response := PostCreateResponse{
		ID: post.ID,
		Author: AuthorSummary{
			ID: post.Author.ID,
			FullName: post.Author.FullName,
		},
		Title: post.Title,
		Content: post.Content,
		IsPublished: post.IsPublished,
		CreatedAt: post.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"results": response,
	})
}

func (h *PostHandler) FindByID(ctx *gin.Context) {
	var req PostPathRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, req, validationErrs)
			return
		}

		httperrors.HandleRequestError(
			ctx,
			http.StatusBadRequest,
			"post_id",
			"invalid postId",
			err,
		)
	}

	postID := entity.PostID(req.PostID)

	var userID entity.UserID
	isStaff := false

	value, exist := ctx.Get("UserID")
	if exist {
		userID = value.(entity.UserID)
		isStaff = ctx.MustGet("isStaff").(bool)
	}

	post, err := h.postService.FindByID(postID, userID, isStaff)

	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	topComments := make([]CommentSummary, len(post.TopComments))

	for i, c := range post.TopComments {
		topComments[i] = CommentSummary{
			ID: c.ID,
			Author: AuthorSummary{
				ID: c.Author.ID,
				FullName: c.Author.FullName,
			},
			PostID: c.PostID,
			Excerpt: c.Excerpt,
			Likes: c.Likes,
			Liked: c.Liked,
			ReplyCount: c.ReplyCount,
			CreatedAt: c.CreatedAt,
		}
	}

	response := PostDetailResponse{
		ID: post.ID,
		Author: AuthorSummary{
			ID: post.Author.ID,
			FullName: post.Author.FullName,
		},
		Title: post.Title,
		Content: post.Content,
		Likes: post.Likes,
		Liked: post.Liked,
		IsPublished: post.IsPublished,
		CommentCount: post.CommentCount,
		TopComments: topComments,
		CreatedAt: post.CreatedAt,
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

	postID := entity.PostID(path.PostID)
	userID := ctx.MustGet("userID").(entity.UserID)

	post := entity.Post{
		ID: postID,
		AuthorID: userID,
		Title: req.Title,
		Content: req.Content,
		IsPublished: req.IsPublished,
	}

	updatedPost, err := h.postService.Update(&post)
	if err != nil {
		httperrors.HandleError(ctx, err)
		return
	}

	topComments := make([]CommentSummary, len(updatedPost.TopComments))

	for i, c := range updatedPost.TopComments {
		topComments[i] = CommentSummary{
			ID: c.ID,
			Author: AuthorSummary{
				ID: c.Author.ID,
				FullName: c.Author.FullName,
			},
			Excerpt: c.Excerpt,
			Likes: c.Likes,
			Liked: c.Liked,
			CreatedAt: c.CreatedAt,

		}
	}

	response := PostDetailResponse{
		ID: updatedPost.ID,
		Author: AuthorSummary{
			ID: updatedPost.Author.ID,
			FullName: updatedPost.Author.FullName,
		},
		Title: updatedPost.Title,
		Content: updatedPost.Content,
		Likes: updatedPost.Likes,
		Liked: updatedPost.Liked,
		IsPublished: updatedPost.IsPublished,
		TopComments: topComments,
		CreatedAt: updatedPost.CreatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": response,
	})
}

func (h *PostHandler) DeleteByID(ctx *gin.Context) {
	var req PostPathRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			httperrors.Validator(ctx, req, validationErrs)
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

	postID := entity.PostID(req.PostID)
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

	query := post.PostQuery{
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
