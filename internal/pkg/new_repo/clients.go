package new_repo

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type ClientRepository struct{}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{}
}

func (r *ClientRepository) GetStudentStatusTx(ctx context.Context, tx *pgxpool.Tx, studentID uuid.UUID) (model.Students, error) {
	stmt, args := table.Students.
		SELECT(table.Students.AllColumns).
		WHERE(table.Students.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)
	student := model.Students{}

	if err := scanStudent(row, &student); err != nil {
		return model.Students{}, errors.Wrap(err, "GetStudentStatusTx()")
	}

	return student, nil
}

func (r *ClientRepository) InsertStudentTx(ctx context.Context, tx *pgxpool.Tx, student model.Students) error {
	stmt, args := table.Students.
		INSERT().
		MODEL(student).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertStudentTx()")
	}

	return nil
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
		&target.StatusID,
		&target.GroupID,
		&target.Status,
		&target.CanEdit,
	)
}
