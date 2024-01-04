package authorization

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/authorization/mapping"
)

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
		Registered: client.Registered,
	}, true, nil
}
