package supervisor

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/handlers/authorization_handler/request_models"
	"uir_draft/internal/pkg/models"
	"uir_draft/internal/pkg/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type (
	TokenRepository interface {
		GetUserIDByTokenTx(ctx context.Context, tx pgx.Tx, token string) (uuid.UUID, error)
	}

	UsersRepository interface {
		GetUserTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (model.Users, error)
	}

	DissertationRepository interface {
		UpsertFeedbackTx(ctx context.Context, tx pgx.Tx, feedback model.Feedback) error
		GetDissertationDataBySemester(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, semester int32) (model.Dissertations, error)
	}

	ClientRepository interface {
		GetSupervisorsStudentsTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) ([]models.Student, error)
		GetSupervisorTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) (models.Supervisor, error)
		GetSupervisorProfile(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) (models.SupervisorProfile, error)
		InsertSupervisor(ctx context.Context, tx pgx.Tx, supervisor model.Supervisors) error
	}

	MarksRepository interface {
		GetStudentsAttestationMarksTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Marks, error)
		GetStudentsExamResults(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Exams, error)
		GetStudentsSupervisorMarks(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.SupervisorMarks, error)
		UpsertStudentsSupervisorMark(ctx context.Context, tx pgx.Tx, model model.SupervisorMarks) error
	}
)

type Service struct {
	dissertationRepo DissertationRepository
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

func (s *Service) InitSupervisor(ctx context.Context, user model.Users, registry request_models.FirstSupervisorRegistry) error {
	supervisor := model.Supervisors{
		SupervisorID: user.KasperID,
		UserID:       user.UserID,
		FullName:     lo.ToPtr(registry.FullName),
		Phone:        registry.Phone,
		Faculty:      lo.ToPtr(registry.Faculty),
		Department:   lo.ToPtr(registry.Department),
		Degree:       lo.ToPtr(registry.Degree),
	}

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		return s.client.InsertSupervisor(ctx, tx, supervisor)
	}); err != nil {
		return errors.Wrap(err, "InitSupervisor()")
	}

	return nil
}
