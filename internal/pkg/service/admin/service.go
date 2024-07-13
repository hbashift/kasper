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
	ClientRepository interface {
		GetSupervisorsTx(ctx context.Context, tx pgx.Tx) ([]models.Supervisor, error)
		SetNewSupervisorTx(ctx context.Context, tx pgx.Tx, studentID, supervisorID uuid.UUID) error
		GetStudentSupervisorPairsTx(ctx context.Context, tx pgx.Tx) ([]models.StudentSupervisorPair, error)
		GetStudentTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (model.Students, error)
		SetStudentStatusTx(ctx context.Context, tx pgx.Tx, status model.ApprovalStatus, studyingStatus model.StudentStatus, studentID uuid.UUID) error
		GetStudentsList(ctx context.Context, tx pgx.Tx) ([]models.Student, error)
		SetStudentFlags(ctx context.Context, tx pgx.Tx, studyingStatus model.StudentStatus, canEdit bool, studentID uuid.UUID) error
		ArchiveSupervisor(ctx context.Context, tx pgx.Tx, supervisors []models.SupervisorStatus) error
		UpdateStudentsSemester(ctx context.Context, tx pgx.Tx, students []model.Students) error
		GetStudentsByStudentsIDs(ctx context.Context, tx pgx.Tx, studentIDs []uuid.UUID) ([]model.Students, error)
		GetAllStudentIDs(ctx context.Context, tx pgx.Tx) ([]uuid.UUID, error)
		GetDataForReportOne(ctx context.Context, tx pgx.Tx, studentIDs []uuid.UUID) ([]models.StudentInfoForReportOne, error)
		GetDataForReportTwo(ctx context.Context, tx pgx.Tx, studentIDs []uuid.UUID) ([]models.StudentInfoForReportTwo, error)
		GetStudentsActualSupervisorTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (models.Supervisor, error)
	}

	UsersRepository interface {
		InsertUsersTx(ctx context.Context, tx pgx.Tx, users []model.Users) error
		GetNotRegisteredUsers(ctx context.Context, tx pgx.Tx) ([]models.UserInfo, error)
		DeleteNotRegisteredUsers(ctx context.Context, tx pgx.Tx, userIDs []uuid.UUID) error
	}

	MarksRepository interface {
		UpsertAttestationMarksTx(ctx context.Context, tx pgx.Tx, models []model.Marks) error
	}
)

type Service struct {
	clientRepo ClientRepository
	marksRepo  MarksRepository
	userRepo   UsersRepository
	db         *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{
		clientRepo: repository.NewClientRepository(),
		marksRepo:  repository.NewMarksRepository(),
		userRepo:   repository.NewUsersRepository(),
		db:         db,
	}
}
