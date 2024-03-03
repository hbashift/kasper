package supervisor

import (
	"context"

	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) GetStudentList(ctx context.Context, supervisorID uuid.UUID) ([]models.StudentList, error) {
	response := make([]models.StudentList, 0)

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		list, err := s.client.GetSupervisorsStudentsTx(ctx, tx, supervisorID)
		if err != nil {
			return err
		}
		response = list

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "GetStudentList()")
	}

	return response, nil
}
