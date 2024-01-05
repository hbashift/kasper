package admin

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/admin/mapping"
)

func (s *Service) GetStudSupPairs(ctx context.Context, token string) (*mapping.GetStudSupPairs, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Admin]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}

	sups, err := s.supRepo.GetSupervisors(ctx, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Admin]")
	}

	var supervisors []*mapping.SupervisorInfo

	for _, sup := range sups {
		supervisor := mapping.SupervisorInfo{
			SupervisorName: sup.FullName,
			SupervisorID:   sup.SupervisorID,
		}

		supervisors = append(supervisors, &supervisor)
	}

	pairs, err := s.studSupRepo.GetPairs(ctx, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Admin]")
	}

	return &mapping.GetStudSupPairs{
		Pairs:       pairs,
		Supervisors: supervisors,
	}, nil
}
