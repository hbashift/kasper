package supervisor_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/service/supervisor/mapping"
)

func (h *supervisorHandler) GetScientificWorks(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.StudentID{}
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	works, err := h.service.GetScientificWorks(ctx, token.String(), reqBody.StudentID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, works)
}
