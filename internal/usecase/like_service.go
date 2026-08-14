package usecase

import (
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/like"
)

type LikeService struct {
	repo like.LikeRepository
}

func (s *LikeService) Create(like *entity.Like) error {
	return s.repo.Create(like)
}

func (s *LikeService) ExistByUserAndPost(userID entity.UserID, postID entity.LikeID) (bool, error) {
	return s.repo.ExistByUserAndPost(userID, postID)
}

func (s *LikeService) ExistByUserAndComment(userID entity.UserID, commentID entity.LikeID) (bool, error) {
	return s.repo.ExistByUserAndComment(userID, commentID)
}

func (s *LikeService) DeleteByUserAndPost(userID entity.UserID, postID entity.LikeID) error {
	return s.repo.DeleteByUserAndPost(userID, postID)
}

func (s *LikeService) DeleteByUserAndComment(userID entity.UserID, commentID entity.LikeID) error {
	return s.repo.DeleteByUserAndPost(userID, commentID)
}

func (s *LikeService) CountByPost(postID entity.LikeID) (int, error) {
	return s.repo.CountByPost(postID)
}

func (s *LikeService) CountByComment(commentID entity.LikeID) (int, error) {
	return s.repo.CountByComment(commentID)
}
