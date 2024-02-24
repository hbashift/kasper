package domain

import (
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
)

type Publication struct {
	PublicationID uuid.UUID               `json:"publication_id,omitempty" db:"publications.publication_id"`
	Name          string                  `json:"name,omitempty" db:"publications.name"`
	Index         model.WorkIndex         `json:"index,omitempty" db:"publications.index"`
	Impact        float64                 `json:"impact,omitempty" db:"publications.impact"`
	Status        model.PublicationStatus `json:"status,omitempty" db:"publications.status"`
	OutputData    *string                 `json:"output_data,omitempty" db:"publications.output_data"`
	CoAuthors     *string                 `json:"co_authors,omitempty" db:"publications.co_authors"`
	Volume        *int32                  `json:"volume,omitempty" db:"publications.volume"`
}

type Conference struct {
	ConferenceID   uuid.UUID              `json:"conference_id,omitempty" db:"conferences.conference_id"`
	Status         model.ConferenceStatus `json:"status,omitempty" db:"conferences.status"`
	Index          model.WorkIndex        `json:"index,omitempty" db:"conferences.index"`
	ConferenceName string                 `json:"conference_name,omitempty" db:"conferences.conference_name"`
	ReportName     string                 `json:"report_name,omitempty" db:"conferences.report_name"`
	Location       string                 `json:"location,omitempty" db:"conferences.location"`
	ReportedAt     time.Time              `json:"reported_at" db:"conferences.reported_at"`
}

type ResearchProject struct {
	ProjectID   uuid.UUID `json:"project_id,omitempty" db:"research_projects.project_id"`
	ProjectName string    `json:"project_name,omitempty" db:"research_projects.project_name"`
	StartAt     time.Time `json:"start_at" db:"research_projects.start_at"`
	EndAt       time.Time `json:"end_at" db:"research_projects.end_at"`
	AddInfo     *string   `json:"add_info,omitempty" db:"research_projects.add_info"`
	Grantee     *string   `json:"grantee,omitempty" db:"research_projects.grantee"`
}

type ScientificWork struct {
	WorksID         uuid.UUID            `json:"works_id,omitempty" db:"scientific_works.works_id"`
	Semester        int                  `json:"semester,omitempty" db:"scientific_works.semester"`
	StudentID       uuid.UUID            `json:"student_id,omitempty" db:"scientific_works.student_id"`
	ApprovalStatus  model.ApprovalStatus `json:"works_status,omitempty" db:"scientific_works.approval_status"`
	UpdatedAt       time.Time            `json:"updated_at" db:"scientific_works.updated_at"`
	AcceptedAt      *time.Time           `json:"accepted_at,omitempty" db:"scientific_works.accepted_at"`
	Publication     Publication          `json:"publication"`
	Conference      Conference           `json:"conference"`
	ResearchProject ResearchProject      `json:"research_project"`
}
