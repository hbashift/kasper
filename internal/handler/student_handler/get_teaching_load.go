package student_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/service/student/mapping"
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

	arr := &mapping.SemesterTeachingLoad{}

	for _, el := range load.Array {
		switch {
		case el.Semester == int32(1):
			arr.Semester1 = append(arr.Semester1, el)
		case el.Semester == int32(2):
			arr.Semester2 = append(arr.Semester2, el)
		case el.Semester == int32(3):
			arr.Semester3 = append(arr.Semester3, el)
		case el.Semester == int32(4):
			arr.Semester4 = append(arr.Semester4, el)
		}
	}

	ctx.JSON(http.StatusOK, arr)
}
