package supervisor

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"
	"uir_draft/internal/pkg/service/student/mapping"
)

var ErrNonValidToken = errors.New("token is expired")

type StudentRepository interface {
	UpdateFeedback(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, feedback string) error
	GetStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (*models.StudentCommonInformation, error)
	GetListOfStudents(ctx context.Context, tx *pgxpool.Pool, supervisorID *uuid.UUID) ([]*model.Students, error)
}

type SemesterRepository interface {
	GetStudentDissertationPlan(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) ([]*models.StudentDissertationPlan, error)
}

type TokenRepository interface {
	Authenticate(ctx context.Context, token string, db *pgxpool.Pool) (*model.AuthorizationToken, error)
}

type DissertationRepository interface {
	GetStatuses(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*mapping.DissertationStatus, error)
	GetDissertationData(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, semester int32) (*model.Dissertation, error)
}

type Service struct {
	studRepo     StudentRepository
	tokenRepo    TokenRepository
	semesterRepo SemesterRepository
	dRepo        DissertationRepository

	db *pgxpool.Pool
}

func NewService(studRepo StudentRepository, tokenRepo TokenRepository, semesterRepo SemesterRepository, dRepo DissertationRepository, db *pgxpool.Pool) *Service {
	return &Service{
		studRepo:     studRepo,
		tokenRepo:    tokenRepo,
		semesterRepo: semesterRepo,
		dRepo:        dRepo,
		db:           db,
	}
}
