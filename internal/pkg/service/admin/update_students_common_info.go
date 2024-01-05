package admin

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/admin/mapping"
)

func (s *Service) UpdateStudentsCommonInfo(ctx context.Context, token string, info *mapping.UpdateStudentsCommonInfo) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "authentication error")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return ErrNonValidToken
	}

	startDate, err := time.Parse("02.01.2006", info.StartDate)
	if err != nil {
		return err
	}
	student := model.Students{
		StudentID:       info.StudentID,
		EnrollmentOrder: info.EnrollmentOrder,
		StartDate:       &startDate,
	}

	if err = s.studRepo.UpdateStudentCommonInfo(ctx, s.db, student); err != nil {
		return errors.Wrap(err, "[Admin]")
	}

	return nil
}
