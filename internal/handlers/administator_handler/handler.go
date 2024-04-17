package administator_handler

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetStudentSupervisorPairs(ctx context.Context) ([]models.StudentSupervisorPair, error)

type (
	UserService interface {
		GetStudentSupervisorPairs(ctx context.Context) ([]models.StudentSupervisorPair, error)
		ChangeSupervisor(ctx context.Context, pairs []models.ChangeSupervisor) error
		SetStudentFlags(ctx context.Context, students []models.SetStudentsFlags) error
		GetSupervisors(ctx context.Context) ([]models.Supervisor, error)
		GetStudentsList(ctx context.Context) ([]models.Student, error)
	}

	Authenticator interface {
		// Authenticate - проводит аутентификацию пользователя
		AuthenticateWithUserType(ctx context.Context, token, userType string) (*model.Users, error)
	}

	EnumService interface {
		GetSpecializations(ctx context.Context) ([]models.Specialization, error)
		InsertSpecializations(ctx context.Context, specializations []models.Specialization) error
		DeleteSpecializations(ctx context.Context, specIDs []int32) error

		GetGroups(ctx context.Context) ([]models.Group, error)
		InsertGroups(ctx context.Context, groups []models.Group) error
		DeleteGroups(ctx context.Context, groupIDs []int32) error
	}

	SupervisorService interface {
		GetSupervisorsStudents(ctx context.Context, supervisorID uuid.UUID) ([]models.Student, error)
		GetSupervisorProfile(ctx context.Context, supervisorID uuid.UUID) (models.SupervisorProfile, error)
	}
)

type AdministratorHandler struct {
	user          UserService
	authenticator Authenticator
	enum          EnumService
	supervisor    SupervisorService
}

func NewHandler(user UserService, authenticator Authenticator, enum EnumService, supervisor SupervisorService) *AdministratorHandler {
	return &AdministratorHandler{user: user, authenticator: authenticator, enum: enum, supervisor: supervisor}
}

func (h *AdministratorHandler) authenticate(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.AuthenticateWithUserType(ctx, token, model.UserType_Admin.String())
	if err != nil {
		return user, err
	}

	return user, nil
}
