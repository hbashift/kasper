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
		p := &mapping.ScientificWork{
			WorkID:     work.WorkID,
			Semester:   work.Semester,
			Name:       work.Name,
			State:      work.State,
			Impact:     work.Impact,
			OutputData: work.OutputData,
			CoAuthors:  work.CoAuthors,
			WorkType:   work.WorkType,
			Volume:     work.Volume,
		}
		scientificWorks = append(scientificWorks, p)
	}

	newWorks, err := h.service.InsertScientificWorks(ctx, token.String(), scientificWorks)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, newWorks)
}
