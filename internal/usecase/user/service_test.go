package user

import (
	"errors"
	"testing"
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
	"github.com/Adejare77/go-BlogPost-API/internal/mocks"
	"go.uber.org/mock/gomock"
)

func TestUserService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	users := []user.UserDetail{
		{
			ID: entity.UserID(1),
			Email: "test1@gmail.com",
			FullName: "John Doe",
			CreatedAt: time.Now().Add(-72 *time.Hour),
			IsStaff: false,
			IsActive: true,
		},
		{
			ID: entity.UserID(2),
			Email: "test2@gmail.com",
			FullName: "Bonnie Clyde",
			CreatedAt: time.Now().Add(-38 *time.Hour),
			IsStaff: true,
			IsActive: true,
		},
	}

	mockRepo.EXPECT().
	FindAll().
	Return(users, nil)

	response, err := service.FindAll()

	if err != nil {
		t.Fatalf("expected err to be nil, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response[0].ID != users[0].ID {
		t.Errorf("expected response[0].ID %v, got %v", response[0].ID, users[0].ID)
	}

	if response[0].Email != users[0].Email {
		t.Errorf("expected response[0].Email %v, got %v", response[0].Email, users[0].Email)
	}

	if response[0].FullName != users[0].FullName {
		t.Errorf("expected response[0].FullName %v, got %v", response[0].FullName, users[0].FullName)
	}

	if response[0].CreatedAt != users[0].CreatedAt {
		t.Errorf("expected response[0].CreatedAt %v, got %v", response[0].CreatedAt, users[0].CreatedAt)
	}

	if response[0].IsStaff != users[0].IsStaff {
		t.Errorf("expected response[0].IsStaff %v, got %v", response[0].IsStaff, users[0].IsStaff)
	}

	if response[0].IsActive != users[0].IsActive {
		t.Errorf("expected response[0].IsActive %v, got %v", response[0].IsActive, users[0].IsActive)
	}
}

func TestUserService_FindAll_RepoErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)
	expectedErr := errors.New("Repository Error")

	mockRepo.EXPECT().
	FindAll().
	Return(nil, expectedErr)

	response, err := service.FindAll()

	if response != nil {
		t.Fatalf("expected response to be nil, got %v", response)
	}

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != expectedErr {
		t.Errorf("expected err %v, got %v", err, expectedErr)
	}
}

func TestUserService_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(6)
	repoResponse := &user.UserDetail{
		ID: targetedID,
		Email: "test@gmail.com",
		FullName: "John Doe",
		CreatedAt: time.Now().Add(-73 * time.Hour),
		IsActive: true,
		IsStaff: false,
	}

	mockRepo.EXPECT().
	FindByID(targetedID).
	Return(repoResponse, nil)

	response, err := service.FindByID(targetedID)

	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID != repoResponse.ID {
		t.Fatalf("expected response.ID %v, got %v", response.ID, repoResponse.ID)
	}

	if response.Email != repoResponse.Email {
		t.Fatalf("expected response.Email %v, got %v", response.Email, repoResponse.Email)
	}

	if response.FullName != repoResponse.FullName {
		t.Fatalf("expected response.FullName %v, got %v", response.FullName, repoResponse.FullName)
	}

	if response.CreatedAt != repoResponse.CreatedAt {
		t.Fatalf("expected response.CreatedAt %v, got %v", response.CreatedAt, repoResponse.CreatedAt)
	}

	if response.IsActive != repoResponse.IsActive {
		t.Fatalf("expected response.IsActive %v, got %v", response.IsActive, repoResponse.IsActive)
	}

	if response.IsStaff != repoResponse.IsStaff {
		t.Fatalf("expected response.IsStaff %v, got %v", response.IsStaff, repoResponse.IsStaff)
	}
}

func TestUserService_FindByID_RepoErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(6)
	expectedErr := errors.New("Repository Error")

	mockRepo.EXPECT().
	FindByID(targetedID).
	Return(nil, expectedErr)

	response, err := service.FindByID(targetedID)

	if response != nil {
		t.Fatalf("expected response to be nil, got %v", response)
	}

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != expectedErr {
		t.Errorf("expected error %v, got %v", err, expectedErr)
	}
}

func TestUserService_DeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(3)

	mockRepo.EXPECT().
	DeleteByID(targetedID).
	Return(nil)

	err := service.DeleteByID(targetedID)

	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}
}
func TestUserService_DeleteByID_RepoErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(9)
	expectedErr := errors.New("Repository Errors")

	mockRepo.EXPECT().
	DeleteByID(targetedID).
	Return(expectedErr)

	err := service.DeleteByID(targetedID)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != expectedErr {
		t.Errorf("expected error %v, got %v", err, expectedErr)
	}
}

func TestUserService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	repoResponse := &user.UserDetail{
		ID: entity.UserID(4),
		Email: "test@gmail.com",
		FullName: "John Doe",
		CreatedAt: time.Now().Add(-776 *time.Hour),
		IsActive: true,
		IsStaff: false,
	}

	password := "newPassWord123"

	userUpdate := &entity.User{
		Password: &password,
		Email: "test@gmai.com",
		FullName: "John Doe",
		IsActive: true,
		IsStaff: false,
	}

	mockRepo.EXPECT().
	Update(userUpdate).
	Return(repoResponse, nil)

	response, err := service.Update(userUpdate)

	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}

	if response == nil {
		t.Fatal("expected response, got nil")
	}

	if response.ID != repoResponse.ID {
		t.Errorf("expected response.ID %v, got %v", response.ID, repoResponse.ID)
	}

	if response.Email != repoResponse.Email {
		t.Errorf("expected response.Email %v, got %v", response.Email, repoResponse.Email)
	}

	if response.FullName != repoResponse.FullName {
		t.Errorf("expected response.FullName %v, got %v", response.FullName, repoResponse.FullName)
	}

	if response.CreatedAt != repoResponse.CreatedAt {
		t.Errorf("expected response.CreatedAt %v, got %v", response.CreatedAt, repoResponse.CreatedAt)
	}

	if response.IsActive != repoResponse.IsActive {
		t.Errorf("expected response.IsActive %v, got %v", response.IsActive, repoResponse.IsActive)
	}

	if response.IsStaff != repoResponse.IsStaff {
		t.Errorf("expected response.IsStaff %v, got %v", response.IsStaff, repoResponse.IsStaff)
	}
}

func TestUserService_Update_RepoErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	expectedErr := errors.New("Repository Error")

	mockRepo.EXPECT().
	Update(&entity.User{}).
	Return(nil, expectedErr)

	response, err := service.Update(&entity.User{})

	if response != nil {
		t.Fatalf("expected response to be nil, got %v", response)
	}

	if err == nil {
		t.Fatal("expected err, got nil")
	}

	if err != expectedErr {
		t.Errorf("expected error %v, to %v", err, expectedErr)
	}
}

func TestUserService_EnableByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(8)

	mockRepo.EXPECT().
	EnableByID(targetedID).
	Return(nil)

	err := service.EnableByID(targetedID)

	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}
}

func TestUserService_EnableByID_RepoErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(8)
	expectedError := errors.New("Repository Error")

	mockRepo.EXPECT().
	EnableByID(targetedID).
	Return(expectedError)

	err := service.EnableByID(targetedID)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != expectedError {
		t.Errorf("expected error %v, got %v", err, expectedError)
	}
}

func TestUserService_DisableByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(8)

	mockRepo.EXPECT().
	DisableByID(targetedID).
	Return(nil)

	err := service.DisableByID(targetedID)

	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}
}

func TestUserService_DisableByID_RepoErr(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	targetedID := entity.UserID(8)
	expectedError := errors.New("Repository Error")

	mockRepo.EXPECT().
	DisableByID(targetedID).
	Return(expectedError)

	err := service.DisableByID(targetedID)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != expectedError {
		t.Errorf("expected error %v, got %v", err, expectedError)
	}
}
