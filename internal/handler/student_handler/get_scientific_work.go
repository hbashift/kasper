package student_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandler) GetScientificWorks(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	works, err := h.service.GetScientificWorks(ctx, token.String())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, works)
}
