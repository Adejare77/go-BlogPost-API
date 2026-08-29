package postgres

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (repo *PostRepository) Create(post *entity.Post) error {
	return MapError(repo.db.Create(post).Error)
}

func (repo *PostRepository) FindByID(postID entity.PostID, userID entity.UserID) (*post.PostDetail, error) {
	var detail PostDetailRow

	err := repo.db.Model(&entity.Post{}).
	Select(`
		posts.id AS id,
		posts.author_id AS author_id,
		users.full_name AS full_name,
		posts.title AS title,
		posts.content AS content,
		posts.is_published AS is_published,
		posts.created_at,

		(
			SELECT COUNT(*),
			FROM likes
			WHERE likes.likeable_id = posts.id
			AND likes.likeable_type = 'post'
		) AS likes,

		(
			SELECT COUNT(*),
			FROM comments
			WHERE comments.post_id = posts.id
			AND comments.parent_id IS NULL
		) AS comment_count,

		EXISTS (
			SELECT 1
			FROM likes
			WHERE likes.likeable_id = posts.id
			AND likes.likeable_type = 'post'
			AND likes.user_id = ?
		) AS liked
	`, userID).
	Joins("JOIN users ON users.id = posts.author_id").
	Where("posts.id = ?", postID).
	Scan(&detail).Error

	if err != nil {
		return nil, MapError(err)
	}

	var commentlist []CommentListRow

	err = repo.db.Model(&entity.Comment{}).
	Select(`
		comments.id AS id,
		comments.author_id AS author_id,
		users.full_name AS full_name,
		comments.post_id,
		comments.content AS content,

		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS likes

		(
			SELECT COUNT(*)
			FROM comments AS replies
			WHERE replies.parent_id = comments.id
		) AS reply_count

		EXISTS (
			SELECT 1
			FROM likes
			WHERE likes.likeable_id = comments.id
			AND likes.likeable_type = 'comment'
		) AS liked
	`).
	Joins("JOIN users ON users.id = comments.author_id").
	Where("comments.post_id = ?", postID).
	Where("comments.parent_id IS NULL").
	Order("likes DESC").
	Limit(3).
	Scan(&commentlist).Error

	if err != nil {
		return nil, MapError(err)
	}

	comments := make([]comment.CommentList, len(commentlist))
	for i, comment := range commentlist {
		comments[i] = comment.ToCommentList()
	}

	detail.TopComments = comments

	return detail.ToPostDetail(), nil
}


func (repo *PostRepository) Update(post *entity.Post) (*post.PostDetail, error) {
	result := repo.db.Model(&entity.Post{}).
	Where("id = ? AND author_id = ?", post.ID, post.AuthorID).Updates(post)

	if result.Error != nil {
		return nil, MapError(result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, domainErrors.ErrNotFound
	}

	return repo.FindByID(post.ID, post.AuthorID)
}

func (repo *PostRepository) DeleteByID(postID entity.PostID, userID entity.UserID) error {
	result := repo.db.
	Where("id = ? AND author_id = ?", postID, userID).
	Delete(&entity.Post{})

	if result.Error != nil {
		return MapError(result.Error)
	}

	if result.RowsAffected == 0 {
		return domainErrors.ErrNotFound
	}

	return nil
}

func (repo *PostRepository) FindAll(userID entity.UserID) ([]post.PostList, error) {
	var list []PostListRow

	if err := repo.db.Table("posts").
	Select(`
		posts.id AS id,
		users.id AS author_id,
		users.full_name AS full_name,
		posts.title AS title,
		posts.content AS content,
		posts.is_published,
		posts.created_at AS created_at,
		(
			SELECT COUNT(*)
			FROM likes
			WHERE likes.likeable_id = posts.id
			AND likes.likeable_type = 'post'
		) AS likes,

		(
			SELECT COUNT(*)
			FROM comments
			WHERE comments.post_id = posts.id
			AND comments.parent_id IS NULL
		) AS comment_count,

		EXISTS (
			SELECT 1
			FROM likes
			WHERE likes.user_id = ?
			AND likes.likeable_type = 'post'
			AND likes.likeable_id = posts.id
		) AS liked
	`, userID).
	Joins("JOIN users ON users.id = posts.author_id").
	Order("posts.created_at DESC, likes DESC").
	Scan(&list).Error; err != nil {
		return nil, MapError(err)
	}

	posts := make([]post.PostList, len(list))

	for i, p := range list {
		posts[i] = p.ToPostList()
	}

	return posts, nil
}
