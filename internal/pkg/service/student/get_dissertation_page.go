package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/pkg/errors"
)

var ErrNonValidToken = errors.New("token is expired")

func (s *Service) GetDissertationPage(ctx context.Context, token string) (*models.DissertationPage, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "authentication error")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}

	return s.studRepo.GetDissertation(ctx, session.ClientID)
}
