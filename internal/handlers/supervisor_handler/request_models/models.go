package request_models

import (
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
)

type UpsertFeedbackRequest struct {
	Feedback  models.FeedbackRequest `json:"feedback"`
	StudentID uuid.UUID              `json:"student_id,omitempty"`
}

type GetByStudentID struct {
	StudentID uuid.UUID `json:"student_id,omitempty"`
}

type AllToStatusRequest struct {
	StudentID uuid.UUID `json:"student_id,omitempty"`
	Status    string    `json:"status,omitempty"`
}
