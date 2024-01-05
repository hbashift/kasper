package student

import (
	"context"
	"time"

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
	startDate, err := time.Parse(time.DateOnly, info.StartDate)
	if err != nil {
		return err
	}

	student := model.Students{
		ClientID:          session.ClientID,
		StudentID:         session.KasperID,
		FullName:          info.FullName,
		Department:        info.Department,
		EnrollmentOrder:   info.EnrollmentOrder,
		Specialization:    lo.ToPtr(info.Specialization),
		ActualSemester:    info.ActualSemester,
		StartDate:         &startDate,
		AcademicLeave:     false,
		DissertationTitle: "unknown",
		Feedback:          nil,
		GroupNumber:       info.GroupNumber,
		SupervisorID:      uuid.Nil,
		NumberOfYears:     info.NumberOfYears,
	}
	supID := uuid.Nil

	if info.SupervisorID != nil {
		student.SupervisorID = *info.SupervisorID
		supID = *info.SupervisorID
	}

	err = s.studRepo.InsertStudentCommonInfo(ctx, s.db, student)
	if err != nil {
		return err
	}

	err = s.cliRepo.SetRegistered(ctx, s.db, session.ClientID)
	if err != nil {
		return errors.Wrap(err, "[Student]")
	}

	err = s.studSupRepo.SetStudentSupervisor(ctx, s.db, model.StudentSupervisor{
		StudentID:    session.KasperID,
		SupervisorID: supID,
	})

	// TODO сделать инициализацию програсса в семестре
	var models []model.SemesterProgress
	progressTypes := []model.ProgressType{
		model.ProgressType_Intro,
		model.ProgressType_Main,
		model.ProgressType_Ch1,
		model.ProgressType_Ch2,
		model.ProgressType_Ch3,
		model.ProgressType_Ch4,
		model.ProgressType_Ch5,
		model.ProgressType_Ch6,
		model.ProgressType_End,
		model.ProgressType_Literature,
		model.ProgressType_Abstract,
	}

	for _, progressType := range progressTypes {
		progress := model.SemesterProgress{
			SemesterProgressID: uuid.New(),
			StudentID:          session.KasperID,
			First:              false,
			Second:             false,
			Third:              false,
			Forth:              false,
			Fifth:              false,
			Sixth:              false,
			Seventh:            false,
			Eighth:             false,
			ProgressName:       progressType,
			LastUpdated:        lo.ToPtr(time.Now()),
			ClientID:           session.ClientID,
		}

		models = append(models, progress)
	}

	err = s.semesterRepo.InitSemesterProgress(ctx, s.db, models)
	if err != nil {
		return err
	}

	return nil
}
