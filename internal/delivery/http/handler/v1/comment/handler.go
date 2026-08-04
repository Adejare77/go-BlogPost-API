package comment

import (
	"net/http"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type CommentHandler struct {
	commentService *usecase.CommentService
}

func (h *CommentHandler) CreateComment(ctx *gin.Context) {
	var req CommentRequest
	userID := ctx.MustGet("userID").(entity.UserID)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	id, err := uuid.Parse(ctx.Param("post_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	postID := entity.PostID(id)

	comment := entity.Comment{
		AuthorID: userID,
		PostID: postID,
		Content: req.Content,
	}

	if err := h.commentService.Create(&comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	parentID := entity.CommentID(id)

	comment := entity.Comment{
		AuthorID: userID,
		ParentID: &parentID,
		Content: req.Content,
	}

	if err := h.commentService.Create(&comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
	var userID *entity.UserID

	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	commentID := entity.CommentID(id)

	if v, ok := ctx.Get("userID"); ok {
		userID = v.(*entity.UserID)
	}

	comment, err := h.commentService.FindByID(commentID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(entity.UserID)
	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}


func (h *CommentHandler) DeleteByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("comment_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	commentID := entity.CommentID(id)
	userID := ctx.MustGet("userID").(entity.UserID)

	if err := h.commentService.DeleteByID(commentID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
