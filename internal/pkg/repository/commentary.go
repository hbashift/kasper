package repository

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type CommentaryRepository struct{}

func NewCommentaryRepository() *CommentaryRepository {
	return &CommentaryRepository{}
}

func (r *CommentaryRepository) UpsertStudentsComment(
	ctx context.Context,
	tx pgx.Tx,
	comment model.StudentsCommentary,
) error {
	sql := table.StudentsCommentary.
		INSERT().
		MODEL(comment)

	if comment.Commentary != nil {
		sql = sql.
			ON_CONFLICT(table.StudentsCommentary.StudentID, table.StudentsCommentary.Semester).
			DO_UPDATE(postgres.
				SET(
					table.StudentsCommentary.Commentary.
						SET(postgres.String(lo.FromPtr(comment.Commentary))),
				),
			)
	} else {
		sql = sql.
			ON_CONFLICT(table.StudentsCommentary.StudentID, table.StudentsCommentary.Semester).
			DO_NOTHING()
	}

	stmt, args := sql.Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpsertStudentsComment()")
	}

	return nil
}

func (r *CommentaryRepository) GetStudentsCommentaries(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.StudentsCommentary, error) {
	stmt, args := table.StudentsCommentary.
		SELECT(table.StudentsCommentary.AllColumns).
		WHERE(table.StudentsCommentary.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentsCommentaries()")
	}

	comments := make([]model.StudentsCommentary, 0)
	for rows.Next() {
		comment := model.StudentsCommentary{}

		if err := scanStudentsCommentary(rows, &comment); err != nil {
			return nil, errors.Wrap(err, "GetStudentsCommentaries(): scanning row")
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *CommentaryRepository) UpsertDissertationComment(ctx context.Context, tx pgx.Tx, comment model.DissertationCommentary) error {
	sql := table.DissertationCommentary.
		INSERT().
		MODEL(comment)

	if comment.Commentary != nil {
		sql = sql.
			ON_CONFLICT(table.DissertationCommentary.StudentID, table.DissertationCommentary.Semester).
			DO_UPDATE(postgres.
				SET(
					table.DissertationCommentary.Commentary.
						SET(postgres.String(lo.FromPtr(comment.Commentary))),
				),
			)
	} else {
		sql = sql.
			ON_CONFLICT(table.DissertationCommentary.StudentID, table.DissertationCommentary.Semester).
			DO_NOTHING()
	}

	stmt, args := sql.Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpsertDissertationComment()")
	}

	return nil
}

func (r *CommentaryRepository) GetDissertationComments(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.DissertationCommentary, error) {
	stmt, args := table.DissertationCommentary.
		SELECT(table.DissertationCommentary.AllColumns).
		WHERE(table.DissertationCommentary.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetDissertationComments()")
	}

	comments := make([]model.DissertationCommentary, 0)
	for rows.Next() {
		comment := model.DissertationCommentary{}

		if err := scanDissertationCommentary(rows, &comment); err != nil {
			return nil, errors.Wrap(err, "GetDissertationComments(): scanning row")
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *CommentaryRepository) UpsertPlanComment(ctx context.Context, tx pgx.Tx, plan model.DissertationPlans) error {
	sql := table.DissertationPlans.
		INSERT().
		MODEL(plan)

	if plan.PlanText != nil {
		sql = sql.
			ON_CONFLICT(table.DissertationPlans.StudentID, table.DissertationPlans.Semester).
			DO_UPDATE(postgres.
				SET(
					table.DissertationPlans.PlanText.
						SET(postgres.String(lo.FromPtr(plan.PlanText))),
				),
			)
	} else {
		sql = sql.
			ON_CONFLICT(table.DissertationPlans.StudentID, table.DissertationPlans.Semester).
			DO_NOTHING()
	}

	stmt, args := sql.Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpsertPlanComment()")
	}

	return nil
}

func (r *CommentaryRepository) GetPlanComments(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.DissertationPlans, error) {
	stmt, args := table.DissertationPlans.
		SELECT(table.DissertationPlans.AllColumns).
		WHERE(table.DissertationPlans.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetPlanComments()")
	}

	comments := make([]model.DissertationPlans, 0)
	for rows.Next() {
		comment := model.DissertationPlans{}

		if err := scanDissertationPlan(rows, &comment); err != nil {
			return nil, errors.Wrap(err, "GetPlanComments(): scanning row")
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func scanStudentsCommentary(row pgx.Row, target *model.StudentsCommentary) error {
	return row.Scan(
		&target.CommentaryID,
		&target.StudentID,
		&target.Semester,
		&target.Commentary,
		&target.CommentedAt,
	)
}

func scanDissertationCommentary(row pgx.Row, target *model.DissertationCommentary) error {
	return row.Scan(
		&target.CommentaryID,
		&target.StudentID,
		&target.Semester,
		&target.Commentary,
	)
}

func scanDissertationPlan(row pgx.Row, target *model.DissertationPlans) error {
	return row.Scan(
		&target.PlanID,
		&target.StudentID,
		&target.Semester,
		&target.PlanText,
	)
}
