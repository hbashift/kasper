package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
)

func (s *Service) UpdateScientificWorks(ctx context.Context, token string, works []*mapping.ScientificWork) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	for _, work := range works {
		workDomain := mapping.MapScientificWorkToDomain(work, session)

		err = s.scienceRepo.UpdateStudentScientificWorks(ctx, s.db, workDomain)
		if err != nil {
			return errors.Wrap(err, "[Student]")
		}
	}

	return nil
}
