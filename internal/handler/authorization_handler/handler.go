package authorization_handler

import (
	"context"

	"uir_draft/internal/pkg/service/authorization/mapping"
)

type AuthorizationService interface {
	Authorize(ctx context.Context, info *mapping.AuthorizeInfo) (*mapping.Authorization, bool, error)
}

type authorizationHandler struct {
	authorization AuthorizationService
}

func NewAuthorizationHandler(authorization AuthorizationService) *authorizationHandler {
	return &authorizationHandler{authorization: authorization}
}
