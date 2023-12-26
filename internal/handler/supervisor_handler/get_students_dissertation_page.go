package supervisor_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Student struct {
	StudentID uuid.UUID `json:"StudentID"`
}

func (h *supervisorHandler) GetStudentsDissertationPage(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := Student{}
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	page, err := h.service.GetDissertationPage(ctx, token.String(), reqBody.StudentID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, page)
}
