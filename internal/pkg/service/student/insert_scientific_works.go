package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *Service) InsertScientificWorks(ctx context.Context, token string, works []*mapping.ScientificWork) ([]*mapping.ScientificWork, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Student]")
	}

	// Upserting scientific works
	for _, work := range works {
		switch {
		case work.WorkID == nil:
			workDomain := mapping.MapScientificWorkToDomain(work, session)
			workDomain.WorkID = uuid.New()

			err = s.scienceRepo.InsertStudentScientificWorks(ctx, s.db, workDomain)
			if err != nil {
				return nil, errors.Wrap(err, "[Student]")
			}

		case work.WorkID != nil:
			workDomain := mapping.MapScientificWorkToDomain(work, session)
			err = s.scienceRepo.UpdateStudentScientificWorks(ctx, s.db, workDomain)
			if err != nil {
				return nil, errors.Wrap(err, "[Student]")
			}

		}
	}

	// Grepping all scientific works
	return s.grepFromDBScientificWorks(ctx, session)
}
