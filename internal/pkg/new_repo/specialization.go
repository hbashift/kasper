package new_repo

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type EnumRepository struct{}

func NewSpecializationRepository() *EnumRepository {
	return &EnumRepository{}
}

func (r *EnumRepository) GetSpecializationsTx(ctx context.Context, tx *pgxpool.Tx) ([]model.Specializations, error) {
	stmt, args := table.Specializations.
		SELECT(table.Specializations.AllColumns).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetSpecializationsTx()")
	}

	specs := make([]model.Specializations, 0)

	for rows.Next() {
		spec := model.Specializations{}

		if err := scanSpecialization(rows, &spec); err != nil {
			return nil, errors.Wrap(err, "GetSpecializationsTx(): scanning row")
		}

		specs = append(specs, spec)
	}

	return specs, nil
}

func (r *EnumRepository) InsertSpecializationsTx(ctx context.Context, tx *pgxpool.Tx, specs []model.Specializations) error {
	stmt, args := table.Specializations.
		INSERT().
		MODELS(specs).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertSpecializationsTx()")
	}

	return nil
}

// TODO добавить CRUD для других энамов

func scanSpecialization(row pgx.Row, target *model.Specializations) error {
	return row.Scan(
		&target.SpecID,
		&target.Title,
	)
}
