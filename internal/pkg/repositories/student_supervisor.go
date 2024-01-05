package repositories

import (
	"context"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
	"uir_draft/internal/pkg/service/admin/mapping"
)

type StudentSupervisorRepository struct {
}

func NewStudentSupervisorRepository() *StudentSupervisorRepository {
	return &StudentSupervisorRepository{}
}

func (r *StudentSupervisorRepository) ChangeSupervisor(ctx context.Context, tx *pgxpool.Pool, studentID, supervisorID uuid.UUID) error {
	stmt, args := table.StudentSupervisor.
		UPDATE(table.StudentSupervisor.SupervisorID).
		SET(supervisorID).
		WHERE(table.StudentSupervisor.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "ChangeSupervisor()")
	}

	return nil
}

func (r *StudentSupervisorRepository) SetStudentSupervisor(ctx context.Context, tx *pgxpool.Pool, model model.StudentSupervisor) error {
	stmt, args := table.StudentSupervisor.
		INSERT(table.StudentSupervisor.MutableColumns).
		MODEL(model).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "SetStudentSupervisor()")
	}

	return nil
}

func (r *StudentSupervisorRepository) GetPairs(ctx context.Context, tx *pgxpool.Pool) ([]*mapping.StudentSupervisorPair, error) {
	stmt, args := table.Students.
		INNER_JOIN(table.StudentSupervisor, table.StudentSupervisor.StudentID.EQ(table.Students.StudentID)).
		INNER_JOIN(table.Supervisors, table.Supervisors.SupervisorID.EQ(table.StudentSupervisor.SupervisorID)).
		SELECT(
			table.Supervisors.SupervisorID.AS("supervisor_id"),
			table.Supervisors.FullName.AS("supervisor_name"),
			table.Students.StudentID.AS("student_id"),
			table.Students.FullName.AS("student_name"),
		).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetPairs()")
	}

	var pairs []*mapping.StudentSupervisorPair

	for rows.Next() {
		pair := &mapping.StudentSupervisorPair{}
		if err = scanPairsRow(rows, pair); err != nil {
			return nil, err
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}

func scanPairsRow(row pgx.Row, target *mapping.StudentSupervisorPair) error {
	return row.Scan(
		&target.SupervisorID,
		&target.SupervisorName,
		&target.StudentID,
		&target.StudentName,
	)
}
