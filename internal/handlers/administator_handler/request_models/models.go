package request_models

import (
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
)

type ChangeSupervisorRequest struct {
	Pairs []models.ChangeSupervisor `json:"pairs" binding:"required"`
}

type SetStudentFlagsRequest struct {
	Students []models.SetStudentsFlags `json:"students"`
}

type AddGroupsRequest struct {
	Groups []models.Group `json:"groups,omitempty"`
}

type AddSpecializationsRequest struct {
	Specs []models.Specialization `json:"specs,omitempty"`
}

type DeleteEnumRequest struct {
	IDs []int32 `json:"ids"`
}

type GetSupervisorsStudents struct {
	SupervisorID uuid.UUID `json:"supervisor_id"`
}
