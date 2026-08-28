package errors

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func jsonFieldName(req any, fieldName string) string {
	field, _ := reflect.TypeOf(req).FieldByName(fieldName)
	jsonName := field.Tag.Get("json")

	return jsonName
}

func Validator(ctx *gin.Context, req any, err validator.ValidationErrors) {
	fields := make(map[string]string)
	var key string
	var value string

	for _, fieldError := range err {
		key = jsonFieldName(req, fieldError.StructField())
		switch {
		case fieldError.Tag() == "required":
			value = fmt.Sprintf("%s is required", key)

		case fieldError.Tag() == "email":
			value = "invalid email address"

		case fieldError.Tag() == "min":
			value = "minimum of 3 length"

		case fieldError.Tag() == "oneOf":
			value = "must be one of: published,draft,all"
		}

		fields[key] = value
	}

	response := errorResponse{
		Code: "validation_error",
		Message: "request validation failed",
		Fields: fields,
	}

	ctx.JSON(http.StatusBadRequest, response)
}
