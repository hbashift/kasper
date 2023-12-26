package student_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *studentHandler) GetDissertationIDs(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ids, err := h.service.GetDissertationIDs(ctx, token.String())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, ids)
}
