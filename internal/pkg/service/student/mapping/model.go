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
	Works []ScientificWork `json:"works,omitempty"`
}

type SemesterProgress struct {
	First        bool       `json:"first,omitempty"`
	Second       bool       `json:"second,omitempty"`
	Third        bool       `json:"third,omitempty"`
	Forth        bool       `json:"forth,omitempty"`
	Fifth        *bool      `json:"fifth,omitempty"`
	Sixth        *bool      `json:"sixth,omitempty"`
	ProgressName string     `json:"progressName,omitempty"`
	LastUpdated  *time.Time `json:"lastUpdated,omitempty"`
}

type Progress struct {
	Progress []SemesterProgress `json:"progress,omitempty"`
}

type DeleteWorkIDs struct {
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
	TeachignLoadType_PRACTICE   TeachingLoadType = "practice"
	TeachingLoadType_LECTURE    TeachingLoadType = "lecture"
	TeachingLoadType_LABORATORY TeachingLoadType = "laboratory"
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
}

type UploadDissertation struct {
	Semester Semester `form:"data"`
}

type Semester struct {
	SemesterNumber int `json:"semester"`
}
