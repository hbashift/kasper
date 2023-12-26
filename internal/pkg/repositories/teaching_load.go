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

func (r *TeachingLoadRepository) UpsertStudentsTeachingLoad(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error {
	if err := r.upsertStudentsTeachingLoadTx(ctx, tx, loads); err != nil {
		return errors.Wrap(err, "UpsertStudentsTeachingLoad()")
	}

	return nil
}

func (r *TeachingLoadRepository) upsertStudentsTeachingLoadTx(ctx context.Context, tx *pgxpool.Pool, loads []*model.TeachingLoad) error {
	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		for _, load := range loads {
			settingParams := []postgres.ColumnAssigment{
				table.TeachingLoad.LoadType.SET(table.TeachingLoad.EXCLUDED.LoadType),
				table.TeachingLoad.Hours.SET(postgres.Int32(load.Hours)),
				table.TeachingLoad.SubjectName.SET(postgres.String(load.SubjectName)),
				table.TeachingLoad.GroupName.SET(postgres.String(load.GroupName)),
				table.TeachingLoad.MainTeacher.SET(postgres.String(load.MainTeacher)),
			}

			if load.AdditionalLoad != nil {
				settingParams = append(settingParams,
					table.TeachingLoad.AdditionalLoad.SET(postgres.String(*load.AdditionalLoad)))
			}

			stmt, args := table.TeachingLoad.
				INSERT(table.TeachingLoad.AllColumns).
				MODEL(load).
				ON_CONFLICT(table.TeachingLoad.LoadID).
				DO_UPDATE(postgres.SET(settingParams...)).
				Sql()

			if _, err := tx.Exec(ctx, stmt, args...); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "upsertStudentsTeachingLoadTx()")
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

func (r *TeachingLoadRepository) DeleteTeachingLoad(ctx context.Context, tx *pgxpool.Pool, loadIDs []*uuid.UUID) error {
	var ids []postgres.Expression

	for _, loadID := range loadIDs {
		id := postgres.UUID(loadID)
		ids = append(ids, id)
	}

	if err := r.deleteTeachingLoadTx(ctx, tx, ids); err != nil {
		return errors.Wrap(err, "DeleteTeachingLoad()")
	}

	return nil
}

func (r *TeachingLoadRepository) deleteTeachingLoadTx(ctx context.Context, tx *pgxpool.Pool, loadIDs []postgres.Expression) error {
	stmt, args := table.TeachingLoad.
		DELETE().
		WHERE(table.TeachingLoad.LoadID.IN(loadIDs...)).
		Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "deleteTeachingLoadTx()")
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
