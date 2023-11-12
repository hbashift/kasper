package student_handler

import (
	"net/http"

	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/gin-gonic/gin"
)

func (h *studentHandler) UpsertSemesterProgress(ctx *gin.Context) {
	id, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res := &mapping.Progress{}
	err = ctx.ShouldBindJSON(res)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.service.UpsertSemesterPlan(ctx, id.String(), res); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "everything is ok")
}
