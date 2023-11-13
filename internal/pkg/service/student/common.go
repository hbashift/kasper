package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
)

func (s *Service) grepFromDBScientificWorks(ctx context.Context, session *model.AuthorizationToken) ([]*mapping.ScientificWork, error) {
	scientificWorks, err := s.scienceRepo.GetScientificWorks(ctx, s.db, session.KasperID)
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
