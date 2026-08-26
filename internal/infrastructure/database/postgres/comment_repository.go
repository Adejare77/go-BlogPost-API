package postgres

import (
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
	return MapError(repo.db.Create(comment).Error)
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
		return nil, MapError(err)
	}


	var topRepliesRow []ReplySummaryRow

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
		) AS comment_count

	`, userID, commentID).
	Joins("JOIN users ON users.id = comments.author_id").
	Joins("JOIN posts ON posts.id = comments.post_id").
	Where("parent_id = ?", commentID).
	Order("likes DESC created_at DESC").
	Limit(3).
	Scan(&topRepliesRow).Error; err != nil {
		return nil, MapError(err)
	}

	topReplies := make([]comment.ReplySummary, len(topRepliesRow))

	for i, reply := range topRepliesRow {
		topReplies[i] = reply.ToReplySummary()
	}

	commentDetail.TopReplies = topReplies

	return commentDetail.ToCommentDetail(), nil
}

func (repo *CommentRepository) Update(comment *entity.Comment) (*comment.CommentDetail, error) {
	result := repo.db.Model(&entity.Comment{}).
	Where("id = ? AND author_id = ?", comment.ID, comment.AuthorID).
	Updates(comment)

	if result.Error != nil {
		return nil, MapError(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return repo.FindByID(comment.ID, comment.AuthorID)
}

func (repo *CommentRepository) DeleteByID(commentID entity.CommentID, userID entity.UserID) error {
	result:= repo.db.
	Where("id = ? AND author_id = ?", commentID, userID).
	Delete(&entity.Comment{})

	if result.Error != nil {
		return MapError(result.Error)
	}

	if result.RowsAffected == 0 {
		return MapError(result.Error)
	}

	return nil
}

func (repo *CommentRepository) FindByPostID(postID entity.PostID, userID entity.UserID) ([]comment.CommentList, error) {
	var list []CommentListRow

	err := repo.db.Model(&entity.Comment{}).
	Select(`
		comments.id AS id,
		comments.author_id AS author_id,
		users.full_name AS full_name,
		comments.post_id AS post_id,
		comments.content AS content
		comments.created_at AS created_at

		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS likes

		EXISTS (
			SELECT 1
			FROM likes
			WHERE likes.likeable_type = 'comment'
			AND likes.likeable_id = comments.id
			AND likes.user_id = ?
		) AS liked

		(
			SELECT COUNT(*)
			FROM comments AS replies
			WHERE replies.parent_id IS NOT NULL
			AND replies.parent_id = comments.id
		) AS reply_count
	`, userID).
	Joins("JOIN users ON users.id = comments.author_id").
	Where("post_id = ? AND parent_id IS NULL", postID).
	Scan(&list).Error

	if err != nil {
		return nil, MapError(err)
	}

	comments := make([]comment.CommentList, len(list))

	for i, comment := range list {
		comments[i] = comment.ToCommentList()
	}

	return comments, nil
}
