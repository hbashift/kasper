package supervisor_handler

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	DissertationService interface {
		AllToStatus(ctx context.Context, studentID uuid.UUID, status string) error
		// GetDissertationPage - возвращает всю информацию для отрисовки страницы диссертации
		GetDissertationPage(ctx context.Context, studentID uuid.UUID) (models.DissertationPageResponse, error)
		// GetDissertationData - возвращает данные для скачивания файла
		GetDissertationData(ctx context.Context, studentID uuid.UUID, semester int32) (model.Dissertations, error)
	}

	ScientificWorksService interface {
		// GetScientificWorks - возвращает все научные работы студента
		GetScientificWorks(ctx context.Context, studentID uuid.UUID) ([]models.ScientificWork, error)
	}

	TeachingLoadService interface {
		// GetTeachingLoad - возвращает всю педагогическую нагрузку студента
		GetTeachingLoad(ctx context.Context, studentID uuid.UUID) ([]models.TeachingLoad, error)
	}

	Authenticator interface {
		// Authenticate - проводит аутентификацию пользователя
		Authenticate(ctx context.Context, token, userType string) (*model.Users, error)
	}

	SupervisorService interface {
		// UpsertFeedback - обновляет или добавляет фидбэк от научного руководителя
		UpsertFeedback(ctx context.Context, studentID uuid.UUID, request models.FeedbackRequest) error

		GetStudentList(ctx context.Context, supervisorID uuid.UUID) ([]models.Student, error)
	}
)

type SupervisorHandler struct {
	dissertation  DissertationService
	scientific    ScientificWorksService
	load          TeachingLoadService
	authenticator Authenticator
	supervisor    SupervisorService
}

func (h *SupervisorHandler) authenticate(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.Authenticate(ctx, token, model.UserType_Supervisor.String())
	if err != nil {
		return user, err
	}

	return user, nil
}
