package usecase

import (
	"fmt"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo user.UserRepository
	token auth.TokenService
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

	refreshToken, err := s.token.GenerateRefreshToken(user.ID)
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
