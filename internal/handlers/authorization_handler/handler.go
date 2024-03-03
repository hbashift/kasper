package authorization_handler

import (
	"context"

	"uir_draft/internal/pkg/models"
)

type (
	Authenticator interface {
		Authorize(ctx context.Context, request models.AuthorizeRequest) (*models.AuthorizeResponse, bool, error)
	}
)

type AuthorizationHandler struct {
	authenticator Authenticator
}
