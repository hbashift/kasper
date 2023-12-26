package supervisor_handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

type SupervisorService interface {
	GetStudentsList(ctx context.Context, token string) (*mapping.ListOfStudents, error)
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
