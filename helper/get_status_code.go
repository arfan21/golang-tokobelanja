package helper

import (
	"errors"
	"net/http"

	"github.com/arfan21/golang-tokobelanja/constant"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
)

func GetStatusCode(err error) int {
	if err.Error() == constant.ErrorEmailAlreadyExists.Error() {
		return http.StatusConflict
	}

	if err.Error() == constant.ErrorInvalidLogin.Error() {
		return http.StatusBadRequest
	}

	if err.Error() == constant.ErrorInvalidRole.Error() {
		return http.StatusBadRequest
	}

	if err.Error() == constant.ErrorOwnership.Error() {
		return http.StatusForbidden
	}

	if err.Error() == constant.ErrorCategoryDoesNotExists.Error() {
		return http.StatusNotFound
	}

	if isValidationError(err) {
		return http.StatusBadRequest
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, gorm.ErrMissingWhereClause) {
		return http.StatusBadRequest
	}

	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return http.StatusBadRequest
	}

	if isPostgresErrorUniqueViolation(err) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}

func isValidationError(err error) bool {
	_, ok := err.(validation.Errors)
	return ok
}

func isPostgresErrorUniqueViolation(err error) bool {
	pgError, ok := err.(*pgconn.PgError)
	if ok {
		return pgError.Code == pgerrcode.UniqueViolation
	}
	return false
}
