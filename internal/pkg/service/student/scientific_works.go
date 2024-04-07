package student

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) GetScientificWorks(ctx context.Context, studentID uuid.UUID) ([]models.ScientificWork, error) {
	dWorks := make([]model.ScientificWorksStatus, 0, 10)
	var (
		publications []models.Publication
		conferences  []models.Conference
		projects     []models.ResearchProject
	)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		works, err := s.scienceRepo.GetScientificWorksStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}
		dWorks = works

		worksIDs, err := s.scienceRepo.GetScientificWorksStatusIDs(ctx, tx, studentID)
		if err != nil {
			return err
		}

		dPublications, err := s.scienceRepo.GetPublicationsTx(ctx, tx, worksIDs)
		if err != nil {
			return err
		}

		dConferences, err := s.scienceRepo.GetConferencesTx(ctx, tx, worksIDs)
		if err != nil {
			return err
		}
		dProjects, err := s.scienceRepo.GetResearchProjectsTx(ctx, tx, worksIDs)
		if err != nil {
			return err
		}

		publications = models.MapPublicationsFromDomain(dPublications)
		conferences = models.MapConferencesFromDomain(dConferences)
		projects = models.MapResearchProjectFromDomain(dProjects)

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetScientificWorks()")
	}

	scientificWorks := models.ConvertScientificWorksToResponse(studentID, dWorks, publications, conferences, projects)

	return scientificWorks, nil
}

func (s *Service) ScientificWorksToStatus(ctx context.Context, studentID uuid.UUID, status model.ApprovalStatus, semester int32) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, studentID, status, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, status, student.StudyingStatus, studentID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "ScientificWorksToStatus()")
	}

	return nil
}

func (s *Service) UpsertPublications(ctx context.Context, studentID uuid.UUID, semester int32, publications []models.Publication) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		worksStatus, err := s.scienceRepo.GetScientificWorksStatusBySemesterTx(ctx, tx, studentID, semester)
		if err != nil {
			return err
		}

		dPublicationsInsert, dPublicationsUpdate, err := models.MapPublicationsToDomain(publications, worksStatus.WorksID)
		if err != nil {
			return err
		}

		if len(dPublicationsInsert) != 0 {
			err = s.scienceRepo.InsertPublicationsTx(ctx, tx, dPublicationsInsert)
			if err != nil {
				return err
			}
		}

		if len(dPublicationsUpdate) != 0 {
			err = s.scienceRepo.UpdatePublicationsTx(ctx, tx, dPublicationsUpdate)
			if err != nil {
				return err
			}
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpdatePublications()")
	}

	return nil
}

func (s *Service) UpsertConferences(ctx context.Context, studentID uuid.UUID, semester int32, conferences []models.Conference) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		worksStatus, err := s.scienceRepo.GetScientificWorksStatusBySemesterTx(ctx, tx, studentID, semester)
		if err != nil {
			return err
		}

		insert, update, err := models.MapConferencesToDomain(conferences, worksStatus.WorksID)
		if err != nil {
			return err
		}

		err = s.scienceRepo.InsertConferencesTx(ctx, tx, insert)
		if err != nil {
			return err
		}

		err = s.scienceRepo.UpdateConferencesTx(ctx, tx, update)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpdateConferences()")
	}

	return nil
}

func (s *Service) UpsertResearchProjects(ctx context.Context, studentID uuid.UUID, semester int32, projects []models.ResearchProject) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		worksStatus, err := s.scienceRepo.GetScientificWorksStatusBySemesterTx(ctx, tx, studentID, semester)
		if err != nil {
			return err
		}

		insert, update := models.MapResearchProjectToDomain(projects, worksStatus.WorksID)
		err = s.scienceRepo.InsertResearchProjectsTx(ctx, tx, insert)
		if err != nil {
			return err
		}

		err = s.scienceRepo.UpdateResearchProjectsTx(ctx, tx, update)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertResearchProjects()")
	}

	return nil
}

func (s *Service) DeletePublications(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		err = s.scienceRepo.DeletePublicationsTx(ctx, tx, loads)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeletePublications()")
	}

	return nil
}

func (s *Service) DeleteConferences(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		err = s.scienceRepo.DeleteConferencesTx(ctx, tx, loads)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteConferences()")
	}

	return nil
}

func (s *Service) DeleteResearchProjects(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		err = s.scienceRepo.DeleteResearchProjectsTx(ctx, tx, loads)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteResearchProjects()")
	}

	return nil
}
