package authorization_handler

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/handlers/authorization_handler/request_models"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

type (
	Authenticator interface {
		Authorize(ctx context.Context, request models.AuthorizeRequest) (*models.AuthorizeResponse, bool, error)
		Authenticate(ctx context.Context, token, userType string) (*model.Users, error)
	}

	StudentService interface {
		InitStudent(ctx context.Context, user model.Users, req request_models.FirstStudentRegistry) error
	}
)

type AuthorizationHandler struct {
	authenticator Authenticator
	student       StudentService
}

func NewAuthorizationHandler(authenticator Authenticator, student StudentService) *AuthorizationHandler {
	return &AuthorizationHandler{authenticator: authenticator, student: student}
}

func (h *AuthorizationHandler) authenticateStudent(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.Authenticate(ctx, token, model.UserType_Student.String())
	if err != nil {
		return user, err
	}

	return user, nil
}
