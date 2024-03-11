package request_models

import (
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
)

type ChangeSupervisorRequest struct {
	Pairs []models.ChangeSupervisor `json:"pairs" binding:"required"`
}

type SetStudentStudyingStatusRequest struct {
	StudentID      uuid.UUID `json:"student_id,omitempty"`
	StudyingStatus string    `json:"status,omitempty" enums:"academic,graduated,studying,expelled"`
}
