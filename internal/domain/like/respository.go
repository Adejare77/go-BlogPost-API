package like

import "github.com/Adejare77/go-BlogPost-API/internal/domain/entity"

type LikeRepository interface {
	Create(like *entity.Like) error
	DeleteByUserAndPost(userID entity.UserID, postID entity.LikeID) error
	DeleteByUserAndComment(userID entity.UserID, commentID entity.LikeID) error
}
