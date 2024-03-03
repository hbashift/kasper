package student_handler

import (
	"net/http"

	"uir_draft/internal/new_handler/student_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h *StudentHandler) UpsertPublications(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.UpsertPublicationsRequest{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.scientific.UpsertPublications(ctx, user.KasperID, reqBody.WorkID, reqBody.Semester, reqBody.Publications)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}
