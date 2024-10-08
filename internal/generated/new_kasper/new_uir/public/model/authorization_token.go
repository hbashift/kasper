//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type AuthorizationToken struct {
	TokenID     int32 `sql:"primary_key"`
	UserID      uuid.UUID
	IsActive    bool
	TokenNumber string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}
