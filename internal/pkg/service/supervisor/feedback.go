package supervisor

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) UpsertFeedback(ctx context.Context, studentID, supervisorID uuid.UUID, request models.FeedbackRequest) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		feedback := models.MapFeedbackToDomain(request, studentID)

		err := s.dissertationRepo.UpsertFeedbackTx(ctx, tx, feedback)
		if err != nil {
			return err
		}

		supMark := model.SupervisorMarks{
			MarkID:       uuid.New(),
			StudentID:    studentID,
			Mark:         request.Mark,
			Semester:     request.Semester,
			SupervisorID: supervisorID,
		}

		if err := s.marksRepo.UpsertStudentsSupervisorMark(ctx, tx, supMark); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "UpsertFeedback()")
	}

	return nil
}

func (s *Service) UpsertSupervisorMark(ctx context.Context, studentID, supervisorID uuid.UUID, semester, mark int32) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		supMark := model.SupervisorMarks{
			MarkID:       uuid.New(),
			StudentID:    studentID,
			Mark:         mark,
			Semester:     semester,
			SupervisorID: supervisorID,
		}

		if err := s.marksRepo.UpsertStudentsSupervisorMark(ctx, tx, supMark); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "UpsertSupervisorMark()")
	}

	return nil
}
