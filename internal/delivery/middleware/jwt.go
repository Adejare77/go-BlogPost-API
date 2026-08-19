package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/Adejare77/go-BlogPost-API/internal/config"
	"github.com/Adejare77/go-BlogPost-API/internal/domain/entity"
	"github.com/golang-jwt/jwt/v5"
)


type JWTTokenService struct {
	secret string
}

func NewJWTTokenService(secret string) *JWTTokenService {
	return &JWTTokenService{
		secret: secret,
	}
}

type Claims struct {
	UserID entity.UserID
	jwt.RegisteredClaims
}

func (j *JWTTokenService) GenerateAccessToken(userID entity.UserID) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: config.Current.App.JWTIssuer,
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(config.Current.App.AccessTokenTTL))),
			IssuedAt: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secret))
}

func (*JWTTokenService) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
