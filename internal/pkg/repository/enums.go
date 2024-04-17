package repository

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type EnumRepository struct{}

func NewEnumRepository() *EnumRepository {
	return &EnumRepository{}
}

func (r *EnumRepository) InsertSpecializationsTx(ctx context.Context, tx pgx.Tx, specs []model.Specializations) error {
	stmt, args := table.Specializations.
		INSERT(table.Specializations.Title).
		MODELS(specs).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertSpecializationsTx()")
	}

	return nil
}

func (r *EnumRepository) GetSpecializationsTx(ctx context.Context, tx pgx.Tx) ([]model.Specializations, error) {
	stmt, args := table.Specializations.
		SELECT(table.Specializations.AllColumns).
		WHERE(table.Specializations.Archived.EQ(postgres.Bool(false))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetSpecializationsTx()")
	}
	defer rows.Close()

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

func (r *EnumRepository) UpdateSpecializationTx(ctx context.Context, tx pgx.Tx, spec model.Specializations) error {
	stmt, args := table.Specializations.
		UPDATE(table.Specializations.Title).
		SET(spec.Title).
		WHERE(table.Specializations.SpecID.EQ(postgres.Int32(spec.SpecID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpdateSpecializationTx()")
	}

	return nil
}

func (r *EnumRepository) ArchiveSpecializations(ctx context.Context, tx pgx.Tx, specsIDs []int32) error {
	expressions := make([]postgres.Expression, 0)

	for _, id := range specsIDs {
		exp := postgres.Expression(postgres.Int32(id))

		expressions = append(expressions, exp)
	}

	stmt, args := table.Specializations.
		UPDATE(table.Specializations.Archived).
		SET(true).
		WHERE(table.Specializations.SpecID.IN(expressions...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "ArchiveSpecializations()")
	}

	return nil
}

// TODO добавить CRUD для других энамов

func (r *EnumRepository) InsertGroupsTx(ctx context.Context, tx pgx.Tx, groups []model.Groups) error {
	stmt, args := table.Groups.
		INSERT(table.Groups.GroupName).
		MODELS(groups).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertGroupsTx()")
	}

	return nil
}

func (r *EnumRepository) GetGroupsTx(ctx context.Context, tx pgx.Tx) ([]model.Groups, error) {
	stmt, args := table.Groups.
		SELECT(table.Groups.AllColumns).
		WHERE(table.Groups.Archived.EQ(postgres.Bool(false))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetGroupsTx()")
	}
	defer rows.Close()

	groups := make([]model.Groups, 0)

	for rows.Next() {
		group := model.Groups{}
		if err := scanGroup(rows, &group); err != nil {
			return nil, errors.Wrap(err, "GetGroupsTx(): scanning row")
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func (r *EnumRepository) UpdateGroupTx(ctx context.Context, tx pgx.Tx, group model.Groups) error {
	stmt, args := table.Groups.
		UPDATE(table.Groups.GroupName).
		SET(group.GroupName).
		WHERE(table.Groups.GroupID.EQ(postgres.Int32(group.GroupID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpdateGroupTx()")
	}

	return nil
}

func (r *EnumRepository) DeleteGroupsTx(ctx context.Context, tx pgx.Tx, groupsIDs []int32) error {
	expressions := make([]postgres.Expression, 0)

	for _, id := range groupsIDs {
		exp := postgres.Expression(postgres.Int32(id))

		expressions = append(expressions, exp)
	}

	stmt, args := table.Groups.
		UPDATE(table.Groups.Archived).
		SET(true).
		WHERE(table.Groups.GroupID.IN(expressions...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteGroupsTx()")
	}

	return nil
}

func scanSpecialization(row pgx.Row, target *model.Specializations) error {
	return row.Scan(
		&target.SpecID,
		&target.Title,
	)
}

func scanGroup(row pgx.Row, target *model.Groups) error {
	return row.Scan(
		&target.GroupID,
		&target.GroupName,
	)
}
