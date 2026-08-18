package usecase

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/like"
)

type LikeService struct {
	repo like.LikeRepository
}

func NewLikeService(repo like.LikeRepository) *LikeService {
	return &LikeService{
		repo: repo,
	}
}

func (s *LikeService) Create(like *entity.Like) error {
	return s.repo.Create(like)
}

func (s *LikeService) DeleteByUserAndPost(userID entity.UserID, postID entity.LikeID) error {
	return s.repo.DeleteByUserAndPost(userID, postID)
}

func (s *LikeService) DeleteByUserAndComment(userID entity.UserID, commentID entity.LikeID) error {
	return s.repo.DeleteByUserAndPost(userID, commentID)
}
