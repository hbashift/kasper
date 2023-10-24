package models

import (
	"time"
)

type StudentDissertationPlan struct {
	Name   string `db:"name"`
	First  bool   `db:"first"`
	Second bool   `db:"second"`
	Third  bool   `db:"third"`
	Forth  bool   `db:"forth"`
	Fifth  bool   `db:"fifth"`
	Sixth  bool   `db:"sixth"`
}

type StudentCommonInformation struct {
	DissertationTitle     string    `db:"dissertation_title"`
	SupervisorName        string    `db:"supervisor_name"`
	EnrollmentOrderNumber string    `db:"enrollment_order_number"`
	StudyingStartDate     time.Time `db:"studying_start_date"`
	Semester              int32     `db:"semester_number"`
	Feedback              *string   `db:"feedback"`
	DissertationStatus    int32     `db:"dissertation_status"`
	TitlePageURL          string    `db:"title_page_url"`
	ExplanatoryNoteURL    string    `db:"explanatory_note_url"`
}

type DissertationPage struct {
	DissertationPlan []*StudentDissertationPlan `json:"dissertationPlan"`
	CommonInfo       StudentCommonInformation   `json:"commonInfo"`
}
