package student

import (
	"context"
	"time"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
)

func (s *Service) DeleteScientificWork(ctx context.Context, token string, deleteIDs *mapping.DeleteWorkIDs) ([]*mapping.ScientificWork, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Student]")
	}

	workIDs, err := mapping.MapWorkIDsToDomain(deleteIDs)

	// Deleting scientific works
	if err = s.scienceRepo.DeleteStudentScientificWorks(ctx, s.db, workIDs); err != nil {
		return nil, errors.Wrap(err, "[Student]")
	}

	time.Sleep(time.Second * 1)
	// Grepping scientific works
	return s.grepFromDBScientificWorks(ctx, session)
}
