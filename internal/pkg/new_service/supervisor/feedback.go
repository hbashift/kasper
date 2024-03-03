package supervisor

import (
	"context"

	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) UpsertFeedback(ctx context.Context, studentID uuid.UUID, request models.FeedbackRequest) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		feedback := models.MapFeedbackToDomain(request, studentID)

		err := s.dissertationRepo.UpsertFeedbackTx(ctx, tx, feedback)

		return err
	}); err != nil {
		return errors.Wrap(err, "UpsertFeedback()")
	}

	return nil
}
