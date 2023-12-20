package authorization

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/authorization/mapping"
)

type ClientUserRepository interface {
	GetClient(ctx context.Context, tx *pgxpool.Pool, email string) (*model.ClientUser, error)
}

type TokenRepository interface {
	InsertToken(ctx context.Context, tx *pgxpool.Pool, token *model.AuthorizationToken) error
}

type Service struct {
	clientRepo ClientUserRepository
	tokenRepo  TokenRepository
	db         *pgxpool.Pool
}

func NewService(clientRepo ClientUserRepository, tokenRepo TokenRepository, db *pgxpool.Pool) *Service {
	return &Service{clientRepo: clientRepo, tokenRepo: tokenRepo, db: db}
}

func (s *Service) Authorize(ctx context.Context, info *mapping.AuthorizeInfo) (*mapping.Authorization, bool, error) {
	client, err := s.clientRepo.GetClient(ctx, s.db, info.Email)
	if err != nil {
		return nil, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(info.Password))
	if err != nil {
		return nil, false, nil
	}

	token := &model.AuthorizationToken{
		TokenID:     uuid.New(),
		ClientID:    client.ClientID,
		KasperID:    client.KasperID,
		TokenStatus: model.TokenStatus_Active,
		TokenNumber: uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   lo.ToPtr(time.Now()),
	}

	err = s.tokenRepo.InsertToken(ctx, s.db, token)
	if err != nil {
		return nil, false, err
	}

	return &mapping.Authorization{
		ClientType: client.ClientType.String(),
		Token:      uuid.MustParse(token.TokenNumber),
	}, true, nil
}
