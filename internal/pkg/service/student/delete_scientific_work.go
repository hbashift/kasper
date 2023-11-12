package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
)

func (s *Service) DeleteScientificWork(ctx context.Context, token string, deleteIDs *mapping.DeleteWorkIDs) error {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return errors.Wrap(err, "authentication error")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	workIDs, err := mapping.MapWorkIDsToDomain(deleteIDs)

	if err = s.scienceRepo.DeleteStudentScientificWorks(ctx, s.db, workIDs); err != nil {
		return errors.Wrap(err, "[Student]")
	}

	return nil
}
