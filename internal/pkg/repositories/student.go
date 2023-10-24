package repositories

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/table"
	"uir_draft/internal/pkg/models"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type studentRepository struct {
	postgres *pgxpool.Pool
}

func NewStudentRepository(postgres *pgxpool.Pool) *studentRepository {
	return &studentRepository{postgres: postgres}
}

func (r *studentRepository) GetDissertation(ctx context.Context, clientID uuid.UUID) (*models.DissertationPage, error) {
	plan, err := r.getStudentDissertationPlan(ctx, clientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetDissertation():")
	}

	commonInfo, err := r.getStudentCommonInformation(ctx, clientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetDissertation():")
	}

	return &models.DissertationPage{
		DissertationPlan: plan,
		CommonInfo:       *commonInfo,
	}, nil
}

func (r *studentRepository) getStudentDissertationPlan(ctx context.Context, clientID uuid.UUID) ([]*models.StudentDissertationPlan, error) {
	stmt, args := table.SemesterProgress.
		SELECT(
			table.SemesterProgress.ProgressName.AS("name"),
			table.SemesterProgress.First,
			table.SemesterProgress.Second,
			table.SemesterProgress.Third,
			table.SemesterProgress.Forth,
			table.SemesterProgress.Fifth,
			table.SemesterProgress.Sixth,
		).
		WHERE(table.SemesterProgress.ClientID.EQ(postgres.UUID(clientID))).Sql()

	studentPlan := make([]*models.StudentDissertationPlan, 0)

	rows, err := r.postgres.Query(ctx, stmt, args...)
	defer rows.Close()

	if err != nil {
		return nil, errors.Wrap(err, "selecting students dissertation plan")
	}

	for rows.Next() {
		var plan = &models.StudentDissertationPlan{}
		if err := scanDissertationPlan(rows, plan); err != nil {
			return nil, errors.Wrap(err, "mapping dissertation plan rows")
		}
		studentPlan = append(studentPlan, plan)
	}

	return studentPlan, nil
}

func (r *studentRepository) getStudentCommonInformation(ctx context.Context, clientID uuid.UUID) (*models.StudentCommonInformation, error) {
	stmt, args := table.Students.
		INNER_JOIN(table.Dissertation, table.Students.StudentID.EQ(table.Dissertation.StudentID)).
		INNER_JOIN(table.Supervisors, table.Students.SupervisorID.EQ(table.Supervisors.SupervisorID)).
		SELECT(
			table.Dissertation.Title.AS("dissertation_title"),
			table.Supervisors.FullName.AS("supervisor_name"),
			table.Students.EnrollmentOrder.AS("enrollment_order_number"),
			table.Students.StartDate.AS("studying_start_date"),
			table.Students.ActualSemester.AS("semester"),
			table.Dissertation.Feedback.AS("feedback"),
			table.Dissertation.Status.AS("dissertation_status"),
			table.Students.TitlePagePath.AS("title_page_url"),
			table.Students.ExplanatoryNoteURL.AS("explanatory_note_url"),
		).
		WHERE(table.Students.ClientID.EQ(postgres.UUID(clientID))).Sql()

	var studentCommonInfo models.StudentCommonInformation

	row := r.postgres.QueryRow(ctx, stmt, args...)

	if err := scanStudentCommonInfo(row, &studentCommonInfo); err != nil {
		return nil, errors.Wrap(err, "mapping student common info")
	}

	return &studentCommonInfo, nil
}

func scanDissertationPlan(row pgx.Row, target *models.StudentDissertationPlan) error {
	return row.Scan(
		&target.Name,
		&target.First,
		&target.Second,
		&target.Third,
		&target.Forth,
		&target.Fifth,
		&target.Sixth,
	)
}

func scanStudentCommonInfo(row pgx.Row, target *models.StudentCommonInformation) error {
	return row.Scan(
		&target.DissertationTitle,
		&target.SupervisorName,
		&target.EnrollmentOrderNumber,
		&target.StudyingStartDate,
		&target.Semester,
		&target.Feedback,
		&target.DissertationStatus,
		&target.TitlePageURL,
		&target.ExplanatoryNoteURL,
	)
}
