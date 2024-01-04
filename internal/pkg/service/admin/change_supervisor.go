package admin

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/admin/mapping"
)

func (s *Service) ChangeSupervisor(ctx context.Context, token string, studSup *mapping.ChangeSupervisor) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "authentication error")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	if err = s.studSupRepo.ChangeSupervisor(ctx, s.db, studSup.StudentID, studSup.SupervisorID); err != nil {
		return errors.Wrap(err, "[Admin]")
	}

	return nil
}
