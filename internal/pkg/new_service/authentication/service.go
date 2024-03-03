package authentication

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	TokenRepository interface {
		GetUserIDByTokenTx(ctx context.Context, tx pgx.Tx, token string) (uuid.UUID, error)
		GetByTokenNumberTx(ctx context.Context, tx pgx.Tx, token string) (model.AuthorizationToken, error)
	}

	UsersRepository interface {
		GetUserTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (model.Users, error)
	}
)

type Service struct {
	tokenRepo TokenRepository
	userRepo  UsersRepository
	db        *pgxpool.Pool
}
