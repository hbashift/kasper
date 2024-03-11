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
)

type AdministratorHandler struct {
	user          UserService
	authenticator Authenticator
}

func NewHandler(user UserService, authenticator Authenticator) *AdministratorHandler {
	return &AdministratorHandler{user: user, authenticator: authenticator}
}

func (h *AdministratorHandler) authenticate(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.Authenticate(ctx, token, model.UserType_Admin.String())
	if err != nil {
		return user, err
	}

	return user, nil
}
