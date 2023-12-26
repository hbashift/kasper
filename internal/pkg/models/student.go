package models

import (
	"time"

	"github.com/google/uuid"
)

type StudentDissertationPlan struct {
	Name   string `db:"name" json:"name,omitempty"`
	First  bool   `db:"first" json:"id1,omitempty"`
	Second bool   `db:"second" json:"id2,omitempty"`
	Third  bool   `db:"third" json:"id3,omitempty"`
	Forth  bool   `db:"forth" json:"id4,omitempty"`
	Fifth  bool   `db:"fifth" json:"id5,omitempty"`
	Sixth  bool   `db:"sixth" json:"id6,omitempty"`
}

type StudentCommonInformation struct {
	DissertationTitle     string    `db:"dissertation_title" json:"theme,omitempty"`
	SupervisorName        string    `db:"supervisor_name" json:"teacherFullName,omitempty"`
	EnrollmentOrderNumber string    `db:"enrollment_order_number" json:"numberOfOrderOfStatement,omitempty"`
	StudyingStartDate     time.Time `db:"studying_start_date" json:"dateOfOrderOfStatement"`
	Semester              int32     `db:"semester_number" json:"actualSemestr,omitempty"`
	Feedback              *string   `db:"feedback" json:"feedback,omitempty"`
	DissertationStatus    *string   `db:"dissertation_status" json:"jobStatus,omitempty"`
	TitlePageURL          string    `db:"title_page_url" json:"titlePageURL,omitempty"`
	ExplanatoryNoteURL    string    `db:"explanatory_note_url" json:"explanatoryNoteURL,omitempty"`
	StudentName           string    `db:"student_name"`
}

type IDs struct {
	ID uuid.UUID `db:"id"`
}
