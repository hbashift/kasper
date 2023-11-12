package student_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandler) GetDissertation(ctx *gin.Context) {
	id, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dissertation, err := h.service.GetDissertationPage(ctx, id.String())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, dissertation)
}
