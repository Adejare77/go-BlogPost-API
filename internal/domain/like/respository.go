package like

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type LikeRepository interface {
	Create(like *entity.Like) error
	ExistByUserAndPost(userID entity.UserID, postID entity.LikeID) (bool, error)
	ExistByUserAndComment(userID entity.UserID, commentID entity.LikeID) (bool, error)
	DeleteByUserAndPost(userID entity.UserID, postID entity.LikeID) error
	DeleteByUserAndComment(userID entity.UserID, commentID entity.LikeID) error
	CountByPost(postID entity.LikeID) (int, error)
	CountByComment(commentID entity.LikeID) (int, error)
}
