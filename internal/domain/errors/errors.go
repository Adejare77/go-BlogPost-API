package errors

import "errors"

var (
	ErrNotFound = errors.New("resource not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrTokenRevoked = errors.New("token already revoked")
	ErrInvalidToken = errors.New("invalid or expired token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrAccountDisabled = errors.New("account disabled")
	ErrUnauthorized = errors.New("authentication required")
	ErrForbidden = errors.New("forbidden")
)
