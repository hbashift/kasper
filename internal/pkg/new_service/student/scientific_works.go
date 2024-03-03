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
	scientificWorks := make([]models.ScientificWork, 0, 10)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		works, err := s.scienceRepo.GetScientificWorksTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		scientificWorks = works

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetScientificWorks()")
	}

	return scientificWorks, nil
}

func (s *Service) UpsertPublications(ctx context.Context, studentID, workID uuid.UUID, semester int32, publications []models.Publication) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		dPublicationsInsert, dPublicationsUpdate, err := models.MapPublicationsToDomain(publications, workID)
		if err != nil {
			return err
		}

		err = s.scienceRepo.InsertPublicationsTx(ctx, tx, dPublicationsInsert)
		if err != nil {
			return err
		}

		err = s.scienceRepo.UpdatePublicationsTx(ctx, tx, dPublicationsUpdate)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpdatePublications()")
	}

	return nil
}

func (s *Service) UpsertConferences(ctx context.Context, studentID, workID uuid.UUID, semester int32, conferences []models.Conference) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		insert, update, err := models.MapConferencesToDomain(conferences, workID)
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

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpdateConferences()")
	}

	return nil
}

func (s *Service) UpsertResearchProjects(ctx context.Context, studentID, workID uuid.UUID, semester int32, projects []models.ResearchProject) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if student.ActualSemester != semester && !student.CanEdit {
			return models.ErrNotActualSemester
		}

		insert, update := models.MapResearchProjectToDomain(projects, workID)
		err = s.scienceRepo.InsertResearchProjectsTx(ctx, tx, insert)
		if err != nil {
			return err
		}

		err = s.scienceRepo.UpdateResearchProjectsTx(ctx, tx, update)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertResearchProjects()")
	}

	return nil
}

func (s *Service) DeletePublications(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentStatusTx(ctx, tx, studentID)
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

		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeletePublications()")
	}

	return nil
}

func (s *Service) DeleteConferences(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentStatusTx(ctx, tx, studentID)
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

		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteConferences()")
	}

	return nil
}

func (s *Service) DeleteResearchProjects(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentStatusTx(ctx, tx, studentID)
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

		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteResearchProjects()")
	}

	return nil
}
