package models

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v4"
)

var (
	ErrInvalidEnumValue      = errors.New("invalid enum value")
	ErrNonMutableStatus      = errors.New("status is 'on review' or 'approved'. user cannot edit any information")
	ErrTokenExpired          = errors.New("provided token is expired")
	ErrWrongUserType         = errors.New("wrong user type")
	ErrNotActualSemester     = errors.New("not actual semester")
	ErrUnknownApprovalStatus = errors.New("unknown approval status")
	ErrWrongPassword         = errors.New("wrong password")
	ErrHigherValueExpected   = errors.New("expected higher value")
)

func MapErrorToCode(err error) int {
	switch {
	case errors.Is(err, ErrTokenExpired) || errors.Is(err, ErrWrongUserType):
		return http.StatusUnauthorized
	case errors.Is(err, pgx.ErrNoRows):
		return http.StatusNoContent
	case errors.Is(err, ErrNotActualSemester) || errors.Is(err, ErrHigherValueExpected):
		return http.StatusBadRequest
	case errors.Is(err, ErrNonMutableStatus) || errors.Is(err, ErrWrongPassword):
		return http.StatusForbidden
	}

	return http.StatusInternalServerError
}
