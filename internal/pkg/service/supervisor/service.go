package supervisor

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
	TokenRepository interface {
		GetUserIDByTokenTx(ctx context.Context, tx pgx.Tx, token string) (uuid.UUID, error)
	}

	UsersRepository interface {
		GetUserTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (model.Users, error)
	}

	FeedbackRepository interface {
		UpsertFeedbackTx(ctx context.Context, tx pgx.Tx, feedback model.Feedback) error
	}

	ClientRepository interface {
		GetSupervisorsStudentsTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) ([]models.Student, error)
		GetSupervisorTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) (models.Supervisor, error)
		GetSupervisorProfile(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) (models.SupervisorProfile, error)
	}

	MarksRepository interface {
		GetStudentsAttestationMarksTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Marks, error)
		GetStudentsExamResults(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Exams, error)
		GetStudentsSupervisorMarks(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.SupervisorMarks, error)
		UpsertStudentsSupervisorMark(ctx context.Context, tx pgx.Tx, model model.SupervisorMarks) error
	}
)

type Service struct {
	dissertationRepo FeedbackRepository
	tokenRepo        TokenRepository
	userRepo         UsersRepository
	client           ClientRepository
	marksRepo        MarksRepository
	db               *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{
		dissertationRepo: repository.NewDissertationRepository(),
		tokenRepo:        repository.NewTokenRepository(),
		userRepo:         repository.NewUsersRepository(),
		client:           repository.NewClientRepository(),
		marksRepo:        repository.NewMarksRepository(),
		db:               db,
	}
}
