package models

import (
	"time"

	"github.com/google/uuid"
)

type StudentCommonInformation struct {
	DissertationTitle     string    `db:"dissertation_title" json:"theme,omitempty"`
	SupervisorName        string    `db:"supervisor_name" json:"teacherFullName,omitempty"`
	EnrollmentOrderNumber string    `db:"enrollment_order_number" json:"numberOfOrderOfStatement,omitempty"`
	StudyingStartDate     time.Time `db:"studying_start_date" json:"dateOfOrderOfStatement"`
	Semester              int32     `db:"semester_number" json:"actualSemestr,omitempty"`
	Feedback              *string   `db:"feedback" json:"feedback,omitempty"`
	TitlePageURL          string    `db:"title_page_url" json:"titlePageURL,omitempty"`
	ExplanatoryNoteURL    string    `db:"explanatory_note_url" json:"explanatoryNoteURL,omitempty"`
	StudentName           string    `db:"student_name"`
	NumberOfYears         int32     `db:"number_of_years" json:"number_of_years"`
}

type IDs struct {
	ID       uuid.UUID `db:"id"`
	Semester int       `db:"semester"`
}
