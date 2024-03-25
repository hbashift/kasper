package supervisor

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

var ErrNonValidToken = errors.New("token is expired")

type (
	StudentRepository interface {
		GetNumberOfYears(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (int32, error)
		UpdateFeedback(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, feedback string) error
		GetStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (*models.StudentCommonInformation, error)
		GetListOfStudents(ctx context.Context, tx *pgxpool.Pool, supervisorID *uuid.UUID) ([]*model.Students, error)
	}
	SemesterRepository interface {
		GetStudentDissertationPlan(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) ([]*models.StudentDissertationPlan, error)
	}

	TokenRepository interface {
		Authenticate(ctx context.Context, token string, db *pgxpool.Pool) (*model.AuthorizationToken, error)
	}

	DissertationRepository interface {
		SetStatus(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, semester int32, status model.DissertationStatus) error
		GetStatuses(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*mapping.DissertationStatus, error)
		GetDissertationData(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, semester int32) (*model.Dissertation, error)
	}

	ScientificWorksRepository interface {
		GetScientificWorks(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.ScientificWork, error)
	}

	TeachingLoadRepo interface {
		GetStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.TeachingLoad, error)
	}
)

type Service struct {
	studRepo     StudentRepository
	tokenRepo    TokenRepository
	semesterRepo SemesterRepository
	dRepo        DissertationRepository
	scienceRepo  ScientificWorksRepository
	loadRepo     TeachingLoadRepo

	db *pgxpool.Pool
}

func NewService(studRepo StudentRepository, tokenRepo TokenRepository, semesterRepo SemesterRepository, dRepo DissertationRepository, scienceRepo ScientificWorksRepository, loadRepo TeachingLoadRepo, db *pgxpool.Pool) *Service {
	return &Service{studRepo: studRepo, tokenRepo: tokenRepo, semesterRepo: semesterRepo, dRepo: dRepo, scienceRepo: scienceRepo, loadRepo: loadRepo, db: db}
}
