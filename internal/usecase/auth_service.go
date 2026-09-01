package usecase

import (
	"fmt"
	"strings"

	"github.com/Adejare77/go-BlogPost-API/internal/domain/auth"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/user"
	"github.com/google/uuid"
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

func CompareHashedCredentials(hashedValue, value string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedValue),
		[]byte(value),
	)
}

func (s *AuthService) Login(email, password string) (*auth.AuthResult, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, domainErrors.ErrInvalidCredentials
	}

	if !user.IsActive {
		return nil, domainErrors.ErrAccountDisabled
	}

	if err := CompareHashedCredentials(*user.Password, password); err != nil {
		return nil, domainErrors.ErrInvalidCredentials
	}

	accessToken, err := s.tokenService.GenerateAccessToken(user.ID, user.IsStaff)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %w", err)
	}

	refreshToken, err := s.tokenService.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %w", err)
	}

	hashedToken, err := hashCredential(refreshToken, bcrypt.MinCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing referesh token: %w", err)
	}


	token := entity.RefreshToken {
		UserID: user.ID,
		IsStaff: user.IsStaff,
		TokenHash: hashedToken,
	}
	if err := s.refreshTokenRepo.Create(&token); err != nil {
		return nil, err
	}

	tokenID := token.ID.String()
	refreshTokenValue := tokenID  + "." + refreshToken

	return &auth.AuthResult{
		AccessToken: accessToken,
		RefreshToken: refreshTokenValue,
		UserID: user.ID,
	}, nil
}

func (s *AuthService) Logout(token string) error {
	refreshToken := strings.Split(token, ".")
	tokenID, _ := uuid.Parse(refreshToken[0])

	_, err := s.refreshTokenRepo.FindByTokenID(tokenID)
	if err != nil {
		return err
	}

	return s.refreshTokenRepo.RevokeToken(tokenID)
}

func (s *AuthService) Register(user *entity.User) error {
	if user.Password != nil {
		password, err := hashCredential(*user.Password, bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing password: %w", err)
		}
		user.Password = &password
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) FindByID(userID entity.UserID) (*user.UserDetail, error) {
	return s.userRepo.FindByID(userID)
}

func (s *AuthService) RefreshToken(token string) (auth.AuthResult, error) {
	oldRefreshTokenValue := strings.Split(token, ".")
	tokenID, _ := uuid.Parse(oldRefreshTokenValue[0])
	tokenValue := oldRefreshTokenValue[1]

	retrieveToken, err := s.refreshTokenRepo.FindByTokenID(tokenID)
	if err != nil {
		return auth.AuthResult{}, err
	}

	if err := CompareHashedCredentials(
		retrieveToken.TokenHash, tokenValue,
	); err != nil {
		return auth.AuthResult{}, domainErrors.ErrInvalidToken
	}

	if err := s.refreshTokenRepo.RevokeToken(retrieveToken.ID); err != nil {
		return auth.AuthResult{}, err
	}

	newRefreshToken, err := s.tokenService.GenerateRefreshToken()
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("error generating refresh token: %w", err)
	}

	newRefreshTokenHash, err := hashCredential(newRefreshToken, bcrypt.MinCost)
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("error hashing token: %w", err)
	}

	refreshToken := entity.RefreshToken{
		UserID: retrieveToken.UserID,
		TokenHash: newRefreshTokenHash,
	}

	if err := s.refreshTokenRepo.Create(&refreshToken); err != nil {
		return auth.AuthResult{}, err
	}

	refreshTokenValue := refreshToken.ID.String() + "." + newRefreshToken
	accessToken, err := s.tokenService.GenerateAccessToken(
		retrieveToken.UserID,
		retrieveToken.IsStaff,
	)
	if err != nil {
		return auth.AuthResult{}, fmt.Errorf("error generating access token: %w", err)
	}

	return auth.AuthResult{
		AccessToken: accessToken,
		RefreshToken: refreshTokenValue,
		UserID: retrieveToken.UserID,
	}, nil
}
