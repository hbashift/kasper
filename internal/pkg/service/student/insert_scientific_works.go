package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *Service) InsertScientificWorks(ctx context.Context, token string, works []*mapping.ScientificWork) error {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	var worksDomain []*model.ScientificWork

	for _, work := range works {
		workDomain := mapping.MapScientificWorkToDomain(work, session)
		workDomain.WorkID = uuid.New()
		worksDomain = append(worksDomain, workDomain)
	}

	err = s.scienceRepo.InsertStudentScientificWorks(ctx, s.db, worksDomain)
	if err != nil {
		return err
	}

	return nil
}
