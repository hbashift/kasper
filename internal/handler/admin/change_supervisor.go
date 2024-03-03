package admin

import (
	"net/http"

	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/service/admin/mapping"

	"github.com/gin-gonic/gin"
)

func (h *administratorHandler) ChangeSupervisor(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.ChangeSupervisorArray{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = h.service.ChangeSupervisor(ctx, token.String(), &reqBody); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
