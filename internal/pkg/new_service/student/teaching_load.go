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
	teachingLoads := make([]models.TeachingLoad, 0, 10)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		loads, err := s.loadRepo.GetTeachingLoadsTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		teachingLoads = loads

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetTeachingLoad()")
	}

	return teachingLoads, nil
}

func (s *Service) UpsertClassroomLoad(ctx context.Context, studentID, tLoadID uuid.UUID, semester int32, loads []models.ClassroomLoad) error {
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

		insert, update, err := models.MapClassroomLoadToDomain(loads, tLoadID)
		if err != nil {
			return err
		}

		err = s.loadRepo.InsertClassroomLoadsTx(ctx, tx, insert)
		if err != nil {
			return err
		}

		err = s.loadRepo.UpdateClassroomLoadsTx(ctx, tx, update)
		if err != nil {
			return err
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertClassroomLoad()")
	}

	return nil
}

func (s *Service) UpsertIndividualLoad(ctx context.Context, studentID, tLoadID uuid.UUID, semester int32, loads []models.IndividualStudentsLoad) error {
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

		insert, update := models.MapIndividualWorkToDomain(loads, tLoadID)
		err = s.loadRepo.InsertIndividualLoadsTx(ctx, tx, insert)
		if err != nil {
			return err
		}

		err = s.loadRepo.UpdateIndividualLoadsTx(ctx, tx, update)
		if err != nil {
			return err
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertIndividualLoad()")
	}

	return nil
}

func (s *Service) UpsertAdditionalLoad(ctx context.Context, studentID, tLoadID uuid.UUID, semester int32, loads []models.AdditionalLoad) error {
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

		insert, update := models.MapAdditionalLoadToDomain(loads, tLoadID)
		err = s.loadRepo.InsertAdditionalLoadsTx(ctx, tx, insert)
		if err != nil {
			return err
		}

		err = s.loadRepo.UpdateAdditionalLoadsTx(ctx, tx, update)
		if err != nil {
			return err
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertAdditionalLoad()")
	}

	return nil
}

func (s *Service) DeleteClassroomLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
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

		err = s.loadRepo.DeleteClassroomLoadsTx(ctx, tx, loads)
		if err != nil {
			return err
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteClassroomLoad()")
	}

	return nil
}

func (s *Service) DeleteIndividualLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
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

		err = s.loadRepo.DeleteIndividualStudentsLoadsTx(ctx, tx, loads)
		if err != nil {
			return err
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertAdditionalLoad()")
	}

	return nil
}

func (s *Service) DeleteAdditionalLoad(ctx context.Context, studentID uuid.UUID, semester int32, loads []uuid.UUID) error {
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

		err = s.loadRepo.DeleteAdditionalLoadsTx(ctx, tx, loads)
		if err != nil {
			return err
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, model.ApprovalStatus_InProgress, semester, nil)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "DeleteAdditionalLoad()")
	}

	return nil
}
