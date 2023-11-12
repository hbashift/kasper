package student

import (
	"context"
	"time"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) UpsertSemesterPlan(ctx context.Context, token string, progress *mapping.Progress) error {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	var progressDomain []*model.SemesterProgress

	for _, prog := range progress.Progress {
		p := &model.SemesterProgress{
			SemesterProgressID: prog.SemesterProgressID,
			StudentID:          prog.StudentID,
			First:              prog.First,
			Second:             prog.Second,
			Third:              prog.Third,
			Forth:              prog.Forth,
			Fifth:              nil,
			Sixth:              nil,
			ProgressName:       model.ProgressType(prog.ProgressName),
			LastUpdated:        lo.ToPtr(time.Now()),
			ClientID:           prog.ClientID,
		}

		progressDomain = append(progressDomain, p)
	}

	return s.semesterRepo.UpsertSemesterPlan(ctx, s.db, progressDomain)
}
