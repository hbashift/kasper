package supervisor_handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/supervisor"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

type SupervisorService interface {
	GetStudentsList(ctx context.Context, token string) (*mapping.ListOfStudents, error)
	GetDissertationPage(ctx context.Context, token string, studentID uuid.UUID) (*supervisor.DissertationPage, error)
	DownloadDissertation(ctx context.Context, token string, info *mapping.DownloadDissertation) (*model.Dissertation, error)
}

type supervisorHandler struct {
	service SupervisorService
}

func NewSupervisorHandler(service SupervisorService) *supervisorHandler {
	return &supervisorHandler{service: service}
}

func getUUID(ctx *gin.Context) (*uuid.UUID, error) {
	stringId := ctx.Param("id")

	id, err := uuid.Parse(stringId)

	if err != nil {
		return nil, err
	}

	return &id, nil
}
