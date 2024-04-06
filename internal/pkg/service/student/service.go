package student

import (
	"context"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

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
		SetUserRegisteredTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) error
	}

	MarksRepository interface {
		GetStudentMarksTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Marks, error)
	}

	StudentRepository interface {
		GetStudentTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (model.Students, error)
		SetStudentStatusTx(ctx context.Context, tx pgx.Tx, status model.ApprovalStatus, studyingStatus model.StudentStatus, studentID uuid.UUID) error
		GetStudentStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (models.Student, error)
		InsertStudentTx(ctx context.Context, tx pgx.Tx, student model.Students) error
		SetNewSupervisorTx(ctx context.Context, tx pgx.Tx, studentID, supervisorID uuid.UUID) error
		GetAllStudentsSupervisors(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]models.SupervisorFull, error)
	}

	DissertationRepository interface {
		SetSemesterProgressStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, acceptedAt *time.Time) error
		SetDissertationStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32) error
		SetDissertationTitleStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32, acceptedAt *time.Time) error

		GetSemesterProgressTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.SemesterProgress, error)
		UpsertSemesterProgressTx(ctx context.Context, tx pgx.Tx, progresses []model.SemesterProgress) error

		GetActualDissertationData(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, semester int32) (model.Dissertations, error)
		GetDissertationsTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Dissertations, error)
		UpsertDissertationTx(ctx context.Context, tx pgx.Tx, model model.Dissertations) error

		GetDissertationTitlesTx(ctx context.Context, tx pgx.Tx, dissertationID uuid.UUID) ([]model.DissertationTitles, error)
		InsertDissertationTitleTx(ctx context.Context, tx pgx.Tx, title model.DissertationTitles) error

		GetFeedbackTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Feedback, error)
	}

	ScientificRepository interface {
		SetScientificWorkStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32, acceptedAt *time.Time) error

		GetScientificWorksStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.ScientificWorksStatus, error)
		UpdateScientificWorksStatusTx(ctx context.Context, tx pgx.Tx, works model.ScientificWorksStatus) error

		InsertPublicationsTx(ctx context.Context, tx pgx.Tx, publications []model.Publications) error
		UpdatePublicationsTx(ctx context.Context, tx pgx.Tx, publications []model.Publications) error
		DeletePublicationsTx(ctx context.Context, tx pgx.Tx, publicationsIDs []uuid.UUID) error

		InsertConferencesTx(ctx context.Context, tx pgx.Tx, conferences []model.Conferences) error
		UpdateConferencesTx(ctx context.Context, tx pgx.Tx, conferences []model.Conferences) error
		DeleteConferencesTx(ctx context.Context, tx pgx.Tx, conferencesIDs []uuid.UUID) error

		InsertResearchProjectsTx(ctx context.Context, tx pgx.Tx, projects []model.ResearchProjects) error
		UpdateResearchProjectsTx(ctx context.Context, tx pgx.Tx, projects []model.ResearchProjects) error
		DeleteResearchProjectsTx(ctx context.Context, tx pgx.Tx, projectsIDs []uuid.UUID) error

		GetScientificWorksTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]models.ScientificWork, error)
		InitScientificWorkStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) error
		GetScientificWorksStatusBySemesterTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, semester int32) (model.ScientificWorksStatus, error)
	}

	TeachingLoadRepository interface {
		SetTeachingLoadStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32, acceptedAt *time.Time) error

		GetTeachingLoadStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.TeachingLoadStatus, error)
		UpdateTeachingLoadStatusTx(ctx context.Context, tx pgx.Tx, loads []model.TeachingLoadStatus) error

		InsertClassroomLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.ClassroomLoad) error
		UpdateClassroomLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.ClassroomLoad) error
		DeleteClassroomLoadsTx(ctx context.Context, tx pgx.Tx, classroomsIDs []uuid.UUID) error

		InsertIndividualLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.IndividualStudentsLoad) error
		UpdateIndividualLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.IndividualStudentsLoad) error
		DeleteIndividualStudentsLoadsTx(ctx context.Context, tx pgx.Tx, individualsIDs []uuid.UUID) error

		InsertAdditionalLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.AdditionalLoad) error
		UpdateAdditionalLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.AdditionalLoad) error
		DeleteAdditionalLoadsTx(ctx context.Context, tx pgx.Tx, additionalIDs []uuid.UUID) error

		GetTeachingLoadsTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]models.TeachingLoad, error)
		InitTeachingLoadsStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) error
		GetTeachingLoadStatusBySemesterTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, semester int32) (model.TeachingLoadStatus, error)
	}
)

type Service struct {
	dissertationRepo DissertationRepository
	loadRepo         TeachingLoadRepository
	scienceRepo      ScientificRepository
	marksRepo        MarksRepository
	studRepo         StudentRepository
	tokenRepo        TokenRepository
	userRepo         UsersRepository
	db               *pgxpool.Pool
}

func NewService(dissertationRepo DissertationRepository, loadRepo TeachingLoadRepository, scienceRepo ScientificRepository, marksRepo MarksRepository, studRepo StudentRepository, tokenRepo TokenRepository, userRepo UsersRepository, db *pgxpool.Pool) *Service {
	return &Service{dissertationRepo: dissertationRepo, loadRepo: loadRepo, scienceRepo: scienceRepo, marksRepo: marksRepo, studRepo: studRepo, tokenRepo: tokenRepo, userRepo: userRepo, db: db}
}
