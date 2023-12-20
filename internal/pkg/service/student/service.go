package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StudentRepository interface {
	GetStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) (*models.StudentCommonInformation, error)
	InsertStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, student model.Students) error
	UpdateStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, student model.Students) error
}

type ScientificWorkRepository interface {
	UpdateStudentScientificWorks(ctx context.Context, tx *pgxpool.Pool, works *model.ScientificWork) error
	InsertStudentScientificWorks(ctx context.Context, tx *pgxpool.Pool, works *model.ScientificWork) error
	GetScientificWorks(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.ScientificWork, error)
	DeleteStudentScientificWorks(ctx context.Context, tx *pgxpool.Pool, workIDs []*uuid.UUID) error
}

type DissertationRepository interface {
}

type SemesterRepository interface {
	UpsertSemesterPlan(ctx context.Context, tx *pgxpool.Pool, progress []*model.SemesterProgress) error
	GetStudentDissertationPlan(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) ([]*models.StudentDissertationPlan, error)
}

type TeachingLoadRepository interface {
	InsertStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error
	GetStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.TeachingLoad, error)
	UpdateStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error
}

type TokenRepository interface {
	// TODO сделать мидлварю
	Authenticate(ctx context.Context, token string) (*model.AuthorizationToken, error)
}

type Service struct {
	studRepo     StudentRepository
	tokenRepo    TokenRepository
	dRepo        DissertationRepository
	semesterRepo SemesterRepository
	scienceRepo  ScientificWorkRepository
	loadRepo     TeachingLoadRepository
	db           *pgxpool.Pool
}

func NewService(studRepo StudentRepository, tokenRepo TokenRepository, dRepo DissertationRepository, semesterRepo SemesterRepository, scienceRepo ScientificWorkRepository, loadRepo TeachingLoadRepository, db *pgxpool.Pool) *Service {
	return &Service{studRepo: studRepo, tokenRepo: tokenRepo, dRepo: dRepo, semesterRepo: semesterRepo, scienceRepo: scienceRepo, loadRepo: loadRepo, db: db}
}
