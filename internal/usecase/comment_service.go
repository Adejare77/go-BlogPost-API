package usecase

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
)

type CommentService struct {
	repo comment.CommentRepository
}

func NewCommentService(repo comment.CommentRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) Create(comment *entity.Comment) error {
	return s.repo.Create(comment)
}

func (s *CommentService) FindByID(commentID entity.CommentID, userID entity.UserID) (*comment.CommentDetail, error) {
	return s.repo.FindByID(commentID, userID)
}

func (s *CommentService) Update(comment *entity.Comment) (*comment.CommentDetail, error) {
	return s.repo.Update(comment)
}

func (s *CommentService) DeleteByID(commentID entity.CommentID, userID entity.UserID) error {
	return s.repo.DeleteByID(commentID, userID)
}

func (s *CommentService) FindByPostID(postID entity.PostID, userID entity.UserID) ([]comment.CommentList, error) {
	return s.repo.FindByPostID(postID, userID)
}
