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

type DissertationRepository struct {
	postgres *pgxpool.Pool
}

func NewDissertationRepository(postgres *pgxpool.Pool) *DissertationRepository {
	return &DissertationRepository{postgres: postgres}
}

func (r *DissertationRepository) updateDissertationStatusTx(ctx context.Context, tx pgx.Tx, dissertationID uuid.UUID, status string) error {
	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		stmt, args := table.Dissertation.UPDATE(table.Dissertation.Status).
			SET(status).
			WHERE(table.Dissertation.DissertationID.EQ(postgres.UUID(dissertationID))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "insert semester progress")
	}

	return nil
}

func (r *DissertationRepository) insertDissertationTx(ctx context.Context, tx pgx.Tx, dissertation model.Dissertation) error {
	stmt, args := table.Dissertation.INSERT(table.Dissertation.AllColumns).
		MODEL(dissertation).Sql()

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "insert dissertation")
	}

	return nil
}
