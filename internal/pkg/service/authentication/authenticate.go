package authentication

import (
	"context"
	"strings"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Authenticate(ctx context.Context, token, userType string) (*model.Users, error) {
	var user model.Users

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		tokenModel, err := s.tokenRepo.GetByTokenNumberTx(ctx, tx, token)
		if err != nil {
			return errors.Wrap(err, "getting user_id by token")
		}

		if !tokenModel.IsActive {
			return models.ErrTokenExpired
		}

		user, err = s.userRepo.GetUserTx(ctx, tx, tokenModel.UserID)
		if err != nil {
			return errors.Wrap(err, "getting user info")
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "Authenticate()")
	}

	if user.UserType.String() != userType {
		if user.UserType == model.UserType_Admin && userType == model.UserType_Supervisor.String() {
			return &user, nil
		}
		return nil, models.ErrWrongUserType
	}

	return &user, nil
}

func (s *Service) Authorize(ctx context.Context, request models.AuthorizeRequest) (*models.AuthorizeResponse, bool, error) {
	var response models.AuthorizeResponse
	var isAuthorized bool

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		user, err := s.userRepo.GetUserByEmailTx(ctx, tx, strings.TrimSpace(request.Email))
		if err != nil {
			return errors.Wrap(err, "getting user info")
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			isAuthorized = false
			return nil
		}

		token := model.AuthorizationToken{
			UserID:      user.UserID,
			IsActive:    true,
			TokenNumber: uuid.New().String(),
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		}

		if err = s.tokenRepo.InsertTokenTx(ctx, tx, &token); err != nil {
			return err
		}

		response = models.AuthorizeResponse{
			UserType:   user.UserType.String(),
			Token:      token.TokenNumber,
			Registered: user.Registered,
		}

		return nil
	}); err != nil {
		return nil, false, errors.Wrap(err, "Authenticate()")
	}

	if !isAuthorized {
		return nil, false, nil
	}

	return &response, isAuthorized, nil
}

func (s *Service) FirstStudentRegistry(ctx context.Context, userID, studentID uuid.UUID, student model.Students) error {
	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		// TODO map values from registry page
		m := model.Students{
			StudentID:      studentID,
			UserID:         userID,
			FullName:       "",
			Department:     "",
			SpecID:         0,
			ActualSemester: 0,
			Years:          0,
			StartDate:      time.Time{},
			StudyingStatus: "",
			GroupID:        0,
			Status:         "",
			CanEdit:        false,
		}

		err := s.studRepo.InsertStudentTx(ctx, tx, m)
		if err != nil {
			return err
		}

		err = s.scienceRepo.InitScientificWorkStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		err = s.loadRepo.InitTeachingLoadsStatusTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		progresses := make([]model.SemesterProgress, 0, 10)
		progressTypes := []model.ProgressType{
			model.ProgressType_Intro,
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
				ProgressID:   uuid.New(),
				StudentID:    studentID,
				ProgressType: progressType,
				First:        false,
				Second:       false,
				Third:        false,
				Forth:        false,
				Fifth:        false,
				Sixth:        false,
				Seventh:      false,
				Eighth:       false,
				UpdatedAt:    time.Now(),
				Status:       model.ApprovalStatus_Empty,
				AcceptedAt:   nil,
			}

			progresses = append(progresses, progress)
		}

		err = s.dissertationRepo.UpsertSemesterProgressTx(ctx, tx, progresses)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "FirstStudentRegistry()")
	}

	return nil
}
