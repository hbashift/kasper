package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
)

type StudentRepository interface {
	GetDissertation(ctx context.Context, clientID uuid.UUID) (*models.DissertationPage, error)
}

type TokenRepository interface {
	// TODO сделать миддлварю
	Authenticate(ctx context.Context, token string) (*model.AuthorizationToken, error)
}

type Service struct {
	studRepo  StudentRepository
	tokenRepo TokenRepository
}

func NewService(studRepo StudentRepository, tokenRepo TokenRepository) *Service {
	return &Service{studRepo: studRepo, tokenRepo: tokenRepo}
}
