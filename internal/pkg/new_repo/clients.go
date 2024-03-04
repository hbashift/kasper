package new_repo

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"
	"uir_draft/internal/pkg/models"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type ClientRepository struct{}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{}
}

func (r *ClientRepository) GetStudentTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (model.Students, error) {
	stmt, args := table.Students.
		SELECT(table.Students.AllColumns).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)
	student := model.Students{}

	if err := scanStudent(row, &student); err != nil {
		return model.Students{}, errors.Wrap(err, "GetStudentTx()")
	}

	return student, nil
}

func (r *ClientRepository) GetStudentStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (models.Student, error) {
	stmt, args := table.Students.
		SELECT(
			table.Students.AllColumns.Except(table.Students.UserID, table.Students.SpecID, table.Students.GroupID),
			table.Specializations.Title,
			table.Groups.GroupName,
		).
		FROM(table.Students.
			INNER_JOIN(table.Groups, table.Students.GroupID.EQ(table.Groups.GroupID)).
			INNER_JOIN(table.Specializations, table.Students.SpecID.EQ(table.Specializations.SpecID)),
		).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)
	student := models.Student{}
	if err := scanStudentList(row, &student); err != nil {
		return models.Student{}, errors.Wrap(err, "GetStudentStatusTx()")
	}

	return student, nil
}

func (r *ClientRepository) InsertStudentTx(ctx context.Context, tx pgx.Tx, student model.Students) error {
	stmt, args := table.Students.
		INSERT().
		MODEL(student).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertStudentTx()")
	}

	return nil
}

func (r *ClientRepository) SetStudentStatusTx(ctx context.Context, tx pgx.Tx, status model.ApprovalStatus, studentID uuid.UUID) error {
	stmt, args := table.Students.
		UPDATE(table.Students.Status).
		SET(status).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "SetStudentStatusTx()")
	}

	return nil
}

func (r *ClientRepository) GetSupervisorsStudentsTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) ([]models.Student, error) {
	stmt, args := table.Students.
		SELECT(
			table.Students.AllColumns.Except(table.Students.UserID, table.Students.SpecID, table.Students.GroupID),
			table.Specializations.Title,
			table.Groups.GroupName,
		).
		FROM(table.StudentsSupervisors.
			INNER_JOIN(table.StudentsSupervisors, table.Students.StudentID.EQ(table.StudentsSupervisors.StudentID)).
			INNER_JOIN(table.Groups, table.Students.GroupID.EQ(table.Groups.GroupID)).
			INNER_JOIN(table.Specializations, table.Students.SpecID.EQ(table.Specializations.SpecID)),
		).
		WHERE(table.StudentsSupervisors.SupervisorID.EQ(postgres.UUID(supervisorID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetSupervisorsStudentsTx()")
	}
	defer rows.Close()

	list := make([]models.Student, 0, 0)

	for rows.Next() {
		el := models.Student{}
		if err = scanStudentList(rows, &el); err != nil {
			return nil, errors.Wrap(err, "GetSupervisorsStudentsTx(): scanning rows")
		}

		list = append(list, el)
	}

	return list, err
}

// TODO доделать для аспера и научника

func scanStudent(row pgx.Row, target *model.Students) error {
	return row.Scan(
		&target.StudentID,
		&target.UserID,
		&target.FullName,
		&target.Department,
		&target.SpecID,
		&target.ActualSemester,
		&target.Years,
		&target.StartDate,
		&target.StudyingStatus,
		&target.GroupID,
		&target.Status,
		&target.CanEdit,
	)
}

func scanStudentList(row pgx.Row, target *models.Student) error {
	return row.Scan(
		&target.StudentID,
		&target.FullName,
		&target.Department,
		&target.ActualSemester,
		&target.Years,
		&target.StartDate,
		&target.StudyingStatus,
		&target.Status,
		&target.CanEdit,
		&target.Specialization,
		&target.GroupName,
	)
}
