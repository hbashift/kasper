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

func (r *StudentRepository) GetStudentCommonInfo(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (*models.StudentCommonInformation, error) {
	commonInfo, err := r.getStudentCommonInformation(ctx, tx, studentID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentDissertationPlan()")
	}

	return commonInfo, nil
}

func (r *StudentRepository) getStudentCommonInformation(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (*models.StudentCommonInformation, error) {
	stmt, args := table.Students.
		INNER_JOIN(table.Dissertation, table.Students.StudentID.EQ(table.Dissertation.StudentID)).
		INNER_JOIN(table.Supervisors, table.Students.SupervisorID.EQ(table.Supervisors.SupervisorID)).
		SELECT(
			table.Students.DissertationTitle.AS("dissertation_title"),
			table.Supervisors.FullName.AS("supervisor_name"),
			table.Students.EnrollmentOrder.AS("enrollment_order_number"),
			table.Students.StartDate.AS("studying_start_date"),
			table.Students.ActualSemester.AS("semester_number"),
			table.Students.Feedback.AS("feedback"),
			table.Students.NumberOfYears.AS("number_of_years"),
		).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

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

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
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

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "update student common info")
	}

	return nil
}

func (r *StudentRepository) GetListOfStudents(ctx context.Context, tx *pgxpool.Pool, supervisorID *uuid.UUID) ([]*model.Students, error) {
	list, err := r.getListOfStudentsTx(ctx, tx, supervisorID)
	if err != nil {
		return nil, errors.Wrap(err, "GetListOfStudents()")
	}

	return list, nil
}

func (r *StudentRepository) getListOfStudentsTx(ctx context.Context, tx *pgxpool.Pool, supervisorID *uuid.UUID) ([]*model.Students, error) {
	stmt, args := table.Students.
		SELECT(table.Students.AllColumns).
		WHERE(table.Students.SupervisorID.EQ(postgres.UUID(supervisorID))).
		Sql()

	row, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "getListOfStudentsTx()")
	}

	var studentsList []*model.Students

	for row.Next() {
		var studentCommonInfo model.Students
		if err = scanStudentRow(row, &studentCommonInfo); err != nil {
			return nil, errors.Wrap(err, "getListOfStudentsTx()")
		}

		studentsList = append(studentsList, &studentCommonInfo)
	}

	return studentsList, nil
}

func (r *StudentRepository) UpdateFeedback(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, feedback string) error {
	stmt, args := table.Students.
		UPDATE(table.Students.Feedback).
		SET(feedback).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "UpdateFeedback()")
	}

	return nil
}

func (r *StudentRepository) SetAcademicLeave(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, isAcademicLeave bool) error {
	stmt, args := table.Students.
		UPDATE(table.Students.AcademicLeave).
		SET(isAcademicLeave).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "SetAcademicLeave()")
	}

	return nil
}

func (r *StudentRepository) GetNumberOfYears(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) (int32, error) {
	stmt, args := table.Students.
		SELECT(table.Students.NumberOfYears).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)
	years := int32(0)

	if err := row.Scan(&years); err != nil {
		return 0, errors.Wrap(err, "GetNumberOfYears()")
	}

	return years, nil
}

func scanStudentCommonInfo(row pgx.Row, target *models.StudentCommonInformation) error {
	return row.Scan(
		&target.DissertationTitle,
		&target.SupervisorName,
		&target.EnrollmentOrderNumber,
		&target.StudyingStartDate,
		&target.Semester,
		&target.Feedback,
		&target.NumberOfYears,
	)
}

func scanStudentRow(row pgx.Row, target *model.Students) error {
	return row.Scan(
		&target.ClientID,
		&target.StudentID,
		&target.FullName,
		&target.Department,
		&target.EnrollmentOrder,
		&target.Specialization,
		&target.ActualSemester,
		&target.SupervisorID,
		&target.StartDate,
		&target.AcademicLeave,
		&target.DissertationTitle,
		&target.Feedback,
		&target.GroupNumber,
		&target.NumberOfYears,
	)
}
