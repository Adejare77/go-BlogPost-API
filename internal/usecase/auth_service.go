package usecase

import (
	"fmt"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo user.UserRepository
	tokenService auth.TokenService
	refreshTokenRepo auth.RefreshTokenRepository
}

func NewAuthService(
	userRepo user.UserRepository,
	tokenService auth.TokenService,
	refreshTokenRepo auth.RefreshTokenRepository,
	) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		tokenService: tokenService,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func hashCredential(value string, cost int) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(value), cost)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (s *AuthService) Login(email, password string) (*auth.AuthResult, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid Credential")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(password)); err != nil {
			return nil, fmt.Errorf("invalid Credential")
		}

	accessToken, err := s.tokenService.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenService.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &auth.AuthResult{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		UserID: user.ID,
	}, nil
}

func (s *AuthService) Logout(refreshToken string) error {
	tokenHash, err := hashCredential(refreshToken, bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("error hashing refresh token: %w", err)
	}

	// retrieve from DB
	retrieveToken, err := s.refreshTokenRepo.FindByTokenHash(tokenHash)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	// revoke token
	if err := s.refreshTokenRepo.RevokeToken(retrieveToken.TokenHash); err != nil {
		return fmt.Errorf("error revoking token: %w", err)
	}

	return nil
}

func (s *AuthService) Register(user *entity.User) error {
	if user.Password != nil {
		password, err := hashCredential(*user.Password, bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}
		user.Password = &password
	}

	// log user created
	return s.userRepo.Create(user)
}

func (s *AuthService) FindByID(userID entity.UserID) (*user.UserDetail, error) {
	return s.userRepo.FindByID(userID)
}

func (s *AuthService) RefreshToken(token string) (auth.AuthResult, error) {
	tokenHash, err := hashCredential(token, bcrypt.MinCost)
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("error hashing token: %w", err)
	}

	retrieveToken, err := s.refreshTokenRepo.FindByTokenHash(tokenHash)
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("invalid token: %w", err)
	}

	if err := s.refreshTokenRepo.RevokeToken(retrieveToken.TokenHash); err != nil {
		return auth.AuthResult{}, fmt.Errorf("error revoking token: %w", err)
	}

	// generate new refresh tokne
	newRefreshToken, err := s.tokenService.GenerateRefreshToken()
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("error generating refresh token: %w", err)
	}

	// hash the new token
	newRefreshTokenHash, err := hashCredential(newRefreshToken, bcrypt.MinCost)
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("error hasing token: %w", err)
	}

	refreshToken := entity.RefreshToken{
		UserID: retrieveToken.UserID,
		TokenHash: newRefreshTokenHash,
	}
	if err := s.refreshTokenRepo.Create(&refreshToken); err != nil {
		return auth.AuthResult{}, fmt.Errorf("error creating new token in DB: %w", err)
	}

	accessToken, err := s.tokenService.GenerateAccessToken(retrieveToken.UserID)
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("error generating access token: %w", err)
	}

	return auth.AuthResult{
		AccessToken: accessToken,
		RefreshToken: newRefreshToken,
		UserID: retrieveToken.UserID,
	}, nil
}
