package student_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (h *studentHandler) DeleteTeachingLoad(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.IDs{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("[Delete] Request body: %+v", reqBody)

	if len(reqBody.IDs) == 0 {
		ctx.JSON(http.StatusOK, "no ids provided")
		return
	}

	newWorks, err := h.service.DeleteTeachingLoad(ctx, token.String(), &reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, newWorks)
}
