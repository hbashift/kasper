package supervisor

import (
	"context"

	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) GetSupervisorsStudents(ctx context.Context, supervisorID uuid.UUID) ([]models.Student, error) {
	response := make([]models.Student, 0)

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		list, err := s.client.GetSupervisorsStudentsTx(ctx, tx, supervisorID)
		if err != nil {
			return err
		}
		response = list

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "GetSupervisorsStudents()")
	}

	return response, nil
}

func (s *Service) GetSupervisorProfile(ctx context.Context, supervisorID uuid.UUID) (models.Supervisor, error) {
	var supervisor models.Supervisor

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		dSup, err := s.client.GetSupervisorTx(ctx, tx, supervisorID)
		if err != nil {
			return err
		}

		supervisor = dSup
		return nil
	}); err != nil {
		return models.Supervisor{}, errors.Wrap(err, "GetSupervisorProfile()")
	}

	return supervisor, nil
}
