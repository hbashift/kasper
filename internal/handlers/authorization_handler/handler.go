package authorization_handler

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/handlers/authorization_handler/request_models"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	Authenticator interface {
		Authorize(ctx context.Context, request models.AuthorizeRequest) (*models.AuthorizeResponse, bool, error)
		AuthenticateWithUserType(ctx context.Context, token, userType string) (*model.Users, error)
		Authenticate(ctx context.Context, token string) (*model.Users, error)
		TokenCheck(ctx context.Context, token string) (*model.Users, error)
		ChangePassword(ctx context.Context, userID uuid.UUID, request request_models.ChangePasswordRequest) error
		GetUserProfile(ctx context.Context, userID uuid.UUID) (model.Users, error)
	}

	StudentService interface {
		InitStudent(ctx context.Context, user model.Users, req request_models.FirstStudentRegistry) error
	}
)

type AuthorizationHandler struct {
	authenticator Authenticator
	student       StudentService
}

func NewHandler(authenticator Authenticator, student StudentService) *AuthorizationHandler {
	return &AuthorizationHandler{authenticator: authenticator, student: student}
}

func (h *AuthorizationHandler) authenticateStudent(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.AuthenticateWithUserType(ctx, token, model.UserType_Student.String())
	if err != nil {
		return user, err
	}

	return user, nil
}

func (h *AuthorizationHandler) authenticate(ctx *gin.Context) (*model.Users, error) {
	token := helpers.GetToken(ctx)

	user, err := h.authenticator.Authenticate(ctx, token)
	if err != nil {
		return user, err
	}

	return user, nil
}
