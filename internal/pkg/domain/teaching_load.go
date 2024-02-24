package domain

import (
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
)

type ClassroomLoad struct {
	LoadID      uuid.UUID               `json:"load_id,omitempty" db:"classroom_load.load_id"`
	Hours       int32                   `json:"hours,omitempty" db:"classroom_load.hours"`
	LoadType    model.ClassroomLoadType `json:"load_type,omitempty" db:"classroom_load.load_type"`
	MainTeacher string                  `json:"main_teacher,omitempty" db:"classroom_load.main_teacher"`
	GroupName   string                  `json:"group_name,omitempty" db:"classroom_load.group_name"`
	SubjectName string                  `json:"subject_name,omitempty" db:"classroom_load.subject_name"`
}

type IndividualStudentsLoad struct {
	LoadID         uuid.UUID `json:"load_id,omitempty" db:"individual_students_load.load_id"`
	StudentsAmount int32     `json:"students_amount,omitempty" db:"individual_students_load.students_amount"`
	Comment        *string   `json:"comment,omitempty" db:"individual_students_load.comment"`
}

type AdditionalLoad struct {
	LoadID  uuid.UUID `json:"load_id,omitempty" db:"additional_load.load_id"`
	Name    string    `json:"name,omitempty" db:"additional_load.name"`
	Volume  *string   `json:"volume,omitempty" db:"additional_load.volume"`
	Comment *string   `json:"comment,omitempty" db:"additional_load.comment"`
}

type TeachingLoad struct {
	LoadsID                uuid.UUID              `json:"loads_id,omitempty" db:"teaching_load.loads_id"`
	StudentID              uuid.UUID              `json:"student_id,omitempty" db:"teaching_load.student_id"`
	Semester               int                    `json:"semester,omitempty" db:"teaching_load.semester"`
	ApprovalStatus         model.ApprovalStatus   `json:"approval_status,omitempty" db:"teaching_load.approval_status"`
	UpdatedAt              time.Time              `json:"updated_at" db:"teaching_load.updated_at"`
	AcceptedAt             *time.Time             `json:"accepted_at,omitempty" db:"teaching_load.accepted_at"`
	ClassroomLoad          ClassroomLoad          `json:"classroom_load"`
	IndividualStudentsLoad IndividualStudentsLoad `json:"individual_students_load"`
	AdditionalLoad         AdditionalLoad         `json:"additional_load"`
}
