package supervisor_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *supervisorHandler) GetListOfStudents(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	list, err := h.service.GetStudentsList(ctx, token.String())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}
