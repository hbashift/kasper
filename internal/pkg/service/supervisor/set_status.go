package supervisor

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

func (s *Service) SetStatus(ctx context.Context, token string, info *mapping.SetStatus) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Supervisor]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return errors.Wrap(ErrNonValidToken, "[Supervisor]")
	}

	return s.dRepo.SetStatus(ctx, s.db, session.KasperID, info.Semester, model.DissertationStatus(info.Status))
}
