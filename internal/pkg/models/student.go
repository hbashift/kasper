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

type Student struct {
	// ID студента
	StudentID uuid.UUID `db:"students.student_id" json:"student_id,omitempty" format:"uuid"`
	// Полное имя
	FullName string `db:"students.full_name" json:"full_name,omitempty"`
	// Кафедра
	Department string `db:"students.department" json:"department,omitempty"`
	// Актуальный семестр
	ActualSemester int32 `db:"students.actual_semester" json:"actual_semester,omitempty"`
	// Количество лет обучения
	Years int32 `db:"students.years" json:"years,omitempty"`
	// Дата начала обучения
	StartDate time.Time `db:"students.start_date" json:"start_date"`
	// Статус обучения
	StudyingStatus model.StudentStatus `db:"students.studying_status" json:"studying_status,omitempty" enums:"academic,graduated,studying,expelled"`
	// Статус проверки и подтверждения
	Status model.ApprovalStatus `db:"students.status" json:"status,omitempty"`
	// Специализация
	Specialization string `db:"specializations.title" json:"specialization,omitempty"`
	// Название группы
	GroupName string `db:"groups.group_name" json:"group_name,omitempty"`
	// Флаг о возможности редактировать всю информацию
	CanEdit bool `db:"students.can_edit" json:"can_edit,omitempty"`
}
