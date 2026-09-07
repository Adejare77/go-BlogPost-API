package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase/text"
)


type CommentList struct {
	ID entity.CommentID
	Author AuthorSummary
	FullName string
	PostID entity.PostID
	Excerpt string
	Likes int
	Liked bool
	ReplyCount int
	CreatedAt time.Time
}

func ToCommentList(comments []comment.CommentListRow) []CommentList {
	result := make([]CommentList, len(comments))

	for i, comment := range comments {
		result[i] = CommentList{
			ID: comment.ID,
			Author: AuthorSummary{
				ID: comment.AuthorID,
				FullName: comment.FullName,
			},
			PostID: comment.PostID,
			Excerpt: text.ToExcerpt(comment.Content),
			Likes: comment.Likes,
			Liked: comment.Liked,
			ReplyCount: comment.ReplyCount,
			CreatedAt: comment.CreatedAt,
		}
	}

	return result
}
