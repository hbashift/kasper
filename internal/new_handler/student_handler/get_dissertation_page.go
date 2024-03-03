package student_handler

import (
	"net/http"

	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

func (h *StudentHandler) GetDissertationPage(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	dissertationPage, err := h.dissertation.GetDissertationPage(ctx, user.KasperID)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, dissertationPage)
}
