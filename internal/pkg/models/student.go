package models

import (
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

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

type StudentList struct {
	StudentID      uuid.UUID            `db:"students.student_id" json:"student_id,omitempty"`
	FullName       string               `db:"students.full_name" json:"full_name,omitempty"`
	Department     string               `db:"students.department" json:"department,omitempty"`
	ActualSemester int32                `db:"students.actual_semester" json:"actual_semester,omitempty"`
	Years          int32                `db:"students.years" json:"years,omitempty"`
	StartDate      time.Time            `db:"students.start_date" json:"start_date"`
	StudyingStatus model.StudentStatus  `db:"students.studying_status" json:"studying_status,omitempty"`
	Status         model.ApprovalStatus `db:"students.status" json:"status,omitempty"`
	Specialization string               `db:"specializations.title" json:"specialization,omitempty"`
	GroupName      string               `db:"groups.group_name" json:"group_name,omitempty"`
}
