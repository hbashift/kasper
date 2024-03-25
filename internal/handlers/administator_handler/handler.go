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
		SetStudentStudyingStatus(ctx context.Context, studentID uuid.UUID, status string) error
		GetSupervisors(ctx context.Context) ([]models.Supervisor, error)
	}

	Authenticator interface {
		// Authenticate - проводит аутентификацию пользователя
		Authenticate(ctx context.Context, token, userType string) (*model.Users, error)
	}

	EnumService interface {
		GetSpecializations(ctx context.Context) ([]models.Specialization, error)
		InsertSpecializations(ctx context.Context, specializations []models.Specialization) error
		DeleteSpecializations(ctx context.Context, specIDs []int32) error

		GetGroups(ctx context.Context) ([]models.Group, error)
		InsertGroups(ctx context.Context, groups []models.Group) error
		DeleteGroups(ctx context.Context, groupIDs []int32) error
	}
)

type AdministratorHandler struct {
	user          UserService
	authenticator Authenticator
	enum          EnumService
}

func NewHandler(user UserService, authenticator Authenticator, enum EnumService) *AdministratorHandler {
	return &AdministratorHandler{user: user, authenticator: authenticator, enum: enum}
}

func (h *AdministratorHandler) authenticate(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.Authenticate(ctx, token, model.UserType_Admin.String())
	if err != nil {
		return user, err
	}

	return user, nil
}
