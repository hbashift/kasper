package authentication

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
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
		return nil, models.ErrWrongUserType
	}

	return &user, nil
}
