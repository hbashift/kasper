package repository

import (
	"context"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type DissertationRepository struct{}

func NewDissertationRepository() *DissertationRepository {
	return &DissertationRepository{}
}

func (r *DissertationRepository) SetSemesterProgressStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, acceptedAt *time.Time) error {
	stmt, args := table.SemesterProgress.
		UPDATE(
			table.SemesterProgress.Status,
			table.SemesterProgress.UpdatedAt,
			table.SemesterProgress.AcceptedAt,
		).
		SET(
			status,
			time.Now(),
			acceptedAt,
		).
		WHERE(table.SemesterProgress.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "SetSemesterProgressStatusTx()")
	}

	return nil
}

func (r *DissertationRepository) SetDissertationStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32) error {
	stmt, args := table.Dissertations.
		UPDATE(table.Dissertations.Status).
		SET(status).
		WHERE(table.Dissertations.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.Dissertations.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "SetDissertationStatusTx()")
	}

	return nil
}

func (r *DissertationRepository) SetDissertationTitleStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32, acceptedAt *time.Time) error {
	stmt, args := table.DissertationTitles.
		UPDATE(
			table.DissertationTitles.Status,
			table.DissertationTitles.AcceptedAt,
		).
		SET(
			status,
			acceptedAt,
		).
		WHERE(table.DissertationTitles.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.DissertationTitles.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "SetDissertationTitleStatusTx()")
	}

	return nil
}

func (r *DissertationRepository) GetSemesterProgressTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.SemesterProgress, error) {
	stmt, args := table.SemesterProgress.
		SELECT(table.SemesterProgress.AllColumns).
		WHERE(table.SemesterProgress.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetSemesterProgressTx()")
	}
	defer rows.Close()

	semesterProgress := make([]model.SemesterProgress, 0, 8)

	for rows.Next() {
		var progress = model.SemesterProgress{}
		if err := scanSemesterProgress(rows, &progress); err != nil {
			return nil, errors.Wrap(err, "GetSemesterProgressTx(): scanning semester progress rows")
		}
		semesterProgress = append(semesterProgress, progress)
	}

	return semesterProgress, nil
}

func (r *DissertationRepository) UpsertSemesterProgressTx(ctx context.Context, tx pgx.Tx, progresses []model.SemesterProgress) error {
	for _, semester := range progresses {
		assignments := []postgres.ColumnAssigment{
			table.SemesterProgress.First.SET(postgres.Bool(semester.First)),
			table.SemesterProgress.Second.SET(postgres.Bool(semester.Second)),
			table.SemesterProgress.Third.SET(postgres.Bool(semester.Third)),
			table.SemesterProgress.Forth.SET(postgres.Bool(semester.Forth)),
			table.SemesterProgress.Fifth.SET(postgres.Bool(semester.Fifth)),
			table.SemesterProgress.Sixth.SET(postgres.Bool(semester.Sixth)),
			table.SemesterProgress.Seventh.SET(postgres.Bool(semester.Seventh)),
			table.SemesterProgress.Eighth.SET(postgres.Bool(semester.Eighth)),
			table.SemesterProgress.UpdatedAt.SET(postgres.NOW()),
			table.SemesterProgress.Status.SET(postgres.NewEnumValue(semester.Status.String())),
		}

		if semester.Status == model.ApprovalStatus_Approved {
			assignments = append(assignments,
				table.SemesterProgress.AcceptedAt.SET(postgres.TimestampzT(lo.FromPtr(semester.AcceptedAt))))
		} else {
			assignments = append(assignments,
				table.SemesterProgress.AcceptedAt.SET(postgres.RawTimestampz(`null`)))
		}

		stmt, args := table.SemesterProgress.
			INSERT().
			MODEL(semester).
			ON_CONFLICT(table.SemesterProgress.StudentID, table.SemesterProgress.ProgressType).
			DO_UPDATE(postgres.
				SET(assignments...),
			).
			Sql()

		_, err := tx.Exec(ctx, stmt, args...)
		if err != nil {
			return errors.Wrap(err, "UpsertSemesterProgressTx()")
		}
	}

	return nil
}

func (r *DissertationRepository) UpsertDissertationTx(ctx context.Context, tx pgx.Tx, model model.Dissertations) error {
	stmt, args := table.Dissertations.
		INSERT().
		MODEL(model).
		ON_CONFLICT(table.Dissertations.StudentID, table.Dissertations.Semester).
		DO_UPDATE(postgres.
			SET(
				table.Dissertations.Status.SET(postgres.NewEnumValue(model.Status.String())),
				table.Dissertations.UpdatedAt.SET(postgres.NOW()),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "UpsertDissertationTx()")
	}

	return nil
}

func (r *DissertationRepository) GetDissertationsTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Dissertations, error) {
	stmt, args := table.Dissertations.
		SELECT(table.Dissertations.AllColumns).
		WHERE(table.Dissertations.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetDissertationsTx()")
	}
	defer rows.Close()

	dissertations := make([]model.Dissertations, 0, 8)

	for rows.Next() {
		dissertation := model.Dissertations{}
		if err := scanDissertation(rows, &dissertation); err != nil {
			return nil, errors.Wrap(err, "GetDissertationsTx(): scanning dissertations rows")
		}
		dissertations = append(dissertations, dissertation)
	}

	return dissertations, nil
}

func (r *DissertationRepository) GetDissertationDataBySemester(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, semester int32) (model.Dissertations, error) {
	stmt, args := table.Dissertations.
		SELECT(table.Dissertations.AllColumns).
		WHERE(table.Dissertations.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.Dissertations.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)
	dissertation := model.Dissertations{}

	if err := scanDissertation(row, &dissertation); err != nil {
		return model.Dissertations{}, errors.Wrap(err, "GetDissertationDataBySemester")
	}

	return dissertation, nil
}

func (r *DissertationRepository) InsertDissertationTitleTx(ctx context.Context, tx pgx.Tx, title model.DissertationTitles) error {
	stmt, args := table.DissertationTitles.
		INSERT().
		MODEL(title).
		Sql()

	deleteStmt, args := table.DissertationTitles.
		DELETE().
		WHERE(table.DissertationTitles.Status.NOT_EQ(postgres.NewEnumValue(model.ApprovalStatus_Approved.String())).
			AND(table.DissertationTitles.StudentID.EQ(postgres.UUID(title.StudentID)))).
		Sql()

	_, err := tx.Exec(ctx, deleteStmt, args...)
	if err != nil {
		return errors.Wrap(err, "InsertDissertationTitleTx(): delete")
	}

	_, err = tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "InsertDissertationTitleTx(): insert")
	}

	return nil
}

func (r *DissertationRepository) GetDissertationTitlesTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.DissertationTitles, error) {
	stmt, args := table.DissertationTitles.
		SELECT(table.DissertationTitles.AllColumns).
		WHERE(table.DissertationTitles.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetDissertationTitlesTx()")
	}
	defer rows.Close()

	titles := make([]model.DissertationTitles, 0)

	for rows.Next() {
		title := model.DissertationTitles{}

		if err := scanDissertationTitle(rows, &title); err != nil {
			return nil, errors.Wrap(err, "GetDissertationTitlesTx()")
		}

		titles = append(titles, title)
	}

	return titles, nil
}

func (r *DissertationRepository) GetFeedbackTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Feedback, error) {
	stmt, args := table.Feedback.
		SELECT(table.Feedback.AllColumns).
		WHERE(table.Feedback.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetFeedbackTx()")
	}
	defer rows.Close()

	feedbacks := make([]model.Feedback, 0, 8)

	for rows.Next() {
		feedback := model.Feedback{}
		if err := scanFeedback(rows, &feedback); err != nil {
			return nil, errors.Wrap(err, "GetFeedbackTx(): scanning feedback row")
		}

		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}

func (r *DissertationRepository) UpsertFeedbackTx(ctx context.Context, tx pgx.Tx, feedback model.Feedback) error {
	stmt, args := table.Feedback.
		INSERT(
			table.Feedback.AllColumns.Except(table.Feedback.CreatedAt, table.Feedback.UpdatedAt),
		).
		MODEL(feedback).
		ON_CONFLICT(table.Feedback.StudentID, table.Feedback.Semester).
		DO_UPDATE(postgres.
			SET(table.Feedback.Feedback.SET(postgres.String(*feedback.Feedback)))).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "UpsertFeedbackTx()")
	}

	return nil
}

func (r *DissertationRepository) GetStudentsProgressiveness(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Progressiveness, error) {
	stmt, args := table.Progressiveness.
		SELECT(table.Progressiveness.AllColumns).
		WHERE(table.Progressiveness.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentsProgressiveness()")
	}

	progresses := make([]model.Progressiveness, 0)
	for rows.Next() {
		progress := model.Progressiveness{}
		if err := scanProgressiveness(rows, &progress); err != nil {
			return nil, errors.Wrap(err, "GetStudentsProgressiveness(): scanning row")
		}

		progresses = append(progresses, progress)
	}

	return progresses, nil
}

func (r *DissertationRepository) UpsertStudentsProgressiveness(ctx context.Context, tx pgx.Tx, progress model.Progressiveness) error {
	stmt, args := table.Progressiveness.
		INSERT().
		MODEL(progress).
		ON_CONFLICT(table.Progressiveness.StudentID, table.Progressiveness.Semester).
		DO_UPDATE(postgres.
			SET(table.Progressiveness.Progressiveness.SET(postgres.Int32(progress.Progressiveness))),
		).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "UpsertStudentsProgressiveness()")
	}

	return nil
}

func scanProgressiveness(row pgx.Row, target *model.Progressiveness) error {
	return row.Scan(
		&target.ProgressID,
		&target.StudentID,
		&target.Semester,
		&target.Progressiveness,
	)
}

func scanSemesterProgress(row pgx.Row, target *model.SemesterProgress) error {
	return row.Scan(
		&target.ProgressID,
		&target.StudentID,
		&target.ProgressType,
		&target.First,
		&target.Second,
		&target.Third,
		&target.Forth,
		&target.Fifth,
		&target.Sixth,
		&target.Seventh,
		&target.Eighth,
		&target.UpdatedAt,
		&target.Status,
		&target.AcceptedAt,
	)
}

func scanDissertation(row pgx.Row, target *model.Dissertations) error {
	return row.Scan(
		&target.DissertationID,
		&target.StudentID,
		&target.Status,
		&target.CreatedAt,
		&target.UpdatedAt,
		&target.Semester,
		&target.FileName,
	)
}

func scanDissertationTitle(row pgx.Row, target *model.DissertationTitles) error {
	return row.Scan(
		&target.TitleID,
		&target.StudentID,
		&target.Title,
		&target.CreatedAt,
		&target.Status,
		&target.AcceptedAt,
		&target.Semester,
		&target.ResearchObject,
		&target.ResearchSubject,
	)
}

func scanFeedback(row pgx.Row, target *model.Feedback) error {
	return row.Scan(
		&target.FeedbackID,
		&target.StudentID,
		&target.DissertationID,
		&target.Feedback,
		&target.Semester,
		&target.CreatedAt,
		&target.UpdatedAt,
	)
}
