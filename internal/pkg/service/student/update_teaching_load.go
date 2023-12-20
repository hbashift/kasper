package student

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) UpdateTeachingLoad(ctx context.Context, token string, loads *mapping.TeachingLoad) error {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return errors.Wrap(ErrNonValidToken, "[Student]")
	}

	domainLoads, err := mapping.MapTeachingLoadToDomain(loads, session)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	return s.loadRepo.InsertStudentsTeachingLoad(ctx, s.db, domainLoads)
}
