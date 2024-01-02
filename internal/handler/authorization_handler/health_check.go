package authorization_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/service/authorization"
)

func (h *authorizationHandler) HealthCheck(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	check, err := h.authorization.HealthCheck(ctx, token.String())
	if errors.Is(err, authorization.ErrNonValidToken) {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, check)
}
