package authorization

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/authorization/mapping"
)

func (s *Service) ChangePassword(ctx context.Context, token string, info *mapping.ChangePassword) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Authorization]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return errors.Wrap(ErrNonValidToken, "[Authorization]")
	}

	client, err := s.clientRepo.GetClient(ctx, s.db, info.Email)
	if err != nil {
		return errors.Wrap(err, "[Authorization]")
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(info.OldPassword))
	if err != nil {
		return errors.Wrap(err, "[Authorization]")
	}

	err = s.clientRepo.ChangePassword(ctx, s.db, session.ClientID, info.NewPassword)
	if err != nil {
		return errors.Wrap(err, "[Authorization]")
	}

	return nil
}
