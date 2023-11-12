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

	for _, work := range works {
		if work.WorkID == nil {
			workDomain := mapping.MapScientificWorkToDomain(work, session)
			workDomain.WorkID = uuid.New()

			err = s.scienceRepo.InsertStudentScientificWorks(ctx, s.db, workDomain)
			if err != nil {
				return err
			}

		} else {
			workDomain := mapping.MapScientificWorkToDomain(work, session)
			err = s.scienceRepo.UpdateStudentScientificWorks(ctx, s.db, workDomain)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
