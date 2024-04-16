package admin

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) GetStudentSupervisorPairs(ctx context.Context) ([]models.StudentSupervisorPair, error) {
	pairs := make([]models.StudentSupervisorPair, 0)

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		dPairs, err := s.clientRepo.GetStudentSupervisorPairsTx(ctx, tx)
		if err != nil {
			return err
		}

		pairs = dPairs

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "GetStudentSupervisorPairs()")
	}

	return pairs, nil
}

func (s *Service) ChangeSupervisor(ctx context.Context, pairs []models.ChangeSupervisor) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		for _, pair := range pairs {
			err := s.clientRepo.SetNewSupervisorTx(ctx, tx, pair.StudentID, pair.SupervisorID)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "ChangeSupervisor()")
	}

	return nil
}

func (s *Service) SetStudentStudyingStatus(ctx context.Context, studentID uuid.UUID, status string) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		var dStatus model.StudentStatus
		if err := dStatus.Scan(status); err != nil {
			return err
		}

		student, err := s.clientRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if err = s.clientRepo.SetStudentStatusTx(ctx, tx, student.Status, dStatus, studentID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "SetStudentStudyingStatus()")
	}

	return nil
}

func (s *Service) GetSupervisors(ctx context.Context) ([]models.Supervisor, error) {
	sups := make([]models.Supervisor, 0)

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		dSups, err := s.clientRepo.GetSupervisorsTx(ctx, tx)
		if err != nil {
			return err
		}

		sups = dSups

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "GetSupervisors()")
	}

	return sups, nil
}

func (s *Service) GetStudentsList(ctx context.Context) ([]models.Student, error) {
	students := make([]models.Student, 0)

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		dStudents, err := s.clientRepo.GetStudentsList(ctx, tx)
		if err != nil {
			return err
		}

		students = dStudents

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "GetStudentsList()")
	}

	return students, nil
}
