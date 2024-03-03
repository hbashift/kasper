package student_handler

import (
	"net/http"

	"uir_draft/internal/new_handler/student_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h *StudentHandler) UpsertIndividualLoad(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.UpsertIndividualLoadRequest{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.load.UpsertIndividualLoad(ctx, user.KasperID, reqBody.TLoadID, reqBody.Semester, reqBody.Loads)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}
