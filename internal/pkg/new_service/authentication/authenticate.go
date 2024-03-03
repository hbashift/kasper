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
