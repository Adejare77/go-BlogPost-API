package usecase

import (
	"fmt"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo user.UserRepository
	token auth.TokenService
}

func NewAuthService(repo user.UserRepository, token auth.TokenService) *AuthService {
	return &AuthService{
		repo: repo,
		token: token,
	}
}

func (s *AuthService) Login(email, password string) (*auth.AuthResult, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("Invalid Credential")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(password)); err != nil {
			return nil, fmt.Errorf("Invalid Credential")
		}

	accessToken, err := s.token.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.token.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &auth.AuthResult{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		UserID: user.ID,
		Email: user.Email,
	}, nil
}

func (s *AuthService) Logout(userID entity.UserID) error {}

func (s *AuthService) Register(user *entity.User) error {
	if user.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}
		password := string(hash)
		user.Password = &password
	}

	// log user created
	return s.repo.Create(user)
}

func (s *AuthService) FindByID(userID entity.UserID) (*user.UserDetail, error) {
	return s.repo.FindByID(userID)
}
