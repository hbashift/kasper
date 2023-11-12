package repositories

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type ScientificWork struct {
	postgres *pgxpool.Pool
}

func NewScientificWork(postgres *pgxpool.Pool) *ScientificWork {
	return &ScientificWork{postgres: postgres}
}

func (r *ScientificWork) GetScientificWorks(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.ScientificWork, error) {
	result, err := r.getScientificWorksTx(ctx, tx, studentID)
	if err != nil {
		return nil, errors.Wrap(err, "GetScientificWorks()")
	}

	return result, nil
}

func (r *ScientificWork) getScientificWorksTx(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.ScientificWork, error) {
	var works []*model.ScientificWork

	stmt, args := table.ScientificWork.
		SELECT(table.ScientificWork.AllColumns).
		WHERE(table.ScientificWork.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, stmt, args...)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Err(); err != nil {
				return err
			}

			work := &model.ScientificWork{}

			if err := scanScientificWork(rows, work); err != nil {
				return err
			}

			works = append(works, work)
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "get scientific works tx")
	}

	return works, nil
}

func (r *ScientificWork) InsertStudentScientificWorks(ctx context.Context, tx *pgxpool.Pool, work *model.ScientificWork) error {
	if err := r.insertStudentScientificWorksTx(ctx, tx, work); err != nil {
		return errors.Wrap(err, "InsertStudentScientificWorks()")
	}

	return nil
}

func (r *ScientificWork) insertStudentScientificWorksTx(ctx context.Context, tx *pgxpool.Pool, work *model.ScientificWork) error {
	stmt, args := table.ScientificWork.
		INSERT(table.ScientificWork.AllColumns).
		MODEL(work).
		ON_CONFLICT(table.ScientificWork.StudentID, table.ScientificWork.Name).
		DO_NOTHING().
		Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "insertStudentScientificWorksTx()")
	}

	return nil
}

func (r *ScientificWork) UpdateStudentScientificWorks(ctx context.Context, tx *pgxpool.Pool, work *model.ScientificWork) error {
	if err := r.updateStudentScientificWorkTx(ctx, tx, work); err != nil {
		return errors.Wrap(err, "UpdateStudentScientificWorks()")
	}

	return nil
}

func (r *ScientificWork) updateStudentScientificWorkTx(ctx context.Context, tx *pgxpool.Pool, work *model.ScientificWork) error {
	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		stmt, args := table.ScientificWork.
			UPDATE(table.ScientificWork.MutableColumns).
			MODEL(work).
			WHERE(table.ScientificWork.WorkID.EQ(postgres.UUID(work.WorkID))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "updateStudentScientificWorkTx()")
	}

	return nil
}

func (r *ScientificWork) DeleteStudentScientificWorks(ctx context.Context, tx *pgxpool.Pool, workIDs []*uuid.UUID) error {
	var ids []postgres.Expression

	for _, workID := range workIDs {
		id := postgres.UUID(workID)
		ids = append(ids, id)
	}

	if err := r.deleteStudentScientificWorks(ctx, tx, ids); err != nil {
		return errors.Wrap(err, "DeleteStudentScientificWorks()")
	}

	return nil
}

func (r *ScientificWork) deleteStudentScientificWorks(ctx context.Context, tx *pgxpool.Pool, workIDs []postgres.Expression) error {
	stmt, args := table.ScientificWork.
		DELETE().
		WHERE(table.ScientificWork.WorkID.IN(workIDs...)).
		Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "deleteStudentScientificWorks()")
	}

	return nil
}

func scanScientificWork(row pgx.Row, target *model.ScientificWork) error {
	return row.Scan(
		&target.WorkID,
		&target.StudentID,
		&target.Semester,
		&target.Name,
		&target.State,
		&target.Impact,
		&target.OutputData,
		&target.CoAuthors,
		&target.WorkType,
	)
}
