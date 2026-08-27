package postgres

import (
	"errors"

	domainErrors "github.com/Adejare77/go-BlogPost-API/internal/domain/errors"
	"gorm.io/gorm"
)

func MapError(err error) error {
	switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return domainErrors.ErrNotFound

		case errors.Is(err, gorm.ErrDuplicatedKey):
			return domainErrors.ErrAlreadyExists

		default:
			return err
	}
}
