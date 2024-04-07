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
)

type TeachingLoadRepository struct{}

func NewTeachingLoadRepository() *TeachingLoadRepository {
	return &TeachingLoadRepository{}
}

func (r *TeachingLoadRepository) InitTeachingLoadsStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) error {
	loads := make([]model.TeachingLoadStatus, 0, 8)

	for i := 1; i < 9; i++ {
		load := model.TeachingLoadStatus{
			LoadsID:    uuid.New(),
			StudentID:  studentID,
			Semester:   int32(i),
			Status:     model.ApprovalStatus_Empty,
			UpdatedAt:  time.Now(),
			AcceptedAt: nil,
		}

		loads = append(loads, load)
	}

	stmt, args := table.TeachingLoadStatus.
		INSERT().
		MODELS(loads).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InitTeachingLoadsStatusTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) SetTeachingLoadStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32, acceptedAt *time.Time) error {
	stmt, args := table.TeachingLoadStatus.
		UPDATE(
			table.TeachingLoadStatus.UpdatedAt,
			table.TeachingLoadStatus.AcceptedAt,
			table.TeachingLoadStatus.Status,
		).
		SET(
			time.Now(),
			acceptedAt,
			status,
		).
		WHERE(table.TeachingLoadStatus.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.TeachingLoadStatus.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "SetTeachingLoadStatusTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) GetTeachingLoadStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.TeachingLoadStatus, error) {
	stmt, args := table.TeachingLoadStatus.
		SELECT(table.TeachingLoadStatus.AllColumns).
		WHERE(table.TeachingLoadStatus.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetTeachingLoadStatusTx()")
	}
	defer rows.Close()

	loads := make([]model.TeachingLoadStatus, 0, 8)

	for rows.Next() {
		load := model.TeachingLoadStatus{}

		if err := scanTeachingLoadStatusStatus(rows, &load); err != nil {
			return nil, errors.Wrap(err, "GetTeachingLoadStatusTx()")
		}

		loads = append(loads, load)
	}

	return loads, nil
}

func (r *TeachingLoadRepository) GetTeachingLoadStatusBySemesterTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, semester int32) (model.TeachingLoadStatus, error) {
	stmt, args := table.TeachingLoadStatus.
		SELECT(table.TeachingLoadStatus.AllColumns).
		WHERE(table.TeachingLoadStatus.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows := tx.QueryRow(ctx, stmt, args...)

	load := model.TeachingLoadStatus{}

	if err := scanTeachingLoadStatusStatus(rows, &load); err != nil {
		return model.TeachingLoadStatus{}, errors.Wrap(err, "GetTeachingLoadStatusTx()")
	}

	return load, nil
}

func (r *TeachingLoadRepository) UpdateTeachingLoadStatusTx(ctx context.Context, tx pgx.Tx, loads []model.TeachingLoadStatus) error {
	for _, load := range loads {
		stmt, args := table.TeachingLoadStatus.
			UPDATE(
				table.TeachingLoadStatus.Status,
				table.TeachingLoadStatus.UpdatedAt,
				table.TeachingLoadStatus.AcceptedAt,
			).
			SET(
				load.Status,
				load.UpdatedAt,
				load.AcceptedAt,
			).
			WHERE(table.TeachingLoadStatus.LoadsID.EQ(postgres.UUID(load.LoadsID)).
				AND(table.TeachingLoadStatus.Semester.EQ(postgres.Int32(load.Semester)))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpdateTeachingLoadStatusTx()")
		}
	}

	return nil
}

func (r *TeachingLoadRepository) InsertClassroomLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.ClassroomLoad) error {
	stmt, args := table.ClassroomLoad.
		INSERT().
		MODELS(loads).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertClassroomLoadTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) UpdateClassroomLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.ClassroomLoad) error {
	for _, load := range loads {
		stmt, args := table.ClassroomLoad.
			UPDATE(
				table.ClassroomLoad.Hours,
				table.ClassroomLoad.LoadType,
				table.ClassroomLoad.MainTeacher,
				table.ClassroomLoad.GroupName,
				table.ClassroomLoad.SubjectName,
			).
			SET(
				load.Hours,
				load.LoadType,
				load.MainTeacher,
				load.GroupName,
				load.SubjectName,
			).
			WHERE(table.ClassroomLoad.LoadID.EQ(postgres.UUID(load.LoadID))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpdateClassroomLoadTx()")
		}
	}

	return nil
}

func (r *TeachingLoadRepository) DeleteClassroomLoadsTx(ctx context.Context, tx pgx.Tx, classroomsIDs []uuid.UUID) error {
	var exps []postgres.Expression
	for _, id := range classroomsIDs {
		exp := postgres.Expression(postgres.UUID(id))

		exps = append(exps, exp)
	}

	stmt, args := table.ClassroomLoad.
		DELETE().
		WHERE(table.ClassroomLoad.LoadID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteClassroomLoadTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) InsertIndividualLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.IndividualStudentsLoad) error {
	stmt, args := table.IndividualStudentsLoad.
		INSERT().
		MODELS(loads).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertIndividualLoadTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) UpdateIndividualLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.IndividualStudentsLoad) error {
	for _, load := range loads {
		stmt, args := table.IndividualStudentsLoad.
			UPDATE(
				table.IndividualStudentsLoad.StudentsAmount,
				table.IndividualStudentsLoad.Comment,
			).
			SET(
				load.StudentsAmount,
				load.Comment,
			).
			WHERE(table.IndividualStudentsLoad.LoadID.EQ(postgres.UUID(load.LoadID))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpdateIndividualLoadTx()")
		}
	}

	return nil
}

func (r *TeachingLoadRepository) DeleteIndividualStudentsLoadsTx(ctx context.Context, tx pgx.Tx, individualsIDs []uuid.UUID) error {
	var exps []postgres.Expression
	for _, id := range individualsIDs {
		exp := postgres.Expression(postgres.UUID(id))

		exps = append(exps, exp)
	}

	stmt, args := table.IndividualStudentsLoad.
		DELETE().
		WHERE(table.IndividualStudentsLoad.LoadID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteIndividualStudentsLoadTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) InsertAdditionalLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.AdditionalLoad) error {
	stmt, args := table.AdditionalLoad.
		INSERT().
		MODELS(loads).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertAdditionalLoadTx()")
	}

	return nil
}

func (r *TeachingLoadRepository) UpdateAdditionalLoadsTx(ctx context.Context, tx pgx.Tx, loads []model.AdditionalLoad) error {
	for _, load := range loads {
		stmt, args := table.AdditionalLoad.
			UPDATE(
				table.AdditionalLoad.Name,
				table.AdditionalLoad.Volume,
				table.AdditionalLoad.Comment,
			).
			SET(
				load.Name,
				load.Volume,
				load.Comment,
			).
			WHERE(table.AdditionalLoad.LoadID.EQ(postgres.UUID(load.LoadID))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpdateAdditionalLoadTx()")
		}
	}

	return nil
}

func (r *TeachingLoadRepository) DeleteAdditionalLoadsTx(ctx context.Context, tx pgx.Tx, additionalIDs []uuid.UUID) error {
	var exps []postgres.Expression
	for _, id := range additionalIDs {
		exp := postgres.Expression(postgres.UUID(id))

		exps = append(exps, exp)
	}

	stmt, args := table.AdditionalLoad.
		DELETE().
		WHERE(table.AdditionalLoad.LoadID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteAdditionalLoadTx()")
	}

	return nil
}

//func (r *TeachingLoadRepository) GetTeachingLoadsTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]models.TeachingLoad, error) {
//	stmt, args := table.TeachingLoadStatus.
//		SELECT(
//			table.TeachingLoadStatus.StudentID,
//			table.TeachingLoadStatus.Semester,
//			table.TeachingLoadStatus.Status.AS("teaching_load.approval_status"),
//			table.TeachingLoadStatus.UpdatedAt,
//			table.TeachingLoadStatus.AcceptedAt,
//			table.ClassroomLoad.AllColumns,
//			table.IndividualStudentsLoad.AllColumns,
//			table.AdditionalLoad.AllColumns,
//		).
//		FROM(table.TeachingLoadStatus.
//			LEFT_JOIN(table.ClassroomLoad, table.TeachingLoadStatus.LoadsID.EQ(table.ClassroomLoad.TLoadID)).
//			LEFT_JOIN(table.IndividualStudentsLoad, table.TeachingLoadStatus.LoadsID.EQ(table.IndividualStudentsLoad.TLoadID)).
//			LEFT_JOIN(table.AdditionalLoad, table.TeachingLoadStatus.LoadsID.EQ(table.AdditionalLoad.TLoadID)),
//		).
//		WHERE(table.TeachingLoadStatus.StudentID.EQ(postgres.UUID(studentID))).
//		Sql()
//
//	rows, err := tx.Query(ctx, stmt, args...)
//	if err != nil {
//		return nil, errors.Wrap(err, "GetTeachingLoadTx()")
//	}
//	defer rows.Close()
//
//	loads := make([]models.TeachingLoad, 0, 10)
//
//	for rows.Next() {
//		load := models.TeachingLoad{}
//
//		if err := scanTeachingLoadStatus(rows, &load); err != nil {
//			return nil, errors.Wrap(err, "GetTeachingLoadTx(): scanning row")
//		}
//
//		loads = append(loads, load)
//	}
//
//	return loads, nil
//}

func scanTeachingLoadStatusStatus(row pgx.Row, target *model.TeachingLoadStatus) error {
	return row.Scan(
		&target.LoadsID,
		&target.StudentID,
		&target.Semester,
		&target.Status,
		&target.UpdatedAt,
		&target.AcceptedAt,
	)
}

//func scanTeachingLoadStatus(row pgx.Row, target *models.TeachingLoad) error {
//	return row.Scan(
//		&target.StudentID,
//		&target.Semester,
//		&target.ApprovalStatus,
//		&target.UpdatedAt,
//		&target.AcceptedAt,
//		&target.ClassroomLoad.LoadID,
//		&target.ClassroomLoad.TLoadID,
//		&target.ClassroomLoad.Hours,
//		&target.ClassroomLoad.LoadType,
//		&target.ClassroomLoad.MainTeacher,
//		&target.ClassroomLoad.GroupName,
//		&target.ClassroomLoad.SubjectName,
//		&target.IndividualStudentsLoad.LoadID,
//		&target.IndividualStudentsLoad.TLoadID,
//		&target.IndividualStudentsLoad.LoadType,
//		&target.IndividualStudentsLoad.StudentsAmount,
//		&target.IndividualStudentsLoad.Comment,
//		&target.AdditionalLoad.LoadID,
//		&target.AdditionalLoad.TLoadID,
//		&target.AdditionalLoad.Name,
//		&target.AdditionalLoad.Volume,
//		&target.AdditionalLoad.Comment,
//	)
//}

func (r *TeachingLoadRepository) GetClassroomLoadsTx(ctx context.Context, tx pgx.Tx, loadsIDs []uuid.UUID) ([]model.ClassroomLoad, error) {
	idExpressions := make([]postgres.Expression, 0, len(loadsIDs))

	for _, id := range loadsIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.ClassroomLoad.
		SELECT(table.ClassroomLoad.AllColumns).
		WHERE(table.ClassroomLoad.TLoadID.IN(idExpressions...)).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetClassroomLoadsTx()")
	}

	loads := make([]model.ClassroomLoad, 0, len(loadsIDs))

	for rows.Next() {
		load := model.ClassroomLoad{}
		if err := scanClassroomLoad(rows, &load); err != nil {
			return nil, errors.Wrap(err, "GetClassroomLoadsTx()")
		}

		loads = append(loads, load)
	}

	return loads, nil
}

func (r *TeachingLoadRepository) GetAdditionalLoadsTx(ctx context.Context, tx pgx.Tx, loadsIDs []uuid.UUID) ([]model.AdditionalLoad, error) {
	idExpressions := make([]postgres.Expression, 0, len(loadsIDs))

	for _, id := range loadsIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.AdditionalLoad.
		SELECT(table.AdditionalLoad.AllColumns).
		WHERE(table.AdditionalLoad.TLoadID.IN(idExpressions...)).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetAdditionalLoadsTx()")
	}

	loads := make([]model.AdditionalLoad, 0, len(loadsIDs))

	for rows.Next() {
		load := model.AdditionalLoad{}
		if err := scanAdditionalLoad(rows, &load); err != nil {
			return nil, errors.Wrap(err, "GetAdditionalLoadsTx()")
		}

		loads = append(loads, load)
	}

	return loads, nil
}

func (r *TeachingLoadRepository) GetIndividualLoadsTx(ctx context.Context, tx pgx.Tx, loadsIDs []uuid.UUID) ([]model.IndividualStudentsLoad, error) {
	idExpressions := make([]postgres.Expression, 0, len(loadsIDs))

	for _, id := range loadsIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.IndividualStudentsLoad.
		SELECT(table.IndividualStudentsLoad.AllColumns).
		WHERE(table.IndividualStudentsLoad.TLoadID.IN(idExpressions...)).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetIndividualLoadsTx()")
	}

	loads := make([]model.IndividualStudentsLoad, 0, len(loadsIDs))

	for rows.Next() {
		load := model.IndividualStudentsLoad{}
		if err := scanIndividualLoad(rows, &load); err != nil {
			return nil, errors.Wrap(err, "GetIndividualLoadsTx()")
		}

		loads = append(loads, load)
	}

	return loads, nil
}

func (r *TeachingLoadRepository) GetTeachingLoadStatusIDs(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]uuid.UUID, error) {
	stmt, args := table.TeachingLoadStatus.
		SELECT(table.TeachingLoadStatus.LoadsID).
		WHERE(table.ScientificWorksStatus.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetTeachingLoadStatusIDs()")
	}

	ids := make([]uuid.UUID, 0, 8)

	for rows.Next() {
		id := uuid.UUID{}
		if err := rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "GetTeachingLoadStatusIDs(): scanning row")
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func scanClassroomLoad(row pgx.Row, target *model.ClassroomLoad) error {
	return row.Scan(
		&target.LoadID,
		&target.TLoadID,
		&target.Hours,
		&target.LoadType,
		&target.MainTeacher,
		&target.GroupName,
		&target.SubjectName,
	)
}

func scanAdditionalLoad(row pgx.Row, target *model.AdditionalLoad) error {
	return row.Scan(
		&target.LoadID,
		&target.TLoadID,
		&target.Name,
		&target.Volume,
		&target.Comment,
	)
}

func scanIndividualLoad(row pgx.Row, target *model.IndividualStudentsLoad) error {
	return row.Scan(
		&target.LoadID,
		&target.TLoadID,
		&target.LoadType,
		&target.StudentsAmount,
		&target.Comment,
	)
}
