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

type ScientificRepository struct{}

func NewScientificRepository() *ScientificRepository {
	return &ScientificRepository{}
}

func (r *ScientificRepository) InitScientificWorkStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, years int32) error {
	works := make([]model.ScientificWorksStatus, 0, 8)

	for i := 1; i < int(years+1); i++ {
		work := model.ScientificWorksStatus{
			WorksID:    uuid.New(),
			StudentID:  studentID,
			Semester:   int32(i),
			Status:     model.ApprovalStatus_Empty,
			UpdatedAt:  time.Now(),
			AcceptedAt: nil,
		}

		works = append(works, work)
	}

	stmt, args := table.ScientificWorksStatus.
		INSERT().
		MODELS(works).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InitScientificWorkStatusTx()")
	}

	return nil
}

func (r *ScientificRepository) SetScientificWorkStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, status model.ApprovalStatus, semester int32, acceptedAt *time.Time) error {
	stmt, args := table.ScientificWorksStatus.
		UPDATE(
			table.ScientificWorksStatus.UpdatedAt,
			table.ScientificWorksStatus.AcceptedAt,
			table.ScientificWorksStatus.Status,
		).
		SET(
			time.Now(),
			acceptedAt,
			status,
		).
		WHERE(table.ScientificWorksStatus.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.ScientificWorksStatus.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "SetScientificWorkStatusTx()")
	}

	return nil
}

func (r *ScientificRepository) GetScientificWorksStatusTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]model.ScientificWorksStatus, error) {
	stmt, args := table.ScientificWorksStatus.
		SELECT(table.ScientificWorksStatus.AllColumns).
		WHERE(table.ScientificWorksStatus.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetScientificWorksStatusTx()")
	}
	defer rows.Close()

	works := make([]model.ScientificWorksStatus, 0, 8)

	for rows.Next() {
		work := model.ScientificWorksStatus{}

		if err := scanScientificWorksStatusStatus(rows, &work); err != nil {
			return nil, errors.Wrap(err, "GetScientificWorksStatusTx(): scanning row")
		}

		works = append(works, work)
	}

	return works, nil
}

func (r *ScientificRepository) GetScientificWorksStatusBySemesterTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID, semester int32) (model.ScientificWorksStatus, error) {
	stmt, args := table.ScientificWorksStatus.
		SELECT(table.ScientificWorksStatus.AllColumns).
		WHERE(table.ScientificWorksStatus.StudentID.EQ(postgres.UUID(studentID)).
			AND(table.ScientificWorksStatus.Semester.EQ(postgres.Int32(semester)))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)

	work := model.ScientificWorksStatus{}

	if err := scanScientificWorksStatusStatus(row, &work); err != nil {
		return model.ScientificWorksStatus{}, errors.Wrap(err, "GetScientificWorksStatusTx(): scanning row")
	}

	return work, nil
}

func (r *ScientificRepository) UpdateScientificWorksStatusTx(ctx context.Context, tx pgx.Tx, work model.ScientificWorksStatus) error {
	stmt, args := table.ScientificWorksStatus.
		UPDATE(
			table.ScientificWorksStatus.Status,
			table.ScientificWorksStatus.UpdatedAt,
			table.ScientificWorksStatus.AcceptedAt,
		).
		SET(
			work.Status,
			work.UpdatedAt,
			work.AcceptedAt,
		).
		WHERE(table.ScientificWorksStatus.WorksID.EQ(postgres.UUID(work.WorksID)).
			AND(table.ScientificWorksStatus.Semester.EQ(postgres.Int32(work.Semester)))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "UpdateScientificWorksStatusTx()")
	}

	return nil
}

func (r *ScientificRepository) InsertPublicationsTx(ctx context.Context, tx pgx.Tx, publications []model.Publications) error {
	stmt, args := table.Publications.
		INSERT().
		MODELS(publications).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "InsertPublicationsTx()")
	}

	return nil
}

func (r *ScientificRepository) UpdatePublicationsTx(ctx context.Context, tx pgx.Tx, publications []model.Publications) error {
	for _, publication := range publications {
		stmt, args := table.Publications.
			UPDATE(
				table.Publications.Name,
				table.Publications.Scopus,
				table.Publications.Rinc,
				table.Publications.Wac,
				table.Publications.Wos,
				table.Publications.Impact,
				table.Publications.Status,
				table.Publications.OutputData,
				table.Publications.CoAuthors,
				table.Publications.Volume,
			).
			SET(
				publication.Name,
				publication.Scopus,
				publication.Rinc,
				publication.Wac,
				publication.Wos,
				publication.Impact,
				publication.Status,
				publication.OutputData,
				publication.CoAuthors,
				publication.Volume,
			).
			WHERE(table.Publications.PublicationID.EQ(postgres.UUID(publication.PublicationID))).
			Sql()

		_, err := tx.Exec(ctx, stmt, args...)
		if err != nil {
			return errors.Wrap(err, "UpdatePublicationsTx()")
		}
	}

	return nil
}

func (r *ScientificRepository) DeletePublicationsTx(ctx context.Context, tx pgx.Tx, publicationsIDs []uuid.UUID) error {
	var exps []postgres.Expression
	for _, id := range publicationsIDs {
		exp := postgres.Expression(postgres.UUID(id))

		exps = append(exps, exp)
	}

	stmt, args := table.Publications.
		DELETE().
		WHERE(table.Publications.PublicationID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeletePublicationsTx()")
	}

	return nil
}

func (r *ScientificRepository) InsertConferencesTx(ctx context.Context, tx pgx.Tx, conferences []model.Conferences) error {
	stmt, args := table.Conferences.
		INSERT().
		MODELS(conferences).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "InsertConferences()")
	}

	return nil
}

func (r *ScientificRepository) UpdateConferencesTx(ctx context.Context, tx pgx.Tx, conferences []model.Conferences) error {
	for _, conference := range conferences {
		stmt, args := table.Conferences.
			UPDATE(
				table.Conferences.ConferenceID,
				table.Conferences.Status,
				table.Conferences.Scopus,
				table.Conferences.Rinc,
				table.Conferences.Wac,
				table.Conferences.Wos,
				table.Conferences.ConferenceName,
				table.Conferences.ReportName,
				table.Conferences.Location,
				table.Conferences.ReportedAt,
			).
			SET(
				conference.ConferenceID,
				conference.Status,
				conference.Scopus,
				conference.Rinc,
				conference.Wac,
				conference.Wos,
				conference.ConferenceName,
				conference.ReportName,
				conference.Location,
				conference.ReportedAt,
			).
			WHERE(table.Conferences.ConferenceID.EQ(postgres.UUID(conference.ConferenceID))).
			Sql()

		_, err := tx.Exec(ctx, stmt, args...)
		if err != nil {
			return errors.Wrap(err, "UpdateConferencesTx()")
		}
	}

	return nil
}

func (r *ScientificRepository) DeleteConferencesTx(ctx context.Context, tx pgx.Tx, conferencesIDs []uuid.UUID) error {
	var exps []postgres.Expression
	for _, id := range conferencesIDs {
		exp := postgres.Expression(postgres.UUID(id))

		exps = append(exps, exp)
	}

	stmt, args := table.Conferences.
		DELETE().
		WHERE(table.Conferences.ConferenceID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteConferencesTx()")
	}

	return nil
}

func (r *ScientificRepository) InsertResearchProjectsTx(ctx context.Context, tx pgx.Tx, projects []model.ResearchProjects) error {
	stmt, args := table.ResearchProjects.
		INSERT().
		MODELS(projects).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertResearchProjectsTx()")
	}

	return nil
}

func (r *ScientificRepository) UpdateResearchProjectsTx(ctx context.Context, tx pgx.Tx, projects []model.ResearchProjects) error {
	for _, project := range projects {
		stmt, args := table.ResearchProjects.
			UPDATE(
				table.ResearchProjects.ProjectName,
				table.ResearchProjects.StartAt,
				table.ResearchProjects.EndAt,
				table.ResearchProjects.AddInfo,
				table.ResearchProjects.Grantee,
			).
			SET(
				project.ProjectName,
				project.StartAt,
				project.EndAt,
				project.AddInfo,
				project.Grantee,
			).
			WHERE(table.ResearchProjects.ProjectID.EQ(postgres.UUID(project.ProjectID))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpdateResearchProjectsTx()")
		}
	}

	return nil
}

func (r *ScientificRepository) DeleteResearchProjectsTx(ctx context.Context, tx pgx.Tx, projectsIDs []uuid.UUID) error {
	var exps []postgres.Expression
	for _, id := range projectsIDs {
		exp := postgres.Expression(postgres.UUID(id))

		exps = append(exps, exp)
	}

	stmt, args := table.ResearchProjects.
		DELETE().
		WHERE(table.ResearchProjects.ProjectID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteConferencesTx()")
	}

	return nil
}

func scanScientificWorksStatusStatus(row pgx.Row, target *model.ScientificWorksStatus) error {
	return row.Scan(
		&target.WorksID,
		&target.StudentID,
		&target.Semester,
		&target.Status,
		&target.UpdatedAt,
		&target.AcceptedAt,
	)
}

func scanPublication(row pgx.Row, target *model.Publications) error {
	return row.Scan(
		&target.PublicationID,
		&target.WorksID,
		&target.Name,
		&target.Scopus,
		&target.Rinc,
		&target.Wac,
		&target.Wos,
		&target.Impact,
		&target.Status,
		&target.OutputData,
		&target.CoAuthors,
		&target.Volume,
	)
}

func scanConference(row pgx.Row, target *model.Conferences) error {
	return row.Scan(
		&target.ConferenceID,
		&target.WorksID,
		&target.Status,
		&target.Scopus,
		&target.Rinc,
		&target.Wac,
		&target.Wos,
		&target.ConferenceName,
		&target.ReportName,
		&target.Location,
		&target.ReportedAt,
	)
}

func scanResearchProject(row pgx.Row, target *model.ResearchProjects) error {
	return row.Scan(
		&target.ProjectID,
		&target.WorksID,
		&target.ProjectName,
		&target.StartAt,
		&target.EndAt,
		&target.AddInfo,
		&target.Grantee,
	)
}

func (r *ScientificRepository) GetPublicationsTx(ctx context.Context, tx pgx.Tx, worksIDs []uuid.UUID) ([]model.Publications, error) {
	idExpressions := make([]postgres.Expression, 0, len(worksIDs))

	for _, id := range worksIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.Publications.
		SELECT(table.Publications.AllColumns).
		WHERE(table.Publications.WorksID.IN(idExpressions...)).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetPublicationsTx()")
	}

	publications := make([]model.Publications, 0, len(worksIDs))

	for rows.Next() {
		publication := model.Publications{}
		if err := scanPublication(rows, &publication); err != nil {
			return nil, errors.Wrap(err, "GetPublicationsTx()")
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (r *ScientificRepository) GetConferencesTx(ctx context.Context, tx pgx.Tx, worksIDs []uuid.UUID) ([]model.Conferences, error) {
	idExpressions := make([]postgres.Expression, 0, len(worksIDs))

	for _, id := range worksIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.Conferences.
		SELECT(table.Conferences.AllColumns).
		WHERE(table.Conferences.WorksID.IN(idExpressions...)).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetConferencesTx()")
	}

	conferences := make([]model.Conferences, 0, len(worksIDs))

	for rows.Next() {
		conference := model.Conferences{}
		if err := scanConference(rows, &conference); err != nil {
			return nil, errors.Wrap(err, "GetConferencesTx()")
		}

		conferences = append(conferences, conference)
	}

	return conferences, nil
}

func (r *ScientificRepository) GetResearchProjectsTx(ctx context.Context, tx pgx.Tx, worksIDs []uuid.UUID) ([]model.ResearchProjects, error) {
	idExpressions := make([]postgres.Expression, 0, len(worksIDs))

	for _, id := range worksIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.ResearchProjects.
		SELECT(table.ResearchProjects.AllColumns).
		WHERE(table.ResearchProjects.WorksID.IN(idExpressions...)).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetResearchProjectsTx()")
	}

	projects := make([]model.ResearchProjects, 0, len(worksIDs))

	for rows.Next() {
		project := model.ResearchProjects{}
		if err := scanResearchProject(rows, &project); err != nil {
			return nil, errors.Wrap(err, "GetResearchProjectsTx()")
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func (r *ScientificRepository) GetScientificWorksStatusIDs(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]uuid.UUID, error) {
	stmt, args := table.ScientificWorksStatus.
		SELECT(table.ScientificWorksStatus.WorksID).
		WHERE(table.ScientificWorksStatus.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetScientificWorksStatusIDs()")
	}

	ids := make([]uuid.UUID, 0, 8)

	for rows.Next() {
		id := uuid.UUID{}
		if err := rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "GetScientificWorksStatusIDs(): scanning row")
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (r *ScientificRepository) InsertPatents(ctx context.Context, tx pgx.Tx, patents []model.Patents) error {
	stmt, args := table.Patents.
		INSERT().
		MODELS(patents).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertPatents()")
	}

	return nil
}

func (r *ScientificRepository) GetPatents(ctx context.Context, tx pgx.Tx, worksIDs []uuid.UUID) ([]model.Patents, error) {
	idExpressions := make([]postgres.Expression, 0, len(worksIDs))

	for _, id := range worksIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.Patents.
		SELECT(table.Patents.AllColumns).
		WHERE(table.Patents.WorksID.IN(idExpressions...)).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetPatents()")
	}

	patents := make([]model.Patents, 0)

	for rows.Next() {
		patent := model.Patents{}
		if err := scanPatent(rows, &patent); err != nil {
			return nil, errors.Wrap(err, "GetPatents()")
		}

		patents = append(patents, patent)
	}

	return patents, nil
}

func (r *ScientificRepository) UpdatePatents(ctx context.Context, tx pgx.Tx, patents []model.Patents) error {
	for _, patent := range patents {
		stmt, args := table.Patents.
			UPDATE(
				table.Patents.Name,
				table.Patents.RegistrationDate,
				table.Patents.Type,
				table.Patents.AddInfo,
			).
			SET(
				patent.Name,
				patent.RegistrationDate,
				patent.Type,
				patent.AddInfo,
			).
			WHERE(table.Patents.PatentID.EQ(postgres.UUID(patent.PatentID))).
			Sql()

		if _, err := tx.Exec(ctx, stmt, args...); err != nil {
			return errors.Wrap(err, "UpdatePatents()")
		}
	}

	return nil
}

func (r *ScientificRepository) DeletePatents(ctx context.Context, tx pgx.Tx, patentIDs []uuid.UUID) error {
	var exps []postgres.Expression
	for _, id := range patentIDs {
		exp := postgres.Expression(postgres.UUID(id))

		exps = append(exps, exp)
	}

	stmt, args := table.Patents.
		DELETE().
		WHERE(table.Patents.PatentID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeletePatents()")
	}

	return nil
}

func scanPatent(row pgx.Row, target *model.Patents) error {
	return row.Scan(
		&target.PatentID,
		&target.WorksID,
		&target.Name,
		&target.RegistrationDate,
		&target.Type,
		&target.AddInfo,
	)
}
