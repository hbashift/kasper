package student_handler

import (
	"fmt"
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

	log.Println(fmt.Sprintf("starting delete"))

	log.Printf(fmt.Sprintf("config value: %s", ctx.Request.FormValue("config")))

	reqBody := mapping.DeleteWorkIDs{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("Request body: %+v", reqBody)

	if err = h.service.DeleteScientificWork(ctx, token.String(), &reqBody); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
