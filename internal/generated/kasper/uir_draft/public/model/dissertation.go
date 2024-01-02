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

type Dissertation struct {
	StudentID      uuid.UUID
	Status         DissertationStatus
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DissertationID uuid.UUID `sql:"primary_key"`
	Semester       int32
	Name           string
}
