package post

import (
	"errors"
	"testing"
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/post"
	"github.com/Adejare77/go-BlogPost-API/internal/mocks"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)


func TestPostService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)

	service := NewPostService(mockRepo)

	userID := entity.UserID(1)
	FullName := "ABC"

	post := &entity.Post{
		Title: "Sample 1",
		Content: "This is sample 1 testing",
		IsPublished: true,
		AuthorID: userID,
	}

	mockRepo.EXPECT().
	Create(post).
	DoAndReturn(func(p *entity.Post) error {
		p.ID = entity.PostID(uuid.New())
		p.IsPublished = post.IsPublished
		p.AuthorID = post.AuthorID
		p.CreatedAt = time.Now()
		p.Author.FullName = FullName

		return nil
	})

	response, err := service.Create(post)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID == (entity.PostID{}) {
		t.Error("expected a response ID")
	}

	if response.Title != post.Title {
		t.Errorf("got title %q, want %q\n", response.Title, post.Title)
	}

	if response.Content != post.Content {
		t.Errorf("got content %q, want %q\n", response.Content, post.Content)
	}

	if response.IsPublished != post.IsPublished {
		t.Errorf("got IsPublished %v, want %v\n", response.IsPublished, post.IsPublished)
	}

	if response.Author.ID != userID {
		t.Errorf("got author.id %d, want %d\n", response.Author.ID, userID)
	}

	if response.Author.FullName != FullName {
		t.Errorf("got author.full_name %s, want %s\n", response.Author.FullName, FullName)
	}

}


func TestPostService_Create_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	userID := entity.UserID(3)

	post := &entity.Post{
		AuthorID: userID,
		Title: "This is a sample title",
		Content: "This is a sample content",
	}

	expectedErr := errors.New("repository error")

	mockRepo.EXPECT().
	Create(post).
	Return(expectedErr)

	response, err := service.Create(post)

	if response != nil {
		t.Fatal("expected response to be nil")
	}

	if err == nil {
		t.Fatal("expected error not be be nil")
	}

	if err != expectedErr {
		t.Errorf("got err %v, want %v\n", err, expectedErr)
	}
}


func TestPostService_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	postID := entity.PostID(uuid.New())
	userID := entity.UserID(5)

	post := &post.PostDetailRow{
		ID: postID,
		AuthorID: 2,
		FullName: "John Roe",
		Title: "This is a sample title",
		Content: "This is a sample content",
		Likes: 3,
		Liked: false,
		IsPublished: true,
		CommentCount: 5,
		CreatedAt: time.Now().Add(-15 *time.Hour),
	}

	mockRepo.EXPECT().
	FindByID(postID, userID).
	Return(post, nil)

	response, err := service.FindByID(postID, userID, false)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID != postID {
		t.Errorf("got response.ID %d, want %d\n", response.ID, postID)
	}

	if response.Author.ID != post.AuthorID {
		t.Errorf("got response.Author.ID %d, want %d\n", response.Author.ID, post.ID)
	}

	if response.Author.FullName != post.FullName {
		t.Errorf("got response.Author.FullName %s, want %s\n", response.Author.FullName, post.FullName)
	}

	if response.Title != post.Title {
		t.Errorf("got response.Title %s, want %s\n", response.Title, post.Title)
	}

	if response.Content != post.Content {
		t.Errorf("got response.Content %s, want %s\n", response.Content, post.Content)
	}

	if response.Likes != post.Likes {
		t.Errorf("got response.likes %d, want %d\n", response.Likes, post.Likes)
	}

	if response.Liked != post.Liked {
		t.Errorf("got response.Liked %v, want %v\n", response.Liked, post.Liked)
	}

	if response.IsPublished != post.IsPublished {
		t.Errorf("got response.IsPublished %v, want %v\n", response.IsPublished, post.IsPublished)
	}

	if response.CommentCount != post.CommentCount {
		t.Errorf("got response.CommentCount %d, want %d\n", response.CommentCount, post.CommentCount)
	}

	if len(response.TopComments) != len(post.TopComments) {
		t.Errorf("got response.TopComments %v, want %v\n", len(response.TopComments), len(post.TopComments))
	}

	if response.CreatedAt != post.CreatedAt {
		t.Errorf("got response.CreatedAt %v, want %v\n", response.CreatedAt, post.CreatedAt)
	}
}


func TestPostService_FindByID_ErrNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	userID := entity.UserID(0)
	postID := entity.PostID(uuid.New())

	mockRepo.EXPECT().
	FindByID(postID, userID).
	Return(nil, domainErrors.ErrNotFound)

	response, err := service.FindByID(postID, userID, false)
	if response != nil {
		t.Fatalf("expected no response, got %v", response)
	}

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != domainErrors.ErrNotFound {
		t.Errorf("got errr %v, want %v", err, domainErrors.ErrNotFound)
	}
}

func TestPostService_FindByID_Draft_Owner(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	postID := entity.PostID(uuid.New())
	userID := entity.UserID(5)

	post := &post.PostDetailRow{
		ID: postID,
		AuthorID: userID,
		FullName: "John Roe",
		Title: "This is a sample title",
		Content: "This is a sample content",
		Likes: 3,
		Liked: false,
		IsPublished: false,
		CommentCount: 5,
		CreatedAt: time.Now().Add(-15 *time.Hour),
	}

	mockRepo.EXPECT().
	FindByID(postID, userID).
	Return(post, nil)

	response, err := service.FindByID(postID, userID, false)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID != postID {
		t.Errorf("got response.ID %d, want %d\n", response.ID, postID)
	}

	if response.Author.ID != post.AuthorID {
		t.Errorf("got response.Author.ID %d, want %d\n", response.Author.ID, post.ID)
	}

	if response.Author.FullName != post.FullName {
		t.Errorf("got response.Author.FullName %s, want %s\n", response.Author.FullName, post.FullName)
	}

	if response.Title != post.Title {
		t.Errorf("got response.Title %s, want %s\n", response.Title, post.Title)
	}

	if response.Content != post.Content {
		t.Errorf("got response.Content %s, want %s\n", response.Content, post.Content)
	}

	if response.Likes != post.Likes {
		t.Errorf("got response.likes %d, want %d\n", response.Likes, post.Likes)
	}

	if response.Liked != post.Liked {
		t.Errorf("got response.Liked %v, want %v\n", response.Liked, post.Liked)
	}

	if response.IsPublished != post.IsPublished {
		t.Errorf("got response.IsPublished %v, want %v\n", response.IsPublished, post.IsPublished)
	}

	if response.CommentCount != post.CommentCount {
		t.Errorf("got response.CommentCount %d, want %d\n", response.CommentCount, post.CommentCount)
	}

	if len(response.TopComments) != len(post.TopComments) {
		t.Errorf("got response.TopComments %v, want %v\n", len(response.TopComments), len(post.TopComments))
	}

	if response.CreatedAt != post.CreatedAt {
		t.Errorf("got response.CreatedAt %v, want %v\n", response.CreatedAt, post.CreatedAt)
	}
}


func TestPostService_FindByID_Draft_Staff(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	postID := entity.PostID(uuid.New())
	userID := entity.UserID(5)

	post := &post.PostDetailRow{
		ID: postID,
		AuthorID: 2,
		FullName: "John Roe",
		Title: "This is a sample title",
		Content: "This is a sample content",
		Likes: 3,
		Liked: false,
		IsPublished: false,
		CommentCount: 5,
		CreatedAt: time.Now().Add(-15 *time.Hour),
	}

	mockRepo.EXPECT().
	FindByID(postID, userID).
	Return(post, nil)

	response, err := service.FindByID(postID, userID, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID != postID {
		t.Errorf("got response.ID %d, want %d\n", response.ID, postID)
	}

	if response.Author.ID != post.AuthorID {
		t.Errorf("got response.Author.ID %d, want %d\n", response.Author.ID, post.ID)
	}

	if response.Author.FullName != post.FullName {
		t.Errorf("got response.Author.FullName %s, want %s\n", response.Author.FullName, post.FullName)
	}

	if response.Title != post.Title {
		t.Errorf("got response.Title %s, want %s\n", response.Title, post.Title)
	}

	if response.Content != post.Content {
		t.Errorf("got response.Content %s, want %s\n", response.Content, post.Content)
	}

	if response.Likes != post.Likes {
		t.Errorf("got response.likes %d, want %d\n", response.Likes, post.Likes)
	}

	if response.Liked != post.Liked {
		t.Errorf("got response.Liked %v, want %v\n", response.Liked, post.Liked)
	}

	if response.IsPublished != post.IsPublished {
		t.Errorf("got response.IsPublished %v, want %v\n", response.IsPublished, post.IsPublished)
	}

	if response.CommentCount != post.CommentCount {
		t.Errorf("got response.CommentCount %d, want %d\n", response.CommentCount, post.CommentCount)
	}

	if len(response.TopComments) != len(post.TopComments) {
		t.Errorf("got response.TopComments %v, want %v\n", len(response.TopComments), len(post.TopComments))
	}

	if response.CreatedAt != post.CreatedAt {
		t.Errorf("got response.CreatedAt %v, want %v\n", response.CreatedAt, post.CreatedAt)
	}
}

func TestPostService_FindByID_Draft_NonOwner(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	postID := entity.PostID(uuid.New())
	userID := entity.UserID(5)

	post := &post.PostDetailRow{
		ID: postID,
		AuthorID: 2,
		IsPublished: false,
	}

	mockRepo.EXPECT().
	FindByID(postID, userID).
	Return(post, nil)

	response, err := service.FindByID(postID, userID, false)
	if response != nil {
		t.Fatalf("expected no response, got %v", response)
	}

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err != domainErrors.ErrNotFound {
		t.Errorf("got error %v, want %v\n", err, domainErrors.ErrNotFound)
	}
}


func TestPostService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	update := &entity.Post{
		ID: entity.PostID(uuid.New()),
		Title: "This is a sample title",
		Content: "This is a sample Content",
		IsPublished: false,
	}

	repoResponse := &post.PostDetailRow{
		ID: update.ID,
		AuthorID: entity.UserID(7),
		FullName: "John Doe",
		Title: update.Title,
		Content: update.Content,
		Likes: 25,
		Liked: false,
		IsPublished: update.IsPublished,
		CommentCount: 9,
		CreatedAt: time.Now().Add(-26 * time.Hour),
	}

	mockRepo.EXPECT().
	Update(update).
	Return(repoResponse, nil)

	response, err := service.Update(update)

	if err != nil {
		t.Fatalf("expected no err, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID != repoResponse.ID {
		t.Errorf("got response.ID %v, want %v\n", response.ID, repoResponse.ID)
	}

	if response.Author.ID != repoResponse.AuthorID {
		t.Errorf("got response.Author.ID %v, want %v\n", response.Author.ID, repoResponse.AuthorID)
	}

	if response.Author.FullName != repoResponse.FullName {
		t.Errorf("got response.FullName %s, want %s\n", response.Author.FullName, repoResponse.FullName)
	}

	if response.Title != repoResponse.Title {
		t.Errorf("got response.Title %s, want %s\n", response.Title, repoResponse.Title)
	}

	if response.Content != repoResponse.Content {
		t.Errorf("got response.Content %s, want %s\n", response.Content, repoResponse.Content)
	}

	if response.Likes != repoResponse.Likes {
		t.Errorf("got response.Likes %d, want %d\n", response.Likes, repoResponse.Likes)
	}

	if response.Liked != repoResponse.Liked {
		t.Errorf("got response.Liked %v, want %v\n", response.Liked, repoResponse.Liked)
	}

	if response.IsPublished != repoResponse.IsPublished {
		t.Errorf("got response.IsPublished %v, want %v\n", response.IsPublished, repoResponse.IsPublished)
	}

	if response.CommentCount != repoResponse.CommentCount {
		t.Errorf("got response.CommentCount %d, want %d\n", response.CommentCount, repoResponse.CommentCount)
	}

	if response.CreatedAt != repoResponse.CreatedAt {
		t.Errorf("got response.CreatedAt %v, want %v\n", response.CreatedAt, repoResponse.CreatedAt)
	}
}

func TestPostService_Update_ErrNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	post := &entity.Post{}
	mockRepo.EXPECT().
	Update(post).
	Return(nil, domainErrors.ErrNotFound)

	response, err := service.Update(post)

	if response != nil {
		t.Fatalf("expected no response, got %v", response)
	}

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != domainErrors.ErrNotFound {
		t.Errorf("got err %v, want %v\n", err, domainErrors.ErrNotFound)
	}
}


func TestPostService_Update_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	expectedErr := errors.New("repository error")
	post := &entity.Post{}
	mockRepo.EXPECT().
	Update(post).
	Return(nil, expectedErr)

	response, err := service.Update(post)

	if response != nil {
		t.Fatal("expected no response, got nil")
	}

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != expectedErr {
		t.Errorf("got err %v, want %v\n", err, expectedErr)
	}
}


func TestPostService_DeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	postID := entity.PostID(uuid.New())
	userID := entity.UserID(21)

	mockRepo.EXPECT().
	DeleteByID(postID, userID).
	Return(nil)

	err := service.DeleteByID(postID, userID)

	if err != nil {
		t.Fatalf("expected no err, got %v", err)
	}
}


func TestPostService_DeleteByID_ErrNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	postID := entity.PostID(uuid.New())
	userID := entity.UserID(27)

	mockRepo.EXPECT().
	DeleteByID(postID, userID).
	Return(domainErrors.ErrNotFound)

	err := service.DeleteByID(postID, userID)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != domainErrors.ErrNotFound {
		t.Errorf("got err %v, want %v", err, domainErrors.ErrNotFound)
	}
}

func TestPostService_DeleteByID_RepositoryErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockPostRepository(ctrl)
	service := NewPostService(mockRepo)

	postID := entity.PostID(uuid.New())
	userID := entity.UserID(27)
	expectedErr := errors.New("repository error")

	mockRepo.EXPECT().
	DeleteByID(postID, userID).
	Return(expectedErr)

	err := service.DeleteByID(postID, userID)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != expectedErr {
		t.Errorf("got err %v, want %v", err, expectedErr)
	}
}


func TestPostService_FindAll_Published(t *testing.T) {
	queries := []struct{
		name string
		userID entity.UserID
		query post.PostQuery
		isStaff bool
		expectedQuery post.PostQuery
	} {
		{
			name: "empty status defaults to published",
			query: post.PostQuery{},
			expectedQuery: post.PostQuery{
				Status: "published",
			},
		},
		{
			name: "published status remains published",
			query: post.PostQuery{
				Status: "published",
			},
			expectedQuery: post.PostQuery{
				Status: "published",
			},
		},
		{
			name: "author filter",
			query: post.PostQuery{
				Author: "John Doe",
			},
			expectedQuery: post.PostQuery{
				Status: "published",
				Author: "John Doe",
			},
		},
		{
			name: "author me",
			userID: entity.UserID(1),
			query: post.PostQuery{
				Author: "me",
			},
			expectedQuery: post.PostQuery{
				Status: "published",
				Author: "me",
			},
		},
	}

	repoResponse := []post.PostListRow{
		{
			ID: entity.PostID(uuid.New()),
			AuthorID: entity.UserID(7),
			FullName: "John Doe",
			Title: "Sample Title",
			Content: "Sample Content",
			Likes: 83468,
			Liked: true,
			IsPublished: true,
			CommentCount: 78,
			CreatedAt: time.Now().Add(-183 * time.Hour),
		},
		{
			ID: entity.PostID(uuid.New()),
			AuthorID: entity.UserID(34),
			FullName: "Monroe",
			Title: "Sample Title 2",
			Content: "Sample Content 2",
			Likes: 547,
			Liked: false,
			IsPublished: true,
			CommentCount: 8,
			CreatedAt: time.Now().Add(-49 * time.Hour),
		},
		{
			ID: entity.PostID(uuid.New()),
			AuthorID: entity.UserID(17),
			FullName: "Tom Jerry",
			Title: "Sample Title 3",
			Content: "Sample Content 3",
			Likes: 2,
			Liked: false,
			IsPublished: true,
			CommentCount: 0,
			CreatedAt: time.Now().Add(-77 * time.Hour),
		},
	}

	for _, tt := range queries {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mocks.NewMockPostRepository(ctrl)
			service := NewPostService(mockRepo)

			mockRepo.EXPECT().
			FindAll(tt.userID, tt.expectedQuery).
			Return(repoResponse, nil)

			response, err := service.FindAll(tt.userID, tt.query, tt.isStaff)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if response == nil {
				t.Fatal("expected response, got nil")
			}

			if len(response) != len(repoResponse) {
				t.Errorf("got len(response) %d, want %d", len(response), len(repoResponse))
			}

			if response[0].ID != repoResponse[0].ID {
				t.Errorf("got response[0].ID %v, want %v", response[0].ID, repoResponse[0].ID)
			}

			if response[0].Author.ID != repoResponse[0].AuthorID {
				t.Errorf("got response[0].Author.ID %v, want %v", response[0].Author.ID, repoResponse[0].AuthorID)
			}

			if response[0].Author.FullName != repoResponse[0].FullName {
				t.Errorf("got response[0].Author.FullName %v, want %v", response[0].Author.FullName, repoResponse[0].FullName)
			}

			if response[0].Title != repoResponse[0].Title {
				t.Errorf("got response[0].Title %v, want %v", response[0].Title, repoResponse[0].Title)
			}

			if response[0].Likes != repoResponse[0].Likes {
				t.Errorf("got response[0].Likes %v, want %v", response[0].Likes, repoResponse[0].Likes)
			}

			if response[0].Liked != repoResponse[0].Liked {
				t.Errorf("got response[0].Liked %v, want %v", response[0].Liked, repoResponse[0].Liked)
			}

			if response[0].CommentCount != repoResponse[0].CommentCount {
				t.Errorf("got response[0].CommentCount %v, want %v", response[0].CommentCount, repoResponse[0].CommentCount)
			}

			if response[0].IsPublished != repoResponse[0].IsPublished {
				t.Errorf("got response[0].IsPublished %v, want %v", response[0].IsPublished, repoResponse[0].IsPublished)
			}

			if response[0].CreatedAt != repoResponse[0].CreatedAt {
				t.Errorf("got response[0].CreatedAt %v, want %v", response[0].CreatedAt, repoResponse[0].CreatedAt)
			}
		})

	}
}

func TestPostService_FindAll_Draft_And_All(t *testing.T) {
	queries := []struct{
		name string
		userID entity.UserID
		query post.PostQuery
		isStaff bool
		expectedQuery post.PostQuery
	} {
		{
			name: "status draft, empty author",
			userID: entity.UserID(1),
			query: post.PostQuery{
				Status: "draft",
			},
			expectedQuery: post.PostQuery{
				Status: "draft",
			},
		},
		{
			name: "status draft, author me",
			userID: entity.UserID(1),
			query: post.PostQuery{
				Status: "draft",
				Author: "me",
			},
			expectedQuery: post.PostQuery{
				Status: "draft",
				Author: "me",
			},
		},
		{
			name: "status draft, isStaff",
			userID: entity.UserID(7),
			query: post.PostQuery{
				Status: "draft",
				Author: "John Doe",
			},
			isStaff: true,
			expectedQuery: post.PostQuery{
				Status: "draft",
				Author: "John Doe",
			},
		},
		{
			name: "status all, empty author",
			userID: entity.UserID(1),
			query: post.PostQuery{
				Status: "all",
			},
			expectedQuery: post.PostQuery{
				Status: "all",
			},
		},
		{
			name: "status all, author me",
			userID: entity.UserID(1),
			query: post.PostQuery{
				Status: "all",
				Author: "me",
			},
			expectedQuery: post.PostQuery{
				Status: "all",
				Author: "me",
			},
		},
		{
			name: "status all, isStaff",
			userID: entity.UserID(7),
			query: post.PostQuery{
				Status: "all",
				Author: "John Doe",
			},
			isStaff: true,
			expectedQuery: post.PostQuery{
				Status: "all",
				Author: "John Doe",
			},
		},
	}

	repoResponse := []post.PostListRow{
		{
			ID: entity.PostID(uuid.New()),
			AuthorID: entity.UserID(1),
			FullName: "John Doe",
			Title: "Sample Title",
			Content: "Sample Content",
			Likes: 0,
			Liked: false,
			IsPublished: false,
			CommentCount: 0,
			CreatedAt: time.Now().Add(-200 * time.Hour),
		},
		{
			ID: entity.PostID(uuid.New()),
			AuthorID: entity.UserID(1),
			FullName: "John Doe",
			Title: "Sample Title 2",
			Content: "Sample Content 2",
			Likes: 83468,
			Liked: true,
			IsPublished: true,
			CommentCount: 78,
			CreatedAt: time.Now().Add(-183 * time.Hour),
		},
		{
			ID: entity.PostID(uuid.New()),
			AuthorID: entity.UserID(1),
			FullName: "John Doe",
			Title: "Sample Title 4",
			Content: "Sample Content 4",
			Likes: 0,
			Liked: false,
			IsPublished: false,
			CommentCount: 0,
			CreatedAt: time.Now().Add(-49 * time.Hour),
		},
		{
			ID: entity.PostID(uuid.New()),
			AuthorID: entity.UserID(17),
			FullName: "Tom Jerry",
			Title: "Sample Title 3",
			Content: "Sample Content 3",
			Likes: 0,
			Liked: false,
			IsPublished: false,
			CommentCount: 0,
			CreatedAt: time.Now().Add(-77 * time.Hour),
		},
	}

	for _, tt := range queries {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mocks.NewMockPostRepository(ctrl)
			service := NewPostService(mockRepo)

			mockRepo.EXPECT().
			FindAll(tt.userID, tt.expectedQuery).
			Return(repoResponse, nil)

			response, err := service.FindAll(tt.userID, tt.query, tt.isStaff)

			if err != nil {
				t.Fatalf("expected error to be nil, got %v", err)
			}

			if response == nil {
				t.Fatal("expected response, got nil")
			}

			if response[0].ID != repoResponse[0].ID {
				t.Errorf("expected response[0].ID %v, got %v", response[0].ID, repoResponse[0].ID)
			}

			if response[0].Author.FullName != repoResponse[0].FullName {
				t.Errorf("expected response[0].Author.FullName %v, got %v", response[0].Author.FullName, repoResponse[0].FullName)
			}

			if response[0].Author.ID != repoResponse[0].AuthorID {
				t.Errorf("expected response[0].Author.ID %v, got %v", response[0].Author.ID, repoResponse[0].AuthorID)
			}

			if response[0].Title != repoResponse[0].Title {
				t.Errorf("expected response[0].Title %v, got %v", response[0].Title, repoResponse[0].Title)
			}

			if response[0].Likes != repoResponse[0].Likes {
				t.Errorf("expected response[0].Likes %v, got %v", response[0].Likes, repoResponse[0].Likes)
			}

			if response[0].Liked != repoResponse[0].Liked {
				t.Errorf("expected response[0].Liked %v, got %v", response[0].Liked, repoResponse[0].Liked)
			}

			if response[0].IsPublished != repoResponse[0].IsPublished {
				t.Errorf("expected response[0].IsPublished %v, got %v", response[0].IsPublished, repoResponse[0].IsPublished)
			}

			if response[0].CommentCount != repoResponse[0].CommentCount {
				t.Errorf("expected response[0].CommentCount %v, got %v", response[0].CommentCount, repoResponse[0].CommentCount)
			}

			if response[0].CreatedAt != repoResponse[0].CreatedAt {
				t.Errorf("expected response[0].CreatedAt %v, got %v", response[0].CreatedAt, repoResponse[0].CreatedAt)
			}
		})
	}
}

func TestPostService_FindAll_Anon_RestrictedStatus(t *testing.T) {
	queries := []struct {
		name string
		userID entity.UserID
		query post.PostQuery
		isStaff bool
		expectedErr error
	} {
		{
			name: "status draft",
			query: post.PostQuery{
				Status: "draft",
			},
			expectedErr: domainErrors.ErrUnauthorized,
		},
		{
			name: "status all",
			query: post.PostQuery{
				Status: "all",
			},
			expectedErr: domainErrors.ErrUnauthorized,
		},
		{
			name: "status draft author me",
			query: post.PostQuery{
				Status: "draft",
			},
			expectedErr: domainErrors.ErrUnauthorized,
		},
		{
			name: "status draft author others",
			query: post.PostQuery{
				Status: "draft",
			},
			expectedErr: domainErrors.ErrUnauthorized,
		},
		{
			name: "status all author me",
			query: post.PostQuery{
				Status: "all",
			},
			expectedErr: domainErrors.ErrUnauthorized,
		},
		{
			name: "status all author others",
			query: post.PostQuery{
				Status: "all",
			},
			expectedErr: domainErrors.ErrUnauthorized,
		},
		{
			name: "status published author me",
			query: post.PostQuery{
				Author: "me",
			},
			expectedErr: domainErrors.ErrUnauthorized,
		},
	}

	for _, tt := range queries {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mocks.NewMockPostRepository(ctrl)
			service := NewPostService(mockRepo)

			response, err := service.FindAll(tt.userID, tt.query, tt.isStaff)

			if response != nil {
				t.Fatalf("expected no response, got %v", response)
			}

			if err == nil {
				t.Fatal("expected error, got nil")
			}

			if err != domainErrors.ErrUnauthorized {
				t.Errorf("expected error %v, got %v", domainErrors.ErrUnauthorized, err)
			}
		})
	}
}

func TestPostService_FindAll_Forbidden(t *testing.T) {
	queries := []struct {
		name string
		userID entity.UserID
		query post.PostQuery
		isStaff bool
		expectedQuery post.PostQuery
	} {
		{
			name: "status draft, author others",
			userID: entity.UserID(3),
			query: post.PostQuery{
				Status: "draft",
				Author: "John Doe",
			},
		},
		{
			name: "status all, author others",
			userID: entity.UserID(7),
			query: post.PostQuery{
				Status: "all",
				Author: "John Doe",
			},
		},
	}

	for _, tt := range queries {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockRepo := mocks.NewMockPostRepository(ctrl)
			service := NewPostService(mockRepo)

			response, err := service.FindAll(tt.userID, tt.query, tt.isStaff)

			if response != nil {
				t.Fatalf("expected response to be nil, got %v", response)
			}

			if err == nil {
				t.Fatal("expected error, got nil")
			}

			if err != domainErrors.ErrForbidden {
				t.Errorf("expected error %v, got %v", domainErrors.ErrForbidden, err)
			}
		})
	}
}
