package supervisor_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

func (h *supervisorHandler) SetStatus(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.SetStatus{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.service.SetStatus(ctx, token.String(), &reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
