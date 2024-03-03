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
)

func MapErrorToCode(err error) int {
	switch {
	case errors.Is(err, ErrTokenExpired) || errors.Is(err, ErrWrongUserType):
		return http.StatusUnauthorized
	case errors.Is(err, pgx.ErrNoRows):
		return http.StatusNoContent
	}

	return http.StatusInternalServerError
}
