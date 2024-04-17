package student

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) GetTeachingLoad(ctx context.Context, studentID uuid.UUID) ([]models.TeachingLoad, error) {
	dLoads := make([]model.TeachingLoadStatus, 0, 10)
	var (
		classroom  []models.ClassroomLoad
		additional []models.AdditionalLoad
		individual []models.IndividualStudentsLoad
	)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		loads, err := s.loadRepo.GetTeachingLoadStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}
		dLoads = loads

		loadsIDs, err := s.loadRepo.GetTeachingLoadStatusIDs(ctx, tx, studentID)
		if err != nil {
			return err
		}

		dClassroom, err := s.loadRepo.GetClassroomLoadsTx(ctx, tx, loadsIDs)
		if err != nil {
			return err
		}

		dAdditional, err := s.loadRepo.GetAdditionalLoadsTx(ctx, tx, loadsIDs)
		if err != nil {
			return err
		}

		dIndividual, err := s.loadRepo.GetIndividualLoadsTx(ctx, tx, loadsIDs)
		if err != nil {
			return err
		}

		classroom = models.MapClassroomLoadFromDomain(dClassroom)
		additional = models.MapAdditionalLoadFromDomain(dAdditional)
		individual = models.MapIndividualWorkFromDomain(dIndividual)

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetTeachingLoad()")
	}

	teachingLoads := models.ConvertTeachingLoadsToResponse(studentID, dLoads, classroom, additional, individual)

	return teachingLoads, nil
}

func (s *Service) TeachingLoadToStatus(ctx context.Context, studentID uuid.UUID, status model.ApprovalStatus, semester int32) error {
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

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, studentID, status, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, status, student.StudyingStatus, studentID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "TeachingLoadToStatus()")
	}

	return nil
}

func (s *Service) UpsertClassroomLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []models.ClassroomLoad) error {
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

		loadsStatus, err := s.loadRepo.GetTeachingLoadStatusBySemesterTx(ctx, tx, studentID, semester)
		if err != nil {
			return err
		}

		insert, update, err := models.MapClassroomLoadToDomain(loads, loadsStatus.LoadsID)
		if err != nil {
			return err
		}

		if len(insert) != 0 {
			err = s.loadRepo.InsertClassroomLoadsTx(ctx, tx, insert)
			if err != nil {
				return err
			}
		}

		if len(update) != 0 {
			err = s.loadRepo.UpdateClassroomLoadsTx(ctx, tx, update)
			if err != nil {
				return err
			}
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertClassroomLoad()")
	}

	return nil
}

func (s *Service) UpsertIndividualLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []models.IndividualStudentsLoad) error {
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

		loadsStatus, err := s.loadRepo.GetTeachingLoadStatusBySemesterTx(ctx, tx, studentID, semester)
		if err != nil {
			return err
		}

		insert, update, err := models.MapIndividualWorkToDomain(loads, loadsStatus.LoadsID)
		if err != nil {
			return errors.Wrap(err, "UpsertIndividualLoad()")
		}

		if len(insert) != 0 {
			err = s.loadRepo.InsertIndividualLoadsTx(ctx, tx, insert)
			if err != nil {
				return err
			}
		}

		if len(update) != 0 {
			err = s.loadRepo.UpdateIndividualLoadsTx(ctx, tx, update)
			if err != nil {
				return err
			}
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertIndividualLoad()")
	}

	return nil
}

func (s *Service) UpsertAdditionalLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []models.AdditionalLoad) error {
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

		loadsStatus, err := s.loadRepo.GetTeachingLoadStatusBySemesterTx(ctx, tx, studentID, semester)
		if err != nil {
			return err
		}

		insert, update := models.MapAdditionalLoadToDomain(loads, loadsStatus.LoadsID)

		if len(insert) != 0 {
			err = s.loadRepo.InsertAdditionalLoadsTx(ctx, tx, insert)
			if err != nil {
				return err
			}
		}

		if len(update) != 0 {
			err = s.loadRepo.UpdateAdditionalLoadsTx(ctx, tx, update)
			if err != nil {
				return err
			}
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertAdditionalLoad()")
	}

	return nil
}

func (s *Service) DeleteClassroomLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
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

		if len(loads) != 0 {
			err = s.loadRepo.DeleteClassroomLoadsTx(ctx, tx, loads)
			if err != nil {
				return err
			}
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteClassroomLoad()")
	}

	return nil
}

func (s *Service) DeleteIndividualLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
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

		if len(loads) != 0 {
			err = s.loadRepo.DeleteIndividualStudentsLoadsTx(ctx, tx, loads)
			if err != nil {
				return err
			}
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertAdditionalLoad()")
	}

	return nil
}

func (s *Service) DeleteAdditionalLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
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

		if len(loads) != 0 {
			err = s.loadRepo.DeleteAdditionalLoadsTx(ctx, tx, loads)
			if err != nil {
				return err
			}
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudyingStatus, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteAdditionalLoad()")
	}

	return nil
}
