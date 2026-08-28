package errors

import (
	"errors"
	"net/http"

	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, err error) {
	var status int
	var code string
	var message string

	switch {
	case errors.Is(err, domainErrors.ErrNotFound):
		status = http.StatusNotFound
		code = "not_found"
		message = "resource not found"

	case errors.Is(err, domainErrors.ErrAlreadyExists):
		status = http.StatusConflict
		code = "already_exists"
		message = "email already exists"

	case errors.Is(err, domainErrors.ErrInvalidCredentials) || errors.Is(err, domainErrors.ErrAccountDisabled):
		status = http.StatusUnauthorized
		code = "invalid_credentials"
		message = "invalid credentials"

	case errors.Is(err, domainErrors.ErrTokenRevoked) || errors.Is(err, domainErrors.ErrInvalidToken):
		status = http.StatusUnauthorized
		code = "invalid_or_expired_token"
		message = "invalid or expired token"

	default:
		status = http.StatusInternalServerError
		code = "internal_error"
		message = "internal server error"
	}

	response := errorResponse{
		Code: code,
		Message: message,
	}

	ctx.JSON(status, response)
}


func HandleRequestError(ctx *gin.Context, status int, field, message string, err error) {
	response := errorResponse{
		Code: field,
		Message: message,
	}

	ctx.JSON(status, response)
}
