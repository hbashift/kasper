package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"
	adminmap "uir_draft/internal/pkg/service/admin/mapping"
	"uir_draft/internal/pkg/service/student/mapping"
)

var ErrNonValidToken = errors.New("token is expired")

type StudentSupervisorRepository interface {
	GetPairs(ctx context.Context, tx *pgxpool.Pool) ([]*adminmap.StudentSupervisorPair, error)
	ChangeSupervisor(ctx context.Context, tx *pgxpool.Pool, pairs []*adminmap.ChangeSupervisor) error
}

type StudentRepository interface {
	GetNumberOfYears(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (int32, error)
	SetAcademicLeave(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, isAcademicLeave bool) error
	UpdateFeedback(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, feedback string) error
	GetStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (*models.StudentCommonInformation, error)
	GetListOfStudents(ctx context.Context, tx *pgxpool.Pool, supervisorID *uuid.UUID) ([]*model.Students, error)
	UpdateStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, student model.Students) error
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

type ScientificWorksRepository interface {
	GetScientificWorks(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.ScientificWork, error)
}

type TeachingLoadRepo interface {
	GetStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.TeachingLoad, error)
}

type SupervisorRepository interface {
	GetSupervisors(ctx context.Context, tx *pgxpool.Pool) ([]*model.Supervisors, error)
}

type Service struct {
	studRepo     StudentRepository
	tokenRepo    TokenRepository
	semesterRepo SemesterRepository
	dRepo        DissertationRepository
	scienceRepo  ScientificWorksRepository
	loadRepo     TeachingLoadRepo
	studSupRepo  StudentSupervisorRepository
	supRepo      SupervisorRepository

	db *pgxpool.Pool
}

func NewService(studRepo StudentRepository, tokenRepo TokenRepository, semesterRepo SemesterRepository, dRepo DissertationRepository, scienceRepo ScientificWorksRepository, loadRepo TeachingLoadRepo, studSupRepo StudentSupervisorRepository, supRepo SupervisorRepository, db *pgxpool.Pool) *Service {
	return &Service{studRepo: studRepo, tokenRepo: tokenRepo, semesterRepo: semesterRepo, dRepo: dRepo, scienceRepo: scienceRepo, loadRepo: loadRepo, studSupRepo: studSupRepo, supRepo: supRepo, db: db}
}
