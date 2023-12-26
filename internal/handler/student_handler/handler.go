package student_handler

import (
	"context"

	"uir_draft/internal/pkg/service/student"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentService interface {
	GetDissertationPage(ctx context.Context, token string) (*student.DissertationPage, error)
	UpsertSemesterPlan(ctx context.Context, token string, progress *mapping.Progress) error
	GetScientificWorks(ctx context.Context, token string) ([]*mapping.ScientificWork, error)
	InsertScientificWorks(ctx context.Context, token string, works []*mapping.ScientificWork) ([]*mapping.ScientificWork, error)
	UpdateScientificWorks(ctx context.Context, token string, works []*mapping.ScientificWork) error
	DeleteScientificWork(ctx context.Context, token string, deleteIDs *mapping.DeleteWorkIDs) ([]*mapping.ScientificWork, error)
	GetTeachingLoad(ctx context.Context, token string) (*mapping.TeachingLoad, error)
	UpsertTeachingLoad(ctx context.Context, token string, loads *mapping.TeachingLoad) (*mapping.TeachingLoad, error)
	DeleteTeachingLoad(ctx context.Context, token string, deleteIDs *mapping.DeleteWorkIDs) (*mapping.TeachingLoad, error)
}

type studentHandler struct {
	service StudentService
}

func NewStudentHandler(service StudentService) *studentHandler {
	return &studentHandler{service: service}
}

func getUUID(ctx *gin.Context) (*uuid.UUID, error) {
	stringId := ctx.Param("id")

	id, err := uuid.Parse(stringId)

	if err != nil {
		return nil, err
	}

	return &id, nil
}
