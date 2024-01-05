package admin

import (
	"context"

	"github.com/google/uuid"
	"uir_draft/internal/pkg/service/admin"
	"uir_draft/internal/pkg/service/admin/mapping"
	studmapping "uir_draft/internal/pkg/service/student/mapping"
)

type AdministratorService interface {
	UpdateStudentsCommonInfo(ctx context.Context, token string, info *mapping.UpdateStudentsCommonInfo) error
	SetAcademicLeave(ctx context.Context, token string, leave *mapping.SetAcademicLeave) error
	GetStudSupPairs(ctx context.Context, token string) (*mapping.GetStudSupPairs, error)
	ChangeSupervisor(ctx context.Context, token string, studSup *mapping.ChangeSupervisorArray) error
	GetTeachingLoad(ctx context.Context, token string, studentID uuid.UUID) (*studmapping.TeachingLoad, error)
	GetScientificWorks(ctx context.Context, token string, studentID uuid.UUID) (*studmapping.Works, error)
	GetDissertationPage(ctx context.Context, token string, studentID uuid.UUID) (*admin.DissertationPage, error)
}

type administratorHandler struct {
	service AdministratorService
}

func NewAdministratorHandler(service AdministratorService) *administratorHandler {
	return &administratorHandler{service: service}
}
