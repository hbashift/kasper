package student

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) DeleteTeachingLoad(ctx context.Context, token string, deleteIDs *mapping.DeleteWorkIDs) (*mapping.TeachingLoad, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Student]")
	}

	ids, err := mapping.MapWorkIDsToDomain(deleteIDs)

	err = s.loadRepo.DeleteTeachingLoad(ctx, s.db, ids)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	return s.grepFromDBTeachingLoad(ctx, session)
}
