package supervisor_handler

import (
	"context"

	"github.com/google/uuid"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/supervisor"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

type SupervisorService interface {
	GetStudentsList(ctx context.Context, token string) (*mapping.ListOfStudents, error)
	GetDissertationPage(ctx context.Context, token string, studentID uuid.UUID) (*supervisor.DissertationPage, error)
	DownloadDissertation(ctx context.Context, token string, info *mapping.DownloadDissertation) (*model.Dissertation, error)
	UpdateFeedback(ctx context.Context, token string, info *mapping.UpdateFeedback) error
	SetStatus(ctx context.Context, token string, info *mapping.SetStatus) error
}

type supervisorHandler struct {
	service SupervisorService
}

func NewSupervisorHandler(service SupervisorService) *supervisorHandler {
	return &supervisorHandler{service: service}
}
