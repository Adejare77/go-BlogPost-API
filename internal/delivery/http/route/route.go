package route

import (
	httpauth "github.com/Adejare77/go-BlogPost-API/internal/delivery/http/handler/v1/auth"
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/http/handler/v1/comment"
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/http/handler/v1/like"
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/http/handler/v1/post"
	"github.com/Adejare77/go-BlogPost-API/internal/delivery/http/handler/v1/user"
	domainauth "github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/gin-gonic/gin"
)


func Setup(
	r *gin.Engine,
	userHandler *user.UserHandler,
	postHandler *post.PostHandler,
	commentHandler *comment.CommentHandler,
	likeHandler *like.LikeHandler,
	authHandler *httpauth.AuthHandler,
	tokenService domainauth.TokenService,
	) {
		api := r.Group("/api")

		v1 := api.Group("/v1")

		user.RegisterRoutes(v1, userHandler, tokenService)
		post.RegisterRoutes(v1, postHandler, tokenService)
		comment.RegisterRoutes(v1, commentHandler, tokenService)
		like.RegisterRoutes(v1, likeHandler, tokenService)
		httpauth.RegisterRoutes(v1, authHandler, tokenService)

}
