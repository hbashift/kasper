package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"

	"github.com/pkg/errors"
)

func (s *Service) GetScientificWorks(ctx context.Context, token string) ([]*model.ScientificWork, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}

	//TODO маппинг в доменную модель (должен приходить json)

	works, err := s.scienceRepo.GetScientificWorks(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	return works, nil
}
