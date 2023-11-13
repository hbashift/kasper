package student_handler

import (
	"log"
	"net/http"

	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/gin-gonic/gin"
)

func (h *studentHandler) DeleteScientificWork(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.DeleteWorkIDs{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("[Delete] Request body: %+v", reqBody)

	newWorks, err := h.service.DeleteScientificWork(ctx, token.String(), &reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var works []mapping.ScientificWork

	for _, m := range newWorks {
		works = append(works, *m)
	}

	log.Printf("[Delete] Response body: %+v", works)

	ctx.JSON(http.StatusOK, newWorks)
}
