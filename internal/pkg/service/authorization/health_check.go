package authorization

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/authorization/mapping"
)

func (s *Service) HealthCheck(ctx context.Context, token string) (*mapping.HealthCheck, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Authorization]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Authorization]")
	}

	userType, err := s.tokenRepo.GetUserTypeByToken(ctx, s.db, token)
	if err != nil {
		return nil, errors.Wrap(err, "[Authorization]")
	}

	return &mapping.HealthCheck{UserType: userType.String()}, nil
}
