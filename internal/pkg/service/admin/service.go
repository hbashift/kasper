package admin

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"
	"uir_draft/internal/pkg/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	UsersRepository interface {
		GetSupervisorsTx(ctx context.Context, tx pgx.Tx) ([]models.Supervisor, error)
		SetNewSupervisorTx(ctx context.Context, tx pgx.Tx, studentID, supervisorID uuid.UUID) error
		GetStudentSupervisorPairsTx(ctx context.Context, tx pgx.Tx) ([]models.StudentSupervisorPair, error)
		GetStudentTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (model.Students, error)
		SetStudentStatusTx(ctx context.Context, tx pgx.Tx, status model.ApprovalStatus, studyingStatus model.StudentStatus, studentID uuid.UUID) error
		GetStudentsList(ctx context.Context, tx pgx.Tx) ([]models.Student, error)
		SetStudentFlags(ctx context.Context, tx pgx.Tx, studyingStatus model.StudentStatus, canEdit bool, studentID uuid.UUID) error
	}

	MarksRepository interface {
		UpsertAttestationMarksTx(ctx context.Context, tx pgx.Tx, models []model.Marks) error
	}
)

type Service struct {
	clientRepo UsersRepository
	marksRepo  MarksRepository
	db         *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{
		clientRepo: repository.NewClientRepository(),
		marksRepo:  repository.NewMarksRepository(),
		db:         db,
	}
}
