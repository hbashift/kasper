package student

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) GetSupervisors(ctx context.Context, token string) (*mapping.Supervisors, error) {
	_, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	sups, err := s.supRepo.GetSupervisors(ctx, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	var supervisors []*mapping.Supervisor

	for _, sup := range sups {
		supervisor := mapping.Supervisor{
			Name:         sup.FullName,
			SupervisorID: sup.SupervisorID,
		}

		supervisors = append(supervisors, &supervisor)
	}

	return &mapping.Supervisors{Supervisors: supervisors}, nil
}
