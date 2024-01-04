package supervisor

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) GetScientificWorks(ctx context.Context, token string, studentID uuid.UUID) ([]*mapping.ScientificWork, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}

	scientificWorks, err := s.scienceRepo.GetScientificWorks(ctx, s.db, studentID)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	var jsonWorks []*mapping.ScientificWork
	for _, work := range scientificWorks {
		jsonWork := mapping.MapScientificWorkFromDomain(work)

		jsonWorks = append(jsonWorks, jsonWork)
	}

	return jsonWorks, nil
}
