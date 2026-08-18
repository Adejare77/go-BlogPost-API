package usecase

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
)

type PostService struct {
	repo post.PostRepository
}

func NewPostService(repo post.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) Create(post *entity.Post) error {
	return s.repo.Create(post)
}

func (s *PostService) FindByID(postID entity.PostID, userID entity.UserID) (*post.PostDetail, error) {
	return s.repo.FindByID(postID, userID)
}

func (s *PostService) Update(post *entity.Post) (*post.PostDetail, error) {
	return s.repo.Update(post)
}

func (s *PostService) DeleteByID(postID entity.PostID, userID entity.UserID) error {
	return s.repo.DeleteByID(postID, userID)
}

func (s *PostService) FindAll(userID entity.UserID) ([]post.PostList, error) {
	return s.repo.FindAll(userID)
}

// func (s *PostService) FindPostDetail(postID entity.PostID, userID entity.UserID) ([]*entity.PostDetail, error) {
// 	return s.repo.FindPostDetail(postID, userID)
// }
