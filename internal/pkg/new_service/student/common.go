package student

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) AllToStatus(ctx context.Context, studentID uuid.UUID, status string) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		dStatus, err := models.MapApprovalStatusToDomain(status)
		if err != nil {
			return errors.Wrap(err, "AllToStatus()")
		}

		err = s.dissertationRepo.SetDissertationStatusTx(ctx, tx, student.StudentID, dStatus, student.ActualSemester)
		if err != nil {
			return err
		}

		//err = s.dissertationRepo.SetDissertationTitleStatusTx(ctx, tx, student.StudentID, dStatus, student.ActualSemester, nil)
		//if err != nil {
		//	return err
		//}

		err = s.dissertationRepo.SetSemesterProgressStatusTx(ctx, tx, student.StudentID, dStatus, nil)
		if err != nil {
			return err
		}

		err = s.scienceRepo.SetScientificWorkStatusTx(ctx, tx, student.StudentID, dStatus, student.ActualSemester, nil)
		if err != nil {
			return err
		}

		err = s.loadRepo.SetTeachingLoadStatusTx(ctx, tx, student.StudentID, dStatus, student.ActualSemester, nil)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, dStatus, student.StudyingStatus, student.StudentID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "AllToStatus()")
	}

	return nil
}

func (s *Service) GetStudentStatus(ctx context.Context, studentID uuid.UUID) (models.Student, error) {
	var student models.Student

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		var err error
		student, err = s.studRepo.GetStudentStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return models.Student{}, errors.Wrap(err, "GetStudentStatus()")
	}

	return student, nil
}

func (s *Service) SetStudentStatus(ctx context.Context, studentID uuid.UUID, status model.ApprovalStatus) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if err = s.studRepo.SetStudentStatusTx(ctx, tx, status, student.StudyingStatus, studentID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "SetStudentStatus()")
	}

	return nil
}
