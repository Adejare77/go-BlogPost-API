package post

import (
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"github.com/Adejare77/go-BlogPost-API/internal/usecase/text"
)

type PostList struct {
	ID entity.PostID
	Author AuthorSummary
	Title string
	Excerpt string
	Likes int
	Liked bool
	IsPublished bool
	CommentCount int
	CreatedAt time.Time
}

func ToPostList(posts []post.PostListRow) []PostList {
	result := make([]PostList, len(posts))

	for i, post := range posts {
		result[i] = PostList{
			ID: post.ID,
			Author: AuthorSummary{
				ID: post.AuthorID,
				FullName: post.FullName,
			},
			Title: post.Title,
			Excerpt: text.ToExcerpt(post.Content),
			Likes: post.Likes,
			Liked: post.Liked,
			IsPublished: post.IsPublished,
			CommentCount: post.CommentCount,
			CreatedAt: post.CreatedAt,
		}
	}
	return result
}
