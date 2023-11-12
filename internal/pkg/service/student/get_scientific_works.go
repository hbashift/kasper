package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
)

func (s *Service) GetScientificWorks(ctx context.Context, token string) ([]*mapping.ScientificWork, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}

	works, err := s.scienceRepo.GetScientificWorks(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	var jsonWorks []*mapping.ScientificWork
	for _, work := range works {
		jsonWork := mapping.MapScientificWorkFromDomain(work)

		jsonWorks = append(jsonWorks, jsonWork)
	}

	return jsonWorks, nil
}
