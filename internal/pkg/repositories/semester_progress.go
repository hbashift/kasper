package repositories

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
	"uir_draft/internal/pkg/models"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type SemesterRepository struct {
	postgres *pgxpool.Pool
}

func NewSemesterRepository(postgres *pgxpool.Pool) *SemesterRepository {
	return &SemesterRepository{postgres: postgres}
}

func (r *SemesterRepository) GetStudentDissertationPlan(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) ([]*models.StudentDissertationPlan, error) {
	plan, err := r.getStudentDissertationPlanTx(ctx, tx, clientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentDissertationPlan():")
	}

	return plan, nil
}

func (r *SemesterRepository) getStudentDissertationPlanTx(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) ([]*models.StudentDissertationPlan, error) {
	stmt, args := table.SemesterProgress.
		SELECT(
			table.SemesterProgress.ProgressName.AS("name"),
			table.SemesterProgress.First,
			table.SemesterProgress.Second,
			table.SemesterProgress.Third,
			table.SemesterProgress.Forth,
			table.SemesterProgress.Fifth,
			table.SemesterProgress.Sixth,
		).
		WHERE(table.SemesterProgress.ClientID.EQ(postgres.UUID(clientID))).Sql()

	studentPlan := make([]*models.StudentDissertationPlan, 0)

	rows, err := tx.Query(ctx, stmt, args...)
	defer rows.Close()

	if err != nil {
		return nil, errors.Wrap(err, "selecting students dissertation plan")
	}

	for rows.Next() {
		var plan = &models.StudentDissertationPlan{}
		if err := scanDissertationPlan(rows, plan); err != nil {
			return nil, errors.Wrap(err, "mapping dissertation plan rows")
		}
		studentPlan = append(studentPlan, plan)
	}

	return studentPlan, nil
}

func (r *SemesterRepository) UpsertSemesterPlan(ctx context.Context, tx *pgxpool.Pool, progress []*model.SemesterProgress) error {
	if err := r.upsertSemesterPlanTx(ctx, tx, progress); err != nil {
		return errors.Wrap(err, "UpsertSemesterPlan(): error during transaction")
	}

	return nil
}

func (r *SemesterRepository) upsertSemesterPlanTx(ctx context.Context, tx *pgxpool.Pool, progress []*model.SemesterProgress) error {
	if err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		for _, semester := range progress {
			stmt, args := table.SemesterProgress.
				INSERT(table.SemesterProgress.MutableColumns).
				MODEL(semester).
				ON_CONFLICT(table.SemesterProgress.StudentID, table.SemesterProgress.ProgressName).
				DO_UPDATE(postgres.
					SET(
						table.SemesterProgress.First.SET(postgres.Bool(semester.First)),
						table.SemesterProgress.Second.SET(postgres.Bool(semester.Second)),
						table.SemesterProgress.Third.SET(postgres.Bool(semester.Third)),
						table.SemesterProgress.Forth.SET(postgres.Bool(semester.Forth)),
						table.SemesterProgress.Fifth.SET(postgres.Bool(*semester.Fifth)),
						table.SemesterProgress.Sixth.SET(postgres.Bool(*semester.Sixth)),
						table.SemesterProgress.LastUpdated.SET(postgres.TimestampzT(*semester.LastUpdated)),
					),
				).
				Sql()

			_, err := tx.Exec(ctx, stmt, args...)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "upsert semester progress")
	}

	return nil
}

func scanDissertationPlan(row pgx.Row, target *models.StudentDissertationPlan) error {
	return row.Scan(
		&target.Name,
		&target.First,
		&target.Second,
		&target.Third,
		&target.Forth,
		&target.Fifth,
		&target.Sixth,
	)
}
