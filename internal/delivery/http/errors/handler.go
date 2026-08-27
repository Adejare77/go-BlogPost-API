package errors

import (
	"errors"
	"fmt"
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

	default:
		status = http.StatusInternalServerError
		code = "internal_error"
		message = "internal server error"
	}

	response := ErrorResponse{
		Code: code,
		Message: message,
	}

	ctx.JSON(status, response)
}


func HandleRequestError(ctx *gin.Context, err error) {
	status := http.StatusBadRequest
	code := "bad_request"
	message := fmt.Sprintf("Bad Request: %s", err.Error())

	response := ErrorResponse{
		Code: code,
		Message: message,
	}

	ctx.JSON(status, response)
}
