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

type TeachingLoadRepository struct {
}

func NewTeachingLoadRepository() *TeachingLoadRepository {
	return &TeachingLoadRepository{}
}

func (r *TeachingLoadRepository) GetStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.TeachingLoad, error) {
	load, err := r.getStudentsTeachingLoadTx(ctx, tx, studentID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentsTeachingLoad()")
	}

	return load, nil
}

func (r *TeachingLoadRepository) getStudentsTeachingLoadTx(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*model.TeachingLoad, error) {
	stmt, args := table.TeachingLoad.
		SELECT(table.TeachingLoad.AllColumns).
		WHERE(table.TeachingLoad.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	var loads []*model.TeachingLoad

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, stmt, args...)
		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			load := &model.TeachingLoad{}
			err := scanTeachingLoad(rows, load)
			if err != nil {
				return err
			}

			loads = append(loads, load)
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "getStudentsTeachingLoadTx()")
	}

	return loads, nil
}

func (r *TeachingLoadRepository) InsertStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error {
	if err := r.insertStudentsTeachingLoadTx(ctx, tx, loads); err != nil {
		return errors.Wrap(err, "InsertStudentsTeachingLoad()")
	}

	return nil
}

func (r *TeachingLoadRepository) insertStudentsTeachingLoadTx(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error {
	stmt, args := table.TeachingLoad.
		INSERT(table.TeachingLoad.AllColumns).
		MODELS(loads).
		ON_CONFLICT().
		DO_NOTHING().
		Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "insertStudentsTeachingLoadTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) UpdateStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error {
	if err := r.updateStudentsTeachingLoadTx(ctx, tx, loads); err != nil {
		return errors.Wrap(err, "UpdateStudentsTeachingLoad()")
	}

	return nil
}

func (r *TeachingLoadRepository) updateStudentsTeachingLoadTx(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error {
	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		for _, load := range loads {
			stmt, args := table.TeachingLoad.
				UPDATE(table.TeachingLoad.MutableColumns).
				MODEL(load).
				WHERE(table.TeachingLoad.LoadID.EQ(postgres.UUID(load.LoadID))).
				Sql()

			if _, err := tx.Exec(ctx, stmt, args...); err != nil {
				return err
			}
		}

		return nil

	}); err != nil {
		return errors.Wrap(err, "updateStudentsTeachingLoadTx()")
	}

	return nil
}

func scanTeachingLoad(row pgx.Row, target *model.TeachingLoad) error {
	return row.Scan(
		&target.LoadID,
		&target.StudentID,
		&target.Semester,
		&target.Hours,
		&target.AdditionalLoad,
		&target.LoadType,
		&target.MainTeacher,
		&target.GroupName,
		&target.SubjectName,
	)
}