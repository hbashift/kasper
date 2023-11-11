package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"

	"github.com/pkg/errors"
)

func (s *Service) InsertScientificWorks(ctx context.Context, token string, works []*model.ScientificWork) error {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	//TODO маппинг в доменную модель (должен приходить json)

	err = s.scienceRepo.InsertStudentScientificWorks(ctx, s.db, works)
	if err != nil {
		return err
	}

	return nil
}
