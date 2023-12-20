package authorization_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/service/authorization/mapping"
)

func (h *authorizationHandler) Authorize(ctx *gin.Context) {
	reqBody := mapping.AuthorizeInfo{}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info, authorized, err := h.authorization.Authorize(ctx, &reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !authorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, info)
}
