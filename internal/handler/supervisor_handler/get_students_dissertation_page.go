package supervisor_handler

import (
	"log"
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

	body := []byte{}
	val, err := ctx.Request.Body.Read(body)
	log.Printf("KeyValue: %+v\n", val)

	log.Printf("Context: %+v", ctx)

	//stringId := ctx.Param("studentID")
	//
	//studentID, err := uuid.Parse(stringId)
	studentID := uuid.New()

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
