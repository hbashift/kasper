package repositories

import (
	"context"
	"time"

	"github.com/samber/lo"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
	"uir_draft/internal/pkg/models"
	"uir_draft/internal/pkg/service/student/mapping"

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

func (r *DissertationRepository) GetDissertationIDs(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*models.IDs, error) {
	return r.getDissertationIDsTx(ctx, tx, studentID)
}

func (r *DissertationRepository) getDissertationIDsTx(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*models.IDs, error) {
	stmt, args := table.Dissertation.
		SELECT(
			table.Dissertation.DissertationID.AS("id"),
			table.Dissertation.Semester,
		).
		WHERE(table.Dissertation.StudentID.EQ(postgres.UUID(studentID))).Sql()

	var ids []*models.IDs

	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, stmt, args...)
		if err != nil {
			return err
		}

		for rows.Next() {
			id := &models.IDs{}
			if err := scanDissertationIDs(rows, id); err != nil {
				return err
			}

			ids = append(ids, id)
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "getDissertationIDsTx()")
	}

	return ids, nil
}

func (r *DissertationRepository) UpsertDissertationData(ctx context.Context, tx *pgxpool.Pool, studentID *uuid.UUID, semester int32, name string) error {
	return r.upsertDissertationDataTx(ctx, tx, studentID, semester, name)
}

func (r *DissertationRepository) upsertDissertationDataTx(ctx context.Context, tx *pgxpool.Pool, studentID *uuid.UUID, semester int32, name string) error {
	dissertation := model.Dissertation{
		StudentID:      *studentID,
		CreatedAt:      lo.ToPtr(time.Now()),
		UpdatedAt:      lo.ToPtr(time.Now()),
		DissertationID: uuid.New(),
		Semester:       semester,
		Name:           name,
		Status:         model.DissertationStatus_Empty,
	}

	stmt, args := table.Dissertation.
		INSERT(table.Dissertation.AllColumns).
		MODEL(dissertation).
		ON_CONFLICT(table.Dissertation.StudentID, table.Dissertation.Semester).
		DO_UPDATE(
			postgres.SET(
				table.Dissertation.Name.SET(postgres.String(name)),
				table.Dissertation.UpdatedAt.SET(postgres.NOW()),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "upsertDissertationDataTx()")
	}

	return nil
}

func (r *DissertationRepository) GetDissertationData(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, semester int32) (*model.Dissertation, error) {
	return r.getDissertationDataTx(ctx, tx, studentID, semester)
}

func (r *DissertationRepository) getDissertationDataTx(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, semester int32) (*model.Dissertation, error) {
	stmt, args := table.Dissertation.
		SELECT(table.Dissertation.AllColumns).
		WHERE(table.Dissertation.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.Dissertation.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	dissertation := &model.Dissertation{}
	row := tx.QueryRow(ctx, stmt, args...)
	if err := scanDissertation(row, dissertation); err != nil {
		return nil, errors.Wrap(err, "getDissertationDataTx()")
	}

	return dissertation, nil
}

func (r *DissertationRepository) GetStatuses(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID) ([]*mapping.DissertationStatus, error) {
	stmt, args := table.Dissertation.
		SELECT(table.Dissertation.Status, table.Dissertation.Semester).
		WHERE(table.Dissertation.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetStatuses()")
	}

	var statuses []*mapping.DissertationStatus

	for rows.Next() {
		status := &mapping.DissertationStatus{}
		if err := scanDissertationStatus(rows, status); err != nil {
			return nil, errors.Wrap(err, "scanDissertationStatus()")
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}

func (r *DissertationRepository) SetStatus(ctx context.Context, tx *pgxpool.Pool, studentID uuid.UUID, semester int32, status model.DissertationStatus) error {
	stmt, args := table.Dissertation.
		UPDATE(table.Dissertation.Status).
		SET(status).
		WHERE(table.Dissertation.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.Dissertation.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "SetStatus()")
	}

	return nil
}

func scanDissertationStatus(rows pgx.Row, target *mapping.DissertationStatus) error {
	return rows.Scan(
		&target.Status,
		&target.Semester,
	)
}

func scanDissertationIDs(rows pgx.Row, target *models.IDs) error {
	return rows.Scan(
		&target.ID,
		&target.Semester,
	)
}

func scanDissertation(row pgx.Row, target *model.Dissertation) error {
	return row.Scan(
		&target.StudentID,
		&target.Status,
		&target.CreatedAt,
		&target.UpdatedAt,
		&target.DissertationID,
		&target.Semester,
		&target.Name,
	)
}
