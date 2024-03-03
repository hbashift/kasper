package request_models

import (
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
)

type DeleteIDs struct {
	IDs      []uuid.UUID `json:"ids,omitempty"`
	Semester int32       `json:"semester,omitempty"`
}

type UpsertAdditionalLoadRequest struct {
	Loads    []models.AdditionalLoad `json:"loads,omitempty"`
	TLoadID  uuid.UUID               `json:"t_load_id,omitempty"`
	Semester int32                   `json:"semester,omitempty"`
}

type UpsertClassroomLoadRequest struct {
	Loads    []models.ClassroomLoad `json:"loads,omitempty"`
	TLoadID  uuid.UUID              `json:"t_load_id,omitempty"`
	Semester int32                  `json:"semester,omitempty"`
}

type UpsertConferencesRequest struct {
	Conferences []models.Conference `json:"conferences,omitempty"`
	WorkID      uuid.UUID           `json:"work_id,omitempty"`
	Semester    int32               `json:"semester,omitempty"`
}

type UpsertIndividualLoadRequest struct {
	Loads    []models.IndividualStudentsLoad `json:"loads,omitempty"`
	TLoadID  uuid.UUID                       `json:"t_load_id,omitempty"`
	Semester int32                           `json:"semester,omitempty"`
}

type UpsertPublicationsRequest struct {
	Publications []models.Publication `json:"publications,omitempty"`
	WorkID       uuid.UUID            `json:"work_id,omitempty"`
	Semester     int32                `json:"semester,omitempty"`
}

type UpsertResearchProjectsRequest struct {
	Projects []models.ResearchProject `json:"projects,omitempty"`
	WorkID   uuid.UUID                `json:"work_id,omitempty"`
	Semester int32                    `json:"semester,omitempty"`
}

type UpsertProgressRequest struct {
	Progresses []models.SemesterProgressRequest `json:"progresses"`
}
