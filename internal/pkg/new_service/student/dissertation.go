package student

import (
	"context"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) GetDissertationPage(ctx context.Context, studentID uuid.UUID) (models.DissertationPageResponse, error) {
	page := models.DissertationPageResponse{}

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		semesterProgress, err := s.dissertationRepo.GetSemesterProgressTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		disTitles, err := s.dissertationRepo.GetDissertationTitlesTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		dissertationsStatuses, err := s.dissertationRepo.GetDissertationsTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		feedback, err := s.dissertationRepo.GetFeedbackTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		page = models.MapDissertationPageFromDomain(semesterProgress, dissertationsStatuses, disTitles, feedback)

		return nil
	})
	if err != nil {
		return models.DissertationPageResponse{}, errors.Wrap(err, "GetDissertationPage()")
	}

	return page, nil
}

func (s *Service) UpsertSemesterProgress(ctx context.Context, studentID uuid.UUID, progress []models.SemesterProgressRequest) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		dProgress, err := models.MapSemesterProgressToDomain(progress, model.ApprovalStatus_InProgress, studentID)
		if err != nil {
			return err
		}

		err = s.dissertationRepo.UpsertSemesterProgressTx(ctx, tx, dProgress)
		if err != nil {
			return err
		}

		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, studentID)

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertSemesterProgress()")
	}

	return nil
}

func (s *Service) UpsertDissertationInfo(ctx context.Context, studentID uuid.UUID, semester int32) error {
	// TODO сделать загрузку файла на диск

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		dissertationStatus := model.Dissertations{
			DissertationID: uuid.New(),
			StudentID:      studentID,
			Status:         model.ApprovalStatus_InProgress,
			CreatedAt:      lo.ToPtr(time.Now()),
			UpdatedAt:      lo.ToPtr(time.Now()),
			Semester:       semester,
		}

		err = s.dissertationRepo.UpsertDissertationTx(ctx, tx, dissertationStatus)
		err = s.studRepo.SetStudentStatusTx(ctx, tx, model.ApprovalStatus_InProgress, student.StudentID)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "InsertDissertationInfo()")
	}

	return nil
}

func (s *Service) UpsertDissertationTitle(ctx context.Context, studentID uuid.UUID, title string) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		err = s.dissertationRepo.InsertDissertationTitleTx(ctx, tx, model.DissertationTitles{
			TitleID:    uuid.New(),
			StudentID:  student.StudentID,
			Title:      title,
			CreatedAt:  time.Now(),
			Status:     model.ApprovalStatus_InProgress,
			AcceptedAt: nil,
			Semester:   student.ActualSemester,
		})

		return err
	})
	if err != nil {
		return errors.Wrap(err, "UpsertDissertationTitle()")
	}

	return nil
}

// GetDissertationData - for downloading dissertation file
func (s *Service) GetDissertationData(ctx context.Context, studentID uuid.UUID, semester int32) (model.Dissertations, error) {
	var dissertation model.Dissertations

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		data, err := s.dissertationRepo.GetActualDissertationData(ctx, tx, studentID, semester)
		if err != nil {
			return err
		}

		dissertation = data
		return nil
	})
	if err != nil {
		return model.Dissertations{}, errors.Wrap(err, "GetDissertationData()")
	}

	return dissertation, nil
}

func (s *Service) getUserInfo(ctx context.Context, tx pgx.Tx, token string) (model.Users, error) {
	userID, err := s.tokenRepo.GetUserIDByTokenTx(ctx, tx, token)
	if err != nil {
		return model.Users{}, errors.Wrap(err, "getting user_id by token")
	}

	user, err := s.userRepo.GetUserTx(ctx, tx, userID)
	if err != nil {
		return model.Users{}, errors.Wrap(err, "getting user info")
	}

	return user, nil
}