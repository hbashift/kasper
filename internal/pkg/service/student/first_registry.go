package student

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (s *Service) FirstRegistry(ctx context.Context, token string, info *mapping.FirstRegistry) error {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return errors.Wrap(ErrNonValidToken, "[Student]")
	}

	student := model.Students{
		ClientID:          session.ClientID,
		StudentID:         session.KasperID,
		FullName:          info.FullName,
		Department:        info.Department,
		EnrollmentOrder:   info.EnrollmentOrder,
		Specialization:    lo.ToPtr(model.StudentSpecialization(info.Specialization)),
		ActualSemester:    info.ActualSemester,
		StartDate:         info.StartDate,
		AcademicLeave:     false,
		DissertationTitle: info.DissertationTitle,
		Feedback:          nil,
		GroupNumber:       info.GroupNumber,
		SupervisorID:      uuid.Nil,
	}

	if info.SupervisorID != nil {
		student.SupervisorID = *info.SupervisorID
	}

	err = s.studRepo.InsertStudentCommonInfo(ctx, s.db, student)
	if err != nil {
		return err
	}

	return nil
}
