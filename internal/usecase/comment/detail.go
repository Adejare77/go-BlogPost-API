package comment

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)


type AuthorSummary struct {
	ID entity.UserID
	FullName string
}

type CommentDetail struct {
	ID entity.CommentID
	Author AuthorSummary
	PostID entity.PostID
	ParentID *entity.CommentID
	Content string
	Likes int
	Liked bool
	ReplyCount int
	TopReplies []CommentList
	CreatedAt time.Time
}

func ToCommentDetail(comment *comment.CommentDetailRow) *CommentDetail {
	return &CommentDetail{
		ID: comment.ID,
		Author: AuthorSummary{
			ID: comment.AuthorID,
			FullName: comment.FullName,
		},
		PostID: comment.PostID,
		ParentID: comment.ParentID,
		Content: comment.Content,
		Likes: comment.Likes,
		Liked: comment.Liked,
		ReplyCount: comment.ReplyCount,
		TopReplies: ToCommentList(comment.TopReplies),
		CreatedAt: comment.CreatedAt,
	}
}

// type ReplySummary struct {
// 	ID entity.CommentID
// 	Author AuthorSummary
// 	PostID entity.PostID
// 	ParentID entity.CommentID
// 	Excerpt string
// 	Likes int
// 	Liked bool
// 	ReplyCount int
// 	CreatedAt time.Time
// }
