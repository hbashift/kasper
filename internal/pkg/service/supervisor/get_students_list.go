package supervisor

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

func (s *Service) GetStudentsList(ctx context.Context, token string) (*mapping.ListOfStudents, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "[Supervisor]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, errors.Wrap(ErrNonValidToken, "[Supervisor]")
	}

	domainStudents, err := s.studRepo.GetListOfStudents(ctx, s.db, &session.KasperID)
	if err != nil {
		return nil, errors.Wrap(err, "[Supervisor]")
	}

	listOfStudents := mapping.MapStudentListFromDomain(domainStudents)

	return listOfStudents, nil
}
