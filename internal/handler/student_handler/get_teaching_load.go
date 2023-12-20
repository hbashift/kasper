package student_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandler) GetTeachingLoad(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	load, err := h.service.GetTeachingLoad(ctx, token.String())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, load)
}
