package authorization

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
)

var ErrNonValidToken = errors.New("token is expired")

type ClientUserRepository interface {
	GetClient(ctx context.Context, tx *pgxpool.Pool, email string) (*model.ClientUser, error)
	ChangePassword(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID, password string) error
}

type TokenRepository interface {
	InsertToken(ctx context.Context, tx *pgxpool.Pool, token *model.AuthorizationToken) error
	Authenticate(ctx context.Context, token string, tx *pgxpool.Pool) (*model.AuthorizationToken, error)
	GetUserTypeByToken(ctx context.Context, tx *pgxpool.Pool, token string) (*model.UserType, error)
}

type Service struct {
	clientRepo ClientUserRepository
	tokenRepo  TokenRepository
	db         *pgxpool.Pool
}

func NewService(clientRepo ClientUserRepository, tokenRepo TokenRepository, db *pgxpool.Pool) *Service {
	return &Service{clientRepo: clientRepo, tokenRepo: tokenRepo, db: db}
}
