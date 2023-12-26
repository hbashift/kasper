package supervisor_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *supervisorHandler) GetStudentsDissertationPage(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stringId := ctx.Param("studentID")

	studentID, err := uuid.Parse(stringId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page, err := h.service.GetDissertationPage(ctx, token.String(), studentID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, page)
}
