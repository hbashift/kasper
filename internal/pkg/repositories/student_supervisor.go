package repositories

import (
	"context"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
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
