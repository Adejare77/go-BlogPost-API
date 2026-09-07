package postgres

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
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
	if err := repo.db.Create(comment).Error; err != nil {
		return MapError(repo.db.Create(comment).Error)
	}

	return MapError(
		repo.db.Preload("Author").
		First(comment, "id = ?", comment.ID).Error,
	)
}

func (repo *CommentRepository) FindByID(commentID entity.CommentID, userID entity.UserID) (*comment.CommentDetailRow, error) {
	var commentDetail comment.CommentDetailRow

	if err := repo.db.Model(&entity.Comment{}).
	Select(`
		comments.id AS id,
		comments.author_id AS author_id,
		users.full_name AS full_name,
		comments.post_id AS post_id,
		comments.content AS content,
		comments.created_at AS created_at,

		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS likes,

		EXISTS (
			SELECT 1
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
			AND likes.user_id = ?
		) AS liked,

		(
			SELECT COUNT(*)
			FROM comments
			WHERE comments.parent_id IS NOT NULL
			AND comments.parent_id = comments.id
		) AS reply_count
	`, userID).
	Joins("JOIN users ON users.id = comments.author_id").
	Joins("JOIN posts ON posts.id = comments.post_id").
	Scan(&commentDetail).Error; err != nil {
		return nil, MapError(err)
	}


	var topRepliesRow []comment.CommentListRow

	if err := repo.db.Model(&entity.Comment{}).
	Select(`
		comments.id AS id,
		comments.author_id AS author_id,
		users.full_name AS full_name,
		posts.id AS post_id,
		comments.content AS content,
		comments.created_at AS created_at,
		comments.parent_id AS parent_id,

		(
			SELECT 1
			FROM likes
			WHERE likes.likeable_type = 'comment'
			AND likes.user_id = ?
			AND likes.likeable_id = comments.id
		) AS liked,

		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS likes,

		(
			SELECT COUNT(*)
			FROM comments
			WHERE parent_id = ?
		) AS reply_count

	`, userID, commentID).
	Joins("JOIN users ON users.id = comments.author_id").
	Joins("JOIN posts ON posts.id = comments.post_id").
	Where("parent_id = ?", commentID).
	Order("likes DESC, created_at DESC").
	Limit(3).
	Scan(&topRepliesRow).Error; err != nil {
		return nil, MapError(err)
	}

	return &commentDetail, nil
}

func (repo *CommentRepository) Update(comment *entity.Comment) (*comment.CommentDetailRow, error) {
	result := repo.db.Model(&entity.Comment{}).
	Where("id = ? AND author_id = ?", comment.ID, comment.AuthorID).
	Updates(comment)

	if result.Error != nil {
		return nil, MapError(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, domainErrors.ErrNotFound
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
		return domainErrors.ErrNotFound
	}

	return nil
}

func (repo *CommentRepository) FindByPostID(postID entity.PostID, userID entity.UserID) ([]comment.CommentListRow, error) {
	var comments []comment.CommentListRow

	if err := repo.db.Model(&entity.Comment{}).
	Select(`
		comments.id AS id,
		comments.author_id AS author_id,
		users.full_name AS full_name,
		comments.post_id AS post_id,
		comments.content AS content,
		comments.created_at AS created_at,

		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS likes,

		EXISTS (
			SELECT 1
			FROM likes
			WHERE likes.likeable_type = 'comment'
			AND likes.likeable_id = comments.id
			AND likes.user_id = ?
		) AS liked,

		(
			SELECT COUNT(*)
			FROM comments AS replies
			WHERE replies.parent_id IS NOT NULL
			AND replies.parent_id = comments.id
		) AS reply_count
	`, userID).
	Joins("JOIN users ON users.id = comments.author_id").
	Where("post_id = ? AND parent_id IS NULL", postID).
	Scan(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}
