package student_handler

import (
	"fmt"
	"log"
	"net/http"

	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/gin-gonic/gin"
)

func (h *studentHandler) InsertScientificWorks(ctx *gin.Context) {
	token, err := getUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.Works{}
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Println(fmt.Sprintf("%+v", reqBody.Works))

	var scientificWorks []*mapping.ScientificWork

	for _, work := range reqBody.Works {
		p := &work
		scientificWorks = append(scientificWorks, p)
	}

	if err := h.service.InsertScientificWorks(ctx, token.String(), scientificWorks); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, scientificWorks)
}
