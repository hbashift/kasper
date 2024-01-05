package mapping

import (
	"time"

	"github.com/google/uuid"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
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
	Volume     int32      `json:"volume,omitempty"`
}

type Works struct {
	Works []*ScientificWork `json:"works,omitempty"`
	Years int32             `json:"years"`
}

type SemesterProgress struct {
	First        bool       `json:"first,omitempty"`
	Second       bool       `json:"second,omitempty"`
	Third        bool       `json:"third,omitempty"`
	Forth        bool       `json:"forth,omitempty"`
	Fifth        bool       `json:"fifth,omitempty"`
	Sixth        bool       `json:"sixth,omitempty"`
	Seventh      bool       `json:"seventh,omitempty"`
	Eighth       bool       `json:"eighth,omitempty"`
	ProgressName string     `json:"progressName,omitempty"`
	LastUpdated  *time.Time `json:"lastUpdated,omitempty"`
}

type Progress struct {
	Progress []SemesterProgress `json:"progress,omitempty"`
}

type IDs struct {
	IDs []string `json:"ids,omitempty"`
}

type SingleLoad struct {
	StudentID      uuid.UUID        `json:"student_id,omitempty"`
	Semester       int32            `json:"semester,omitempty"`
	Hours          int32            `json:"numberOfHours,omitempty"`
	AdditionalLoad *string          `json:"additional_load,omitempty"`
	LoadType       TeachingLoadType `json:"typeOfClasses,omitempty"`
	MainTeacher    string           `json:"mainTeacher,omitempty"`
	GroupName      string           `json:"numberOfGroup,omitempty"`
	SubjectName    string           `json:"subject,omitempty"`
	LoadID         *uuid.UUID       `json:"loadID"`
}

type TeachingLoadType string

const (
	TeachignLoadType_PRACTICE   TeachingLoadType = "Семинар"
	TeachingLoadType_LECTURE    TeachingLoadType = "Лекция"
	TeachingLoadType_LABORATORY TeachingLoadType = "Лабораторная"
	TeachingLoadType_UNKNOWN    TeachingLoadType = "unknown"
)

var TeachingLoadTypeMapFromDomain = map[model.TeachingLoadType]TeachingLoadType{
	model.TeachingLoadType_Practice:   TeachignLoadType_PRACTICE,
	model.TeachingLoadType_Lectures:   TeachingLoadType_LECTURE,
	model.TeachingLoadType_Laboratory: TeachingLoadType_LABORATORY,
}

var TeachingLoadTypeMapToDomain = map[TeachingLoadType]model.TeachingLoadType{
	TeachignLoadType_PRACTICE:   model.TeachingLoadType_Practice,
	TeachingLoadType_LECTURE:    model.TeachingLoadType_Lectures,
	TeachingLoadType_LABORATORY: model.TeachingLoadType_Laboratory,
}

type TeachingLoad struct {
	Array []SingleLoad `json:"array"`
	Years int32        `json:"years"`
}

type UploadDissertation struct {
	Semester Semester `form:"semester"`
}

type Semester struct {
	SemesterNumber int32 `json:"semester"`
}

type DownloadDissertation struct {
	Semester int32 `json:"semester"`
}

type DissertationStatus struct {
	Semester int32                    `db:"semester" json:"semester"`
	Status   model.DissertationStatus `db:"status" json:"status"`
}

type FirstRegistry struct {
	FullName        string     `json:"fullName,omitempty"`
	Department      string     `json:"department,omitempty"`
	EnrollmentOrder string     `json:"enrollmentOrder,omitempty"`
	Specialization  string     `json:"specialization,omitempty"`
	ActualSemester  int32      `json:"actualSemester,omitempty"`
	SupervisorID    *uuid.UUID `json:"supervisorID,omitempty"`
	StartDate       string     `json:"startDate,omitempty"`
	GroupNumber     *string    `json:"groupNumber,omitempty"`
	NumberOfYears   int32      `json:"numberOfYears"`
}

type Supervisor struct {
	Name         string    `json:"name,omitempty"`
	SupervisorID uuid.UUID `json:"supervisorID,omitempty"`
}

type Supervisors struct {
	Supervisors []*Supervisor `json:"supervisors,omitempty"`
}
