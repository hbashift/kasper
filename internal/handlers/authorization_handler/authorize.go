package authorization_handler

import (
	"net/http"

	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *AuthorizationHandler) Authorize(ctx *gin.Context) {
	reqBody := models.AuthorizeRequest{}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, authorized, err := h.authenticator.Authorize(ctx, reqBody)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}
	if !authorized {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("wrong password or email"))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
