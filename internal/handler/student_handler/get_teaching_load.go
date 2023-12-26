package student_handler

import (
	"net/http"
	"sort"

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

	sort.Slice(load.Array, func(i, j int) bool {
		return load.Array[i].Semester < load.Array[j].Semester
	})

	ctx.JSON(http.StatusOK, load)
}
