package postgres

import (
	"fmt"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (repo *CommentRepository) Create(comment *entity.Comment) error {
	return repo.db.Create(comment).Error
}

func (repo *CommentRepository) FindByID(commentID entity.CommentID, userID entity.UserID) (*comment.CommentDetail, error) {
	var commentDetail CommentDetailRow

	if err := repo.db.Model(&entity.Comment{}).
	Select(`
		comments.id AS id,
		comments.author_id AS author_id
		users.full_name AS full_name
		comments.content AS content
		comments.created_at AS

		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS likes

		EXISTS (
			SELECT 1
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
			AND likes.user_id = ?
		) AS liked

		(
			SELECT COUNT(*)
			FROM comments
			WHERE comments.parent_id IS NOT NULL
			AND comments.parent_id = comments.id
		)
	`, userID).
	Joins("JOIN users ON users.id = comments.author_id").
	Joins("JOIN posts ON posts.id = comments.post_id").
	Scan(&commentDetail).Error; err != nil {
		return nil, fmt.Errorf("error fetching comment with ID %s: %w", commentID, err)
	}


	var topReplies []ReplySummaryRow

	if err := repo.db.Model(&entity.Comment{}).
	Select(`
		comments.id AS id,
		comments.author_id AS author_id,
		users.full_name AS full_name,
		posts.id AS post_id,
		comments.content AS content,

		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS likes

		(
			SELECT 1
			FROM likes
			WHERE likes.likeable_type = 'comment'
			AND likes.user_id = ?
			AND likes.likeable_id = comments.id
		) AS liked

		(
			SELECT COUNT(*)
			FROM comments
			WHERE parent_id = ?
		)

	`, userID, commentID).
	Joins("JOIN users ON users.id = comments.author_id").
	Joins("JOIN posts ON posts.id = comments.post_id").
	Where("parent_id = ?", commentID).
	Order("likes DESC created_at DESC").
	Limit(3).
	Scan(&topReplies).Error; err != nil {
		return nil, fmt.Errorf("error fetching comment with id %s replies: %w", commentID, err)
	}

	commentDetail.TopReplies = topReplies

	// convert to commentDetail
	return &CommentDetailRow, nil
}

func (repo *CommentRepository) Update(comment *entity.CommentID) (*comment.CommentDetail, error) {

}
