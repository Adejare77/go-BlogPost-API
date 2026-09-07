package post

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

func (s *PostService) Create(post *entity.Post) (*PostCreateResponse, error) {
	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	response := &PostCreateResponse{
		ID: post.ID,
		Author: AuthorSummary{
			ID: post.AuthorID,
			FullName: post.Author.FullName,
		},
		Title: post.Title,
		Content: post.Content,
		IsPublished: post.IsPublished,
		CreatedAt: post.CreatedAt,
	}

	return response, nil
}

func (s *PostService) FindByID(postID entity.PostID, userID entity.UserID, isStaff bool) (*PostDetail, error) {
	post, err := s.repo.FindByID(postID, userID)
	if err != nil {
		return nil, err
	}

	if post.IsPublished || post.AuthorID == userID || isStaff {
		return ToPostDetail(post), err
	}

	return nil, domainErrors.ErrNotFound
}

func (s *PostService) Update(post *entity.Post) (*PostDetail, error) {
	result, err := s.repo.Update(post)
	if err != nil {
		return nil, err
	}

	return ToPostDetail(result), nil
}

func (s *PostService) DeleteByID(postID entity.PostID, userID entity.UserID) error {
	return s.repo.DeleteByID(postID, userID)
}

func (s *PostService) FindAll(userID entity.UserID, query post.PostQuery, isStaff bool) ([]PostList, error) {
	if query.Status == "" {
		query.Status = "published"
	}

	if userID == 0 && (query.Status == "draft" || query.Status == "all") {
		return nil, domainErrors.ErrUnauthorized
	}

	if query.Status != "published" && query.Author != "me" && query.Author == "" && !isStaff {
		return nil, domainErrors.ErrForbidden
	}

	result, err := s.repo.FindAll(userID, query)
	if err != nil {
		return nil, err
	}
	return ToPostList(result), nil
}
