package models

import (
	"github.com/google/uuid"
)

type Specialization struct {
	SpecializationID int32  `json:"specialization_id,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Group struct {
	GroupID int32  `json:"group_id,omitempty"`
	Name    string `json:"name,omitempty"`
}

type SemesterAmount struct {
	AmountID uuid.UUID `json:"amount_id,omitempty"`
	Amount   int32     `json:"amount,omitempty"`
}
