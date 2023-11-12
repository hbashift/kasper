package student_handler

import (
	"net/http"

	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
)

func (h *studentHandler) InsertScientificWorks(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Warningf(ctx, "%s", ctx.Request.Body)

	reqBody := []*mapping.ScientificWork{}
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.service.InsertScientificWorks(ctx, token.String(), reqBody); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
