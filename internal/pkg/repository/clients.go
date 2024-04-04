package repository

import (
	"context"
	"time"

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
		INSERT(
			table.Students.AllColumns.
				Except(
					table.Students.StudyingStatus,
					table.Students.Status,
					table.Students.CanEdit,
				),
		).
		VALUES(
			student.StudentID,
			student.UserID,
			student.FullName,
			//student.Department,
			student.SpecID,
			student.ActualSemester,
			student.Years,
			student.StartDate,
			student.GroupID,
		).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertStudentTx()")
	}

	return nil
}

func (r *ClientRepository) SetStudentStatusTx(ctx context.Context, tx pgx.Tx, status model.ApprovalStatus, studyingStatus model.StudentStatus, studentID uuid.UUID) error {
	stmt, args := table.Students.
		UPDATE(
			table.Students.Status,
			table.Students.StudyingStatus,
		).
		SET(
			status,
			studyingStatus,
		).
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
		FROM(table.Students.
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

	list := make([]models.Student, 0)

	for rows.Next() {
		el := models.Student{}
		if err = scanStudentList(rows, &el); err != nil {
			return nil, errors.Wrap(err, "GetSupervisorsStudentsTx(): scanning rows")
		}

		list = append(list, el)
	}

	return list, err
}

func (r *ClientRepository) GetStudentSupervisorPairsTx(ctx context.Context, tx pgx.Tx) ([]models.StudentSupervisorPair, error) {
	stmt, args := table.Students.
		SELECT(
			table.Students.AllColumns.Except(table.Students.UserID, table.Students.SpecID, table.Students.GroupID),
			table.Specializations.Title,
			table.Groups.GroupName,
			table.Supervisors.AllColumns.Except(table.Supervisors.UserID),
		).
		FROM(table.Students.
			LEFT_JOIN(table.StudentsSupervisors, table.StudentsSupervisors.StudentID.EQ(table.Students.StudentID)).
			LEFT_JOIN(table.Supervisors, table.StudentsSupervisors.SupervisorID.EQ(table.Supervisors.SupervisorID)).
			INNER_JOIN(table.Groups, table.Students.GroupID.EQ(table.Groups.GroupID)).
			INNER_JOIN(table.Specializations, table.Students.SpecID.EQ(table.Specializations.SpecID)),
		).
		WHERE(table.StudentsSupervisors.EndAt.IS_NULL()).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentSupervisorPairsTx()")
	}
	defer rows.Close()

	list := make([]models.StudentSupervisorPair, 0)

	for rows.Next() {
		el := models.StudentSupervisorPair{}
		if err = scanStudentSupervisorPair(rows, &el); err != nil {
			return nil, errors.Wrap(err, "GetStudentSupervisorPairsTx(): scanning rows")
		}

		list = append(list, el)
	}

	return list, nil
}

func (r *ClientRepository) SetNewSupervisorTx(ctx context.Context, tx pgx.Tx, studentID, supervisorID uuid.UUID) error {
	insertStmt, instArgs := table.StudentsSupervisors.
		INSERT(
			table.StudentsSupervisors.ID,
			table.StudentsSupervisors.SupervisorID,
			table.StudentsSupervisors.StudentID,
		).
		VALUES(
			uuid.New(),
			supervisorID,
			studentID,
		).
		Sql()

	updateStmt, updArgs := table.StudentsSupervisors.
		UPDATE(table.StudentsSupervisors.EndAt).
		SET(time.Now()).
		WHERE(
			table.StudentsSupervisors.EndAt.IS_NULL().
				AND(table.StudentsSupervisors.StudentID.EQ(postgres.UUID(studentID))),
		).
		Sql()

	_, err := tx.Exec(ctx, updateStmt, updArgs...)
	if err != nil {
		return errors.Wrap(err, "SetNewSupervisorTx(): update")
	}

	_, err = tx.Exec(ctx, insertStmt, instArgs...)
	if err != nil {
		return errors.Wrap(err, "SetNewSupervisorTx(): insert")
	}

	return nil
}

func (r *ClientRepository) GetSupervisorsTx(ctx context.Context, tx pgx.Tx) ([]models.Supervisor, error) {
	stmt, args := table.Supervisors.
		SELECT(
			table.Supervisors.SupervisorID,
			table.Supervisors.FullName,
		).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetSupervisorsTx()")
	}
	defer rows.Close()

	supervisors := make([]models.Supervisor, 0)

	for rows.Next() {
		el := models.Supervisor{}
		if err = scanSupervisor(rows, &el); err != nil {
			return nil, errors.Wrap(err, "GetSupervisorsTx()")
		}

		supervisors = append(supervisors, el)
	}

	return supervisors, nil
}

func (r *ClientRepository) GetStudentsActualSupervisorTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) (models.Supervisor, error) {
	stmt, args := table.Supervisors.
		SELECT(
			table.Supervisors.SupervisorID,
			table.Supervisors.FullName,
		).
		FROM(
			table.StudentsSupervisors.
				INNER_JOIN(table.Supervisors, table.StudentsSupervisors.SupervisorID.EQ(table.Supervisors.SupervisorID)),
		).
		WHERE(table.StudentsSupervisors.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.StudentsSupervisors.EndAt.IS_NULL())).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)
	supervisor := models.Supervisor{}
	if err := scanSupervisor(row, &supervisor); err != nil {
		return models.Supervisor{}, errors.Wrap(err, "GetStudentsActualSupervisorTx()")
	}

	return supervisor, nil
}

func (r *ClientRepository) GetSupervisorTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) (models.Supervisor, error) {
	stmt, args := table.Supervisors.
		SELECT(
			table.Supervisors.SupervisorID,
			table.Supervisors.FullName,
		).
		WHERE(table.Supervisors.SupervisorID.EQ(postgres.UUID(supervisorID))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)
	supervisor := models.Supervisor{}
	if err := scanSupervisor(row, &supervisor); err != nil {
		return models.Supervisor{}, errors.Wrap(err, "GetSupervisorTx()")
	}

	return supervisor, nil
}

func (r *ClientRepository) GetAllStudentsSupervisors(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]models.SupervisorFull, error) {
	stmt, args := table.Supervisors.
		SELECT(
			table.Supervisors.SupervisorID,
			table.Supervisors.FullName,
			table.StudentsSupervisors.StartAt,
			table.StudentsSupervisors.EndAt,
		).
		FROM(table.Supervisors.
			INNER_JOIN(table.StudentsSupervisors, table.Supervisors.SupervisorID.EQ(table.StudentsSupervisors.SupervisorID))).
		WHERE(table.StudentsSupervisors.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllStudentsSupervisors()")
	}
	defer rows.Close()

	supervisors := make([]models.SupervisorFull, 0)

	for rows.Next() {
		supervisor := models.SupervisorFull{}
		if err := scanSupervisorFull(rows, &supervisor); err != nil {
			return nil, errors.Wrap(err, "GetAllStudentsSupervisors(): scanning")
		}

		supervisors = append(supervisors, supervisor)
	}

	return supervisors, nil
}

// TODO доделать для аспера и научника

func scanStudent(row pgx.Row, target *model.Students) error {
	return row.Scan(
		&target.StudentID,
		&target.UserID,
		&target.FullName,
		//&target.Department,
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
		//&target.Department,
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

func scanStudentSupervisorPair(row pgx.Row, target *models.StudentSupervisorPair) error {
	return row.Scan(
		&target.Student.StudentID,
		&target.Student.FullName,
		//&target.Student.Department,
		&target.Student.ActualSemester,
		&target.Student.Years,
		&target.Student.StartDate,
		&target.Student.StudyingStatus,
		&target.Student.Status,
		&target.Student.CanEdit,
		&target.Student.Specialization,
		&target.Student.GroupName,
		&target.Supervisor.SupervisorID,
		&target.Supervisor.FullName,
	)
}

func scanSupervisor(row pgx.Row, target *models.Supervisor) error {
	return row.Scan(
		&target.SupervisorID,
		&target.FullName,
	)
}

func scanSupervisorFull(row pgx.Row, target *models.SupervisorFull) error {
	return row.Scan(
		&target.SupervisorID,
		&target.FullName,
		&target.StartAt,
		&target.EndAt,
	)
}
