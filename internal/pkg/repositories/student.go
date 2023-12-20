package repositories

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
	"uir_draft/internal/pkg/models"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type StudentRepository struct {
	postgres *pgxpool.Pool
}

func NewStudentRepository(postgres *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{postgres: postgres}
}

func (r *StudentRepository) GetStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) (*models.StudentCommonInformation, error) {
	commonInfo, err := r.getStudentCommonInformation(ctx, tx, clientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentDissertationPlan()")
	}

	return commonInfo, nil
}

func (r *StudentRepository) getStudentCommonInformation(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) (*models.StudentCommonInformation, error) {
	stmt, args := table.Students.
		INNER_JOIN(table.Dissertation, table.Students.StudentID.EQ(table.Dissertation.StudentID)).
		INNER_JOIN(table.Supervisors, table.Students.SupervisorID.EQ(table.Supervisors.SupervisorID)).
		SELECT(
			table.Students.DissertationTitle.AS("dissertation_title"),
			table.Supervisors.FullName.AS("supervisor_name"),
			table.Students.EnrollmentOrder.AS("enrollment_order_number"),
			table.Students.StartDate.AS("studying_start_date"),
			table.Students.ActualSemester.AS("semester_number"),
			table.Dissertation.Feedback.AS("feedback"),
			table.Dissertation.Status.AS("dissertation_status"),
			table.Students.TitlePagePath.AS("title_page_url"),
			table.Students.ExplanatoryNoteURL.AS("explanatory_note_url"),
		).
		WHERE(table.Students.ClientID.EQ(postgres.UUID(clientID))).Sql()

	var studentCommonInfo models.StudentCommonInformation

	row := tx.QueryRow(ctx, stmt, args...)

	if err := scanStudentCommonInfo(row, &studentCommonInfo); err != nil {
		return nil, errors.Wrap(err, "mapping student common info")
	}

	return &studentCommonInfo, nil
}

func (r *StudentRepository) InsertStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, student model.Students) error {
	if err := r.insertStudentCommonInfoTx(ctx, tx, student); err != nil {
		return errors.Wrap(err, "InsertStudentCommonInfo(): error during transaction")
	}

	return nil
}

func (r *StudentRepository) insertStudentCommonInfoTx(ctx context.Context, tx *pgxpool.Pool, student model.Students) error {
	// TODO ограничение по столбцам
	stmt, args := table.Students.
		INSERT(table.Students.AllColumns).
		MODEL(student).Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "insert student common info")
	}

	return nil
}

func (r *StudentRepository) UpdateStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, student model.Students) error {
	if err := r.updateStudentCommonInfoTx(ctx, tx, student); err != nil {
		return errors.Wrap(err, "UpdateStudentCommonInfo(): error during transaction")
	}

	return nil
}

func (r *StudentRepository) updateStudentCommonInfoTx(ctx context.Context, tx *pgxpool.Pool, student model.Students) error {
	stmt, args := table.Students.
		UPDATE(
			table.Students.FullName,
			table.Students.Specialization,
			table.Students.AcademicLeave,
			table.Students.DissertationTitle,
		).
		MODEL(student).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(student.StudentID))).Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "update student common info")
	}

	return nil
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
