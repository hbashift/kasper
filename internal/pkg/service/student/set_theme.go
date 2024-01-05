package student

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
)

func (s *Service) SetTheme(ctx context.Context, token, theme string) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return errors.Wrap(ErrNonValidToken, "[Student]")
	}

	return s.studRepo.SetTheme(ctx, s.db, session.KasperID, theme)
}
