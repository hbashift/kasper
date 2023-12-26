package student_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (h *studentHandler) InsertTeachingLoad(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.TeachingLoad{}
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newTeachingLoad, err := h.service.UpsertTeachingLoad(ctx, token.String(), &reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, newTeachingLoad)
}
