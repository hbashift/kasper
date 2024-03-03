package new_repo

import (
	"context"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"
	"uir_draft/internal/pkg/models"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type ScientificRepository struct{}

func NewScientificRepository() *ScientificRepository {
	return &ScientificRepository{}
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
				table.Publications.Index,
				table.Publications.Impact,
				table.Publications.Status,
				table.Publications.OutputData,
				table.Publications.CoAuthors,
				table.Publications.Volume,
			).
			SET(
				publication.Name,
				publication.Index,
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
				table.Conferences.Index,
				table.Conferences.ConferenceName,
				table.Conferences.ReportName,
				table.Conferences.Location,
				table.Conferences.ReportedAt,
			).
			SET(
				conference.ConferenceID,
				conference.Status,
				conference.Index,
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
		WHERE(table.Conferences.ConferenceID.IN(exps...)).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteConferencesTx()")
	}

	return nil
}

func (r *ScientificRepository) GetScientificWorksTx(ctx context.Context, tx pgx.Tx, studentID uuid.UUID) ([]models.ScientificWork, error) {
	stmt, args := table.ScientificWorksStatus.
		SELECT(
			table.ScientificWorksStatus.WorksID,
			table.ScientificWorksStatus.Semester,
			table.ScientificWorksStatus.StudentID,
			table.ScientificWorksStatus.Status.AS("scientific_works.approval_status"),
			table.ScientificWorksStatus.UpdatedAt,
			table.ScientificWorksStatus.AcceptedAt,
			table.Publications.AllColumns.Except(table.Publications.WorksID),
			table.Conferences.AllColumns.Except(table.Conferences.WorksID),
			table.ResearchProjects.AllColumns.Except(table.ResearchProjects.WorksID),
		).
		FROM(table.ScientificWorksStatus.
			LEFT_JOIN(table.Publications, table.ScientificWorksStatus.WorksID.EQ(table.Publications.WorksID)).
			LEFT_JOIN(table.Conferences, table.ScientificWorksStatus.WorksID.EQ(table.Conferences.WorksID)).
			LEFT_JOIN(table.ResearchProjects, table.ScientificWorksStatus.WorksID.EQ(table.ResearchProjects.WorksID)),
		).
		WHERE(table.ScientificWorksStatus.StudentID.EQ(postgres.UUID(studentID))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetScientificWorksTx()")
	}
	defer rows.Close()

	works := make([]models.ScientificWork, 0, 10)

	for rows.Next() {
		work := models.ScientificWork{}

		if err := scanScientificWork(rows, &work); err != nil {
			return nil, errors.Wrap(err, "GetScientificWorksTx(): scanning rows")
		}

		works = append(works, work)
	}

	return works, nil
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

func scanScientificWork(row pgx.Row, target *models.ScientificWork) error {
	return row.Scan(
		&target.WorksID,
		&target.Semester,
		&target.StudentID,
		&target.ApprovalStatus,
		&target.UpdatedAt,
		&target.AcceptedAt,
		&target.Publication.PublicationID,
		&target.Publication.Name,
		&target.Publication.Index,
		&target.Publication.Impact,
		&target.Publication.Status,
		&target.Publication.OutputData,
		&target.Publication.CoAuthors,
		&target.Publication.Volume,
		&target.Conference.ConferenceID,
		&target.Conference.Status,
		&target.Conference.Index,
		&target.Conference.ConferenceName,
		&target.Conference.ReportName,
		&target.Conference.Location,
		&target.Conference.ReportedAt,
		&target.ResearchProject.ProjectID,
		&target.ResearchProject.ProjectName,
		&target.ResearchProject.StartAt,
		&target.ResearchProject.EndAt,
		&target.ResearchProject.AddInfo,
		&target.ResearchProject.Grantee,
	)
}

func scanPublication(row pgx.Row, target *models.Publication) error {
	return row.Scan(
		&target.PublicationID,
		&target.Name,
		&target.Index,
		&target.Impact,
		&target.Status,
		&target.OutputData,
		&target.CoAuthors,
		&target.Volume,
	)
}

func scanConference(row pgx.Row, target *models.Conference) error {
	return row.Scan(
		&target.ConferenceID,
		&target.Status,
		&target.Index,
		&target.ConferenceName,
		&target.ReportName,
		&target.Location,
		&target.ReportedAt,
	)
}

func scanResearchProject(row pgx.Row, target *models.ResearchProject) error {
	return row.Scan(
		&target.ProjectID,
		&target.ProjectName,
		&target.StartAt,
		&target.EndAt,
		&target.AddInfo,
		&target.Grantee,
	)
}

//func (r *ScientificRepository) GetPublicationsTx(ctx context.Context, tx pgx.Tx, publicationIDs []uuid.UUID) ([]model.Publications, error) {
//	idExpressions := make([]postgres.Expression, 0, len(publicationIDs))
//
//	for _, id := range publicationIDs {
//		idExp := postgres.UUID(id)
//
//		idExpressions = append(idExpressions, idExp)
//	}
//
//	stmt, args := table.Publications.
//		SELECT(table.Publications.AllColumns).
//		WHERE(table.Publications.PublicationID.IN(idExpressions...)).
//		Sql()
//
//	rows, err := tx.Query(ctx, stmt, args...)
//	if err != nil {
//		return nil, errors.Wrap(err, "GetPublicationsTx()")
//	}
//
//	publications := make([]model.Publications, 0, len(publicationIDs))
//
//	for rows.Next() {
//		publication := model.Publications{}
//		if err := scanPublication(rows, &publication); err != nil {
//			return nil, errors.Wrap(err, "GetPublicationsTx()")
//		}
//
//		publications = append(publications, publication)
//	}
//
//	return publications, nil
//}

//func (r *ScientificRepository) GetConferencesTx(ctx context.Context, tx pgx.Tx, conferenceIDs []uuid.UUID) ([]model.Conferences, error) {
//	idExpressions := make([]postgres.Expression, 0, len(conferenceIDs))
//
//	for _, id := range conferenceIDs {
//		idExp := postgres.UUID(id)
//
//		idExpressions = append(idExpressions, idExp)
//	}
//
//	stmt, args := table.Conferences.
//		SELECT(table.Conferences.AllColumns).
//		WHERE(table.Conferences.ConferenceID.IN(idExpressions...)).
//		Sql()
//
//	rows, err := tx.Query(ctx, stmt, args...)
//	if err != nil {
//		return nil, errors.Wrap(err, "GetPublicationsTx()")
//	}
//
//	conferences := make([]model.Conferences, 0, len(conferenceIDs))
//
//	for rows.Next() {
//		conference := model.Conferences{}
//		if err := scanConference(rows, &conference); err != nil {
//			return nil, errors.Wrap(err, "GetPublicationsTx()")
//		}
//
//		conferences = append(conferences, conference)
//	}
//
//	return conferences, nil
//}

//func (r *ScientificRepository) GetResearchProjectsTx(ctx context.Context, tx pgx.Tx, projectIDs []uuid.UUID) ([]model.ResearchProjects, error) {
//	idExpressions := make([]postgres.Expression, 0, len(projectIDs))
//
//	for _, id := range projectIDs {
//		idExp := postgres.UUID(id)
//
//		idExpressions = append(idExpressions, idExp)
//	}
//
//	stmt, args := table.Conferences.
//		SELECT(table.Conferences.AllColumns).
//		WHERE(table.Conferences.ConferenceID.IN(idExpressions...)).
//		Sql()
//
//	rows, err := tx.Query(ctx, stmt, args...)
//	if err != nil {
//		return nil, errors.Wrap(err, "GetPublicationsTx()")
//	}
//
//	projects := make([]model.ResearchProjects, 0, len(projectIDs))
//
//	for rows.Next() {
//		project := model.ResearchProjects{}
//		if err := scanResearchProject(rows, &project); err != nil {
//			return nil, errors.Wrap(err, "GetPublicationsTx()")
//		}
//
//		projects = append(projects, project)
//	}
//
//	return projects, nil
//}
