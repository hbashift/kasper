package models

import (
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
)

type StudentDissertationPlan struct {
	Name    string `db:"name" json:"name,omitempty"`
	First   bool   `db:"first" json:"id1,omitempty"`
	Second  bool   `db:"second" json:"id2,omitempty"`
	Third   bool   `db:"third" json:"id3,omitempty"`
	Forth   bool   `db:"forth" json:"id4,omitempty"`
	Fifth   bool   `db:"fifth" json:"id5,omitempty"`
	Sixth   bool   `db:"sixth" json:"id6,omitempty"`
	Seventh bool   `db:"seventh" json:"id7,omitempty"`
	Eighth  bool   `db:"eighth" json:"id8,omitempty"`
}

type DissertationPage struct {
	Dissertations []Dissertation
	Titles        []DissertationTitle
}

type Dissertation struct {
	DissertationID uuid.UUID
	StudentID      uuid.UUID
	Status         model.DissertationStatus
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	Semester       int32
}

type DissertationTitle struct {
	TitleID        uuid.UUID
	DissertationID uuid.UUID
	Title          string
	CreatedAt      time.Time
	Status         *model.ApprovalStatus
	AcceptedAt     *time.Time
}
