package admin

import (
	"context"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/admin/mapping"
)

func (s *Service) SetAcademicLeave(ctx context.Context, token string, leave *mapping.SetAcademicLeave) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "authentication error")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	if err = s.studRepo.SetAcademicLeave(ctx, s.db, leave.StudentID, leave.IsAcademicLeave); err != nil {
		return errors.Wrap(err, "[Admin]")
	}

	return nil
}
