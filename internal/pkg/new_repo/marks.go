package new_repo

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type MarksRepository struct{}

func NewMarksRepository() *MarksRepository {
	return &MarksRepository{}
}

func (r *MarksRepository) GetStudentMarksTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Marks, error) {
	stmt, args := table.Marks.
		SELECT(table.Marks.AllColumns).
		WHERE(table.Marks.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentMarksTx()")
	}
	defer rows.Close()

	marks := make([]model.Marks, 0, 8)

	for rows.Next() {
		mark := model.Marks{}

		if err := scanMarks(rows, &mark); err != nil {
			return nil, errors.Wrap(err, "GetStudentMarksTx(): scanning row")
		}

		marks = append(marks, mark)
	}

	return marks, nil
}

func (r *MarksRepository) UpsertMarkTx(ctx context.Context, tx pgx.Tx, model model.Marks) error {
	stmt, args := table.Marks.
		INSERT().
		MODELS(model).
		ON_CONFLICT(table.Marks.StudentID, table.Marks.Semester).
		DO_UPDATE(postgres.
			SET(table.Marks.Mark.SET(postgres.Int32(model.Mark)))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpsertMarkTx()")
	}

	return nil
}

func scanMarks(row pgx.Row, target *model.Marks) error {
	return row.Scan(
		&target.StudentID,
		&target.Mark,
		&target.Semester,
	)
}
