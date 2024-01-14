package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"uir_draft/internal/pkg/helpers"
)

type GetSupervisorsStudents struct {
	SupervisorID uuid.UUID `json:"teacherID"`
}

func (h *administratorHandler) GetSupervisorsStudents(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := GetSupervisorsStudents{}
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	list, err := h.service.GetSupervisorsStudents(ctx, token.String(), reqBody.SupervisorID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}
