package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
)

func (s *Service) GetScientificWorks(ctx context.Context, token string) (*mapping.Works, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}

	works, err := s.grepFromDBScientificWorks(ctx, session)
	if err != nil {
		return nil, err
	}

	years, err := s.studRepo.GetNumberOfYears(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	result := &mapping.Works{
		Works: works,
		Years: years,
	}

	return result, nil
}
