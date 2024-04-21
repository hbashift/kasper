package authentication

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	TokenRepository interface {
		GetByTokenNumberTx(ctx context.Context, tx pgx.Tx, token string) (model.AuthorizationToken, error)
		InsertTokenTx(ctx context.Context, tx pgx.Tx, token *model.AuthorizationToken) error
	}

	UsersRepository interface {
		GetUserTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (model.Users, error)
		GetUserByEmailTx(ctx context.Context, tx pgx.Tx, email string) (model.Users, error)
		ChangeUsersPasswordTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID, password string) error
	}
)

type Service struct {
	tokenRepo TokenRepository
	userRepo  UsersRepository
	db        *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{
		tokenRepo: repository.NewTokenRepository(),
		userRepo:  repository.NewUsersRepository(),
		db:        db,
	}
}
