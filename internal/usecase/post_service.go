package usecase

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
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

func (s *PostService) FindByID(postID entity.PostID, userID entity.UserID, isStaff bool) (*post.PostDetail, error) {
	post, err := s.repo.FindByID(postID, userID)
	if err != nil {
		return nil, err
	}

	if post.IsPublished || post.Author.ID == userID || isStaff {
		return post, err
	}

	return nil, domainErrors.ErrNotFound
}

func (s *PostService) Update(post *entity.Post) (*post.PostDetail, error) {
	return s.repo.Update(post)
}

func (s *PostService) DeleteByID(postID entity.PostID, userID entity.UserID) error {
	return s.repo.DeleteByID(postID, userID)
}

func (s *PostService) FindAll(userID entity.UserID, query post.PostQuery, isStaff bool) ([]post.PostList, error) {
	if userID == 0 && (query.Status == "draft" || query.Status == "all") {
		return []post.PostList{}, domainErrors.ErrUnauthorized
	}

	if query.Status != "published" && query.Author != "me" && query.Author == "" && !isStaff {
		return []post.PostList{}, domainErrors.ErrForbidden
	}

	return s.repo.FindAll(userID, query)
}
