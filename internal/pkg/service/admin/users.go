package admin

import (
	"context"
	"strings"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

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

func (s *Service) SetStudentFlags(ctx context.Context, students []models.SetStudentsFlags) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		for _, student := range students {
			var dStatus model.StudentStatus
			if err := dStatus.Scan(strings.TrimSpace(student.StudyingStatus)); err != nil {
				return err
			}

			if err := s.clientRepo.SetStudentFlags(ctx, tx, dStatus, student.CanEdit, student.StudentID); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "SetStudentFlags()")
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
		return nil, errors.Wrap(err, "GetSupervisorsStudents()")
	}

	return students, nil
}
