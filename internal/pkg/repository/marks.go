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

type MarksRepository struct{}

func NewMarksRepository() *MarksRepository {
	return &MarksRepository{}
}

func (r *MarksRepository) GetStudentsAttestationMarksTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Marks, error) {
	stmt, args := table.Marks.
		SELECT(table.Marks.AllColumns).
		WHERE(table.Marks.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentsAttestationMarksTx()")
	}
	defer rows.Close()

	marks := make([]model.Marks, 0, 8)

	for rows.Next() {
		mark := model.Marks{}

		if err := scanMarks(rows, &mark); err != nil {
			return nil, errors.Wrap(err, "GetStudentsAttestationMarksTx(): scanning row")
		}

		marks = append(marks, mark)
	}

	return marks, nil
}

func (r *MarksRepository) UpsertAttestationMarksTx(ctx context.Context, tx pgx.Tx, models []model.Marks) error {
	for _, mark := range models {
		stmt, args := table.Marks.
			INSERT().
			MODEL(mark).
			ON_CONFLICT(table.Marks.StudentID, table.Marks.Semester).
			DO_UPDATE(postgres.
				SET(table.Marks.Mark.SET(postgres.Int32(mark.Mark)))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpsertAttestationMarksTx()")
		}
	}
	return nil
}

//func (r *MarksRepository) InsertAttestationMarks(ctx context.Context, tx pgx.Tx, models model.Marks) error {
//	stmt, args := table.Marks.
//		INSERT().
//		MODEL(models).
//		Sql()
//
//	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
//		return errors.Wrap(err, "InsertAttestationMarks()")
//	}
//
//	return nil
//}

func (r *MarksRepository) GetStudentsExamResults(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.Exams, error) {
	stmt, args := table.Exams.
		SELECT(table.Exams.AllColumns).
		WHERE(table.Exams.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentsExamResults()")
	}

	exams := make([]model.Exams, 0)
	for rows.Next() {
		exam := model.Exams{}
		if err := scanExams(rows, &exam); err != nil {
			return nil, errors.Wrap(err, "GetStudentsExamResults(): scanning rows")
		}

		exams = append(exams, exam)
	}

	return exams, nil
}

func (r *MarksRepository) UpsertExamResults(ctx context.Context, tx pgx.Tx, models []model.Exams) error {
	for _, exam := range models {
		assignments := []postgres.ColumnAssigment{
			table.Exams.Mark.SET(postgres.Int32(exam.Mark)),
		}

		if exam.SetAt != nil {
			assignments = append(assignments, table.Exams.SetAt.SET(postgres.TimestampzT(lo.FromPtr(exam.SetAt))))
		} else {
			assignments = append(assignments, table.Exams.SetAt.SET(postgres.RawTimestampz("null")))
		}

		stmt, args := table.Exams.
			INSERT().
			MODEL(exam).
			ON_CONFLICT(table.Exams.StudentID, table.Exams.Semester, table.Exams.ExamType).
			DO_UPDATE(postgres.
				SET(assignments...)).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpsertExamResults()")
		}
	}

	return nil
}

func (r *MarksRepository) GetStudentsSupervisorMarks(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.SupervisorMarks, error) {
	stmt, args := table.SupervisorMarks.
		SELECT(table.SupervisorMarks.AllColumns).
		WHERE(table.SupervisorMarks.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentsSupervisorMarks()")
	}

	supervisorMarks := make([]model.SupervisorMarks, 0)
	for rows.Next() {
		exam := model.SupervisorMarks{}
		if err := scanSupervisorMarks(rows, &exam); err != nil {
			return nil, errors.Wrap(err, "GetStudentsSupervisorMarks(): scanning rows")
		}

		supervisorMarks = append(supervisorMarks, exam)
	}

	return supervisorMarks, nil
}

func (r *MarksRepository) UpsertStudentsSupervisorMark(ctx context.Context, tx pgx.Tx, model model.SupervisorMarks) error {
	stmt, args := table.SupervisorMarks.
		INSERT().
		MODEL(model).
		ON_CONFLICT(table.SupervisorMarks.StudentID, table.SupervisorMarks.Semester).
		DO_UPDATE(postgres.
			SET(table.SupervisorMarks.Mark.SET(postgres.Int32(model.Mark)))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpsertStudentsSupervisorMark()")
	}

	return nil
}

func scanMarks(row pgx.Row, target *model.Marks) error {
	return row.Scan(
		&target.StudentID,
		&target.Mark,
		&target.Semester,
	)
}

func scanExams(row pgx.Row, target *model.Exams) error {
	return row.Scan(
		&target.ExamID,
		&target.StudentID,
		&target.ExamType,
		&target.Semester,
		&target.Mark,
		&target.SetAt,
	)
}

func scanSupervisorMarks(row pgx.Row, target *model.SupervisorMarks) error {
	return row.Scan(
		&target.MarkID,
		&target.StudentID,
		&target.Mark,
		&target.Semester,
		&target.SupervisorID,
	)
}
