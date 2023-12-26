package student

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) GetDissertationIDs(ctx context.Context, token string) ([]*mapping.DissertationIDs, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Student]")
	}

	ids, err := s.dRepo.GetDissertationIDs(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, err
	}

	return mapping.MapIDsFromDomain(ids), nil
}
