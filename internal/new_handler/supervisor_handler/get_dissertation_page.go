package supervisor_handler

import (
	"net/http"

	"uir_draft/internal/new_handler/supervisor_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h *SupervisorHandler) GetDissertationPage(ctx *gin.Context) {
	_, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.GetByStudentID{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.dissertation.GetDissertationPage(ctx, reqBody.StudentID)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}
