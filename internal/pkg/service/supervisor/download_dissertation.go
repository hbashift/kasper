package supervisor

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

func (s *Service) DownloadDissertation(ctx context.Context, token string, info *mapping.DownloadDissertation) (*model.Dissertation, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Supervisor]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Supervisor]")
	}

	dissertation, err := s.dRepo.GetDissertationData(ctx, s.db, info.StudentID, info.Semester)
	if err != nil {
		return nil, errors.Wrap(err, "DownloadDissertation()")
	}

	return dissertation, nil
}
