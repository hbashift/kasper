package models

import (
	"time"

	"github.com/google/uuid"
)

type AllMarks struct {
	Exams            []Exam            `json:"exams,omitempty"`
	SupervisorMarks  []SupervisorMark  `json:"supervisor_marks,omitempty"`
	AttestationMarks []AttestationMark `json:"attestation_marks,omitempty"`
}

type Exam struct {
	ExamID    uuid.UUID  `json:"exam_id,omitempty"`
	StudentID uuid.UUID  `json:"student_id,omitempty"`
	ExamType  int32      `json:"exam_type,omitempty"`
	Semester  int32      `json:"semester,omitempty"`
	Mark      int32      `json:"mark,omitempty"`
	SetAt     *time.Time `json:"set_at,omitempty"`
}

type SupervisorMark struct {
	MarkID       uuid.UUID `json:"mark_id,omitempty"`
	StudentID    uuid.UUID `json:"student_id,omitempty"`
	Mark         int32     `json:"mark,omitempty"`
	Semester     int32     `json:"semester,omitempty"`
	SupervisorID uuid.UUID `json:"supervisor_id,omitempty"`
}

type AttestationMark struct {
	StudentID uuid.UUID `json:"student_id,omitempty"`
	Mark      int32     `json:"mark,omitempty"`
	Semester  int32     `json:"semester,omitempty"`
}

type ExamRequest struct {
	ExamType int32 `json:"exam_type,omitempty"`
	Semester int32 `json:"semester,omitempty"`
	Mark     int32 `json:"mark,omitempty"`
}

type AttestationMarkRequest struct {
	StudentID uuid.UUID `json:"student_id,omitempty"`
	Mark      int32     `json:"mark,omitempty"`
	Semester  int32     `json:"semester,omitempty"`
}
