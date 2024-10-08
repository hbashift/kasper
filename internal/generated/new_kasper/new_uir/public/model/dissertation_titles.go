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

type DissertationTitles struct {
	TitleID         uuid.UUID `sql:"primary_key"`
	StudentID       uuid.UUID
	Title           string
	CreatedAt       time.Time
	Status          ApprovalStatus
	AcceptedAt      *time.Time
	Semester        int32
	ResearchObject  string
	ResearchSubject string
}
