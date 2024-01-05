package student_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/helpers"
)

type Theme struct {
	Theme string `json:"theme"`
}

func (h *studentHandler) SetTheme(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := Theme{}
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = h.service.SetTheme(ctx, token.String(), reqBody.Theme); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
