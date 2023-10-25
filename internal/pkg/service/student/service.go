package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StudentRepository interface {
	GetStudentCommonInfo(ctx context.Context, tx pgx.Tx, clientID uuid.UUID) (*models.StudentCommonInformation, error)
	InsertStudentCommonInfo(ctx context.Context, tx pgx.Tx, student model.Students) error
	UpdateStudentCommonInfo(ctx context.Context, tx pgx.Tx, student model.Students) error
}

type DissertationRepository interface {
}

type SemesterRepository interface {
	UpsertSemesterPlan(ctx context.Context, tx pgx.Tx, progress []model.SemesterProgress) error
	GetSemesterProgress(ctx context.Context, tx pgx.Tx, clientID uuid.UUID) ([]*models.StudentDissertationPlan, error)
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
	db           *pgxpool.Conn
}

func NewService(studRepo StudentRepository, tokenRepo TokenRepository, dRepo DissertationRepository) *Service {
	return &Service{studRepo: studRepo, tokenRepo: tokenRepo, dRepo: dRepo}
}
