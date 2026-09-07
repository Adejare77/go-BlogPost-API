package comment

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

func (s *CommentService) Create(comment *entity.Comment) (*CommentCreateResponse, error) {
	if err := s.repo.Create(comment); err != nil {
		return nil, err
	}

	response := CommentCreateResponse{
		ID: comment.ID,
		Author: AuthorSummary{
			ID: comment.AuthorID,
			FullName: comment.Author.FullName,
		},
		PostID: comment.PostID,
		ParentID: comment.ParentID,
		Content: comment.Content,
		CreatedAt: comment.CreatedAt,
	}

	return &response, nil
}

func (s *CommentService) FindByID(commentID entity.CommentID, userID entity.UserID) (*CommentDetail, error) {
	result, err := s.repo.FindByID(commentID, userID)
	if err != nil {
		return nil, err
	}

	return ToCommentDetail(result), nil
}

func (s *CommentService) Update(comment *entity.Comment) (*CommentDetail, error) {
	result, err := s.repo.Update(comment)
	if err != nil {
		return nil, err
	}

	return ToCommentDetail(result), nil
}

func (s *CommentService) DeleteByID(commentID entity.CommentID, userID entity.UserID) error {
	return s.repo.DeleteByID(commentID, userID)
}

func (s *CommentService) FindByPostID(postID entity.PostID, userID entity.UserID) ([]CommentList, error) {
	result, err := s.repo.FindByPostID(postID, userID)
	if err != nil {
		return nil, err
	}

	return ToCommentList(result), nil
}
