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
	if err != nil {
		return nil, errors.Wrap(err, "transaction initialization")
	}

	commonInfo, err := s.studRepo.GetStudentCommonInfo(ctx, s.db, session.ClientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentCommonInfo()")
	}

	plan, err := s.semesterRepo.GetSemesterProgress(ctx, s.db, session.ClientID)
	if err != nil {
		return nil, errors.Wrap(err, "GetSemesterProgress()")
	}

	planMap := make(map[string]*models.StudentDissertationPlan, len(plan))

	for _, chapter := range plan {
		planMap[chapter.Name] = chapter
	}

	return &models.DissertationPage{
		DissertationPlan: planMap,
		CommonInfo:       *commonInfo,
	}, nil
}
