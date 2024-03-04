package student_handler

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	StudentRepository interface {
		// GetStudentStatus - возвращает статус студента
		GetStudentStatus(ctx context.Context, studentID uuid.UUID) (models.Student, error)
	}

	DissertationService interface {
		AllToStatus(ctx context.Context, studentID uuid.UUID, status string) error
		// GetDissertationPage - возвращает всю информацию для отрисовки страницы диссертации
		GetDissertationPage(ctx context.Context, studentID uuid.UUID) (models.DissertationPageResponse, error)
		// UpsertSemesterProgress - обновляет план подготовки диссертации
		UpsertSemesterProgress(ctx context.Context, studentID uuid.UUID, progress []models.SemesterProgressRequest) error
		UpsertDissertationInfo(ctx context.Context, studentID uuid.UUID, semester int32) error
		UpsertDissertationTitle(ctx context.Context, studentID uuid.UUID, title string) error
		GetDissertationData(ctx context.Context, studentID uuid.UUID, semester int32) (model.Dissertations, error)
	}

	ScientificWorksService interface {
		// GetScientificWorks - возвращает все научные работы студента
		GetScientificWorks(ctx context.Context, studentID uuid.UUID) ([]models.ScientificWork, error)
		// UpsertPublications - добавляет или обновляет научные публикации
		UpsertPublications(ctx context.Context, studentID, workID uuid.UUID, semester int32, publications []models.Publication) error
		// UpsertConferences - добавляет или обновляет научные конференции
		UpsertConferences(ctx context.Context, studentID, workID uuid.UUID, semester int32, conferences []models.Conference) error
		// UpsertResearchProjects - добавляет или обновляет научные исследования
		UpsertResearchProjects(ctx context.Context, studentID, workID uuid.UUID, semester int32, projects []models.ResearchProject) error
		// DeletePublications - удаляет научные публикации
		DeletePublications(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error
		// DeleteConferences - удаляет научные конференции
		DeleteConferences(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error
		// DeleteResearchProjects - удаляет научные исследования
		DeleteResearchProjects(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error
	}

	TeachingLoadService interface {
		// GetTeachingLoad - возвращает всю педагогическую нагрузку студента
		GetTeachingLoad(ctx context.Context, studentID uuid.UUID) ([]models.TeachingLoad, error)
		// UpsertClassroomLoad - добавляет или обновляет аудиторную педагогическую нагрузку
		UpsertClassroomLoad(ctx context.Context, studentID, tLoadID uuid.UUID, semester int32, loads []models.ClassroomLoad) error
		// UpsertIndividualLoad - добавляет или обновляет индивидуальную педагогическую нагрузку
		UpsertIndividualLoad(ctx context.Context, studentID, tLoadID uuid.UUID, semester int32, loads []models.IndividualStudentsLoad) error
		// UpsertAdditionalLoad - добавляет или обновляет дополнительную педагогическую нагрузку
		UpsertAdditionalLoad(ctx context.Context, studentID, tLoadID uuid.UUID, semester int32, loads []models.AdditionalLoad) error
		// DeleteClassroomLoad - удаляет аудиторную педагогическую нагрузку студента
		DeleteClassroomLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error
		// DeleteIndividualLoad - удаляет индивидуальную педагогическую нагрузку студента
		DeleteIndividualLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error
		// DeleteAdditionalLoad - удаляет дополнительную аудиторную педагогическую нагрузку студента
		DeleteAdditionalLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error
	}

	Authenticator interface {
		// Authenticate - проводит аутентификацию пользователя
		Authenticate(ctx context.Context, token, userType string) (*model.Users, error)
	}
)

type StudentHandler struct {
	student       StudentRepository
	dissertation  DissertationService
	scientific    ScientificWorksService
	load          TeachingLoadService
	authenticator Authenticator
}

func (h *StudentHandler) authenticate(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.Authenticate(ctx, token, model.UserType_Student.String())
	if err != nil {
		return user, err
	}

	return user, nil
}
