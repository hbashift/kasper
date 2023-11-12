package mapping

import (
	"time"

	"github.com/google/uuid"
)

type ScientificWork struct {
	WorkID     *uuid.UUID `json:"work_id,omitempty"`
	Semester   int32      `json:"semester,omitempty"`
	Name       string     `json:"name,omitempty"`
	State      string     `json:"state,omitempty"`
	Impact     float64    `json:"impact,omitempty"`
	OutputData *string    `json:"output_data,omitempty"`
	CoAuthors  *string    `json:"co_authors,omitempty"`
	WorkType   *string    `json:"work_type,omitempty"`
}

type SemesterProgress struct {
	SemesterProgressID int32      `json:"semesterProgressID,omitempty"`
	StudentID          uuid.UUID  `json:"studentID,omitempty"`
	First              bool       `json:"first,omitempty"`
	Second             bool       `json:"second,omitempty"`
	Third              bool       `json:"third,omitempty"`
	Forth              bool       `json:"forth,omitempty"`
	Fifth              *bool      `json:"fifth,omitempty"`
	Sixth              *bool      `json:"sixth,omitempty"`
	ProgressName       string     `json:"progressName,omitempty"`
	LastUpdated        *time.Time `json:"lastUpdated,omitempty"`
	ClientID           uuid.UUID  `json:"clientID,omitempty"`
}

type Progress struct {
	Progress []SemesterProgress `json:"progress,omitempty"`
}
