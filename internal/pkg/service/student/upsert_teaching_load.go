package student

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) UpsertTeachingLoad(ctx context.Context, token string, loads *mapping.TeachingLoad) (*mapping.TeachingLoad, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Student]")
	}
	var domainLoads []*model.TeachingLoad

	for _, load := range loads.Array {
		domainLoad, err := mapping.MapTeachingLoadToDomain(&load, session)
		if err != nil {
			return nil, errors.Wrap(err, "[Student]")
		}

		switch {
		case load.LoadID == nil:
			domainLoad.LoadID = uuid.New()
			domainLoads = append(domainLoads, domainLoad)

		case load.LoadID != nil:
			domainLoads = append(domainLoads, domainLoad)
		}
	}

	err = s.loadRepo.UpsertStudentsTeachingLoad(ctx, s.db, domainLoads)
	if err != nil {
		return nil, err
	}

	return s.grepFromDBTeachingLoad(ctx, session)
}
