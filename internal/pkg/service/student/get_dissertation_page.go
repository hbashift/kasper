package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/pkg/errors"
)

var ErrNonValidToken = errors.New("token is expired")

func (s *Service) GetDissertationPage(ctx context.Context, token string) (*models.DissertationPage, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "authentication error")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "transaction initialization")
	}

	commonInfo, err := s.studRepo.GetStudentCommonInfo(ctx, tx, session.ClientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentCommonInfo()")
	}

	plan, err := s.semesterRepo.GetSemesterProgress(ctx, tx, session.ClientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetSemesterProgress()")
	}

	return &models.DissertationPage{
		DissertationPlan: plan,
		CommonInfo:       *commonInfo,
	}, nil
}
