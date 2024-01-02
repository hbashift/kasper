package supervisor_handler

import (
	"context"

	"github.com/google/uuid"
	"uir_draft/internal/pkg/service/supervisor"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

type SupervisorService interface {
	GetStudentsList(ctx context.Context, token string) (*mapping.ListOfStudents, error)
	GetDissertationPage(ctx context.Context, token string, studentID uuid.UUID) (*supervisor.DissertationPage, error)
}

type supervisorHandler struct {
	service SupervisorService
}

func NewSupervisorHandler(service SupervisorService) *supervisorHandler {
	return &supervisorHandler{service: service}
}
