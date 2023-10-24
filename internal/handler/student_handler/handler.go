package student_handler

import (
	"context"
	"net/http"

	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentService interface {
	GetDissertationPage(ctx context.Context, token string) (*models.DissertationPage, error)
}

type studentHandler struct {
	service StudentService
}

func NewStudentHandler(service StudentService) *studentHandler {
	return &studentHandler{service: service}
}

func (h *studentHandler) GetDissertation(ctx *gin.Context) {
	id, err := getUUID(ctx)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dissertation, err := h.service.GetDissertationPage(ctx, id.String())

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, dissertation)
}

func getUUID(ctx *gin.Context) (*uuid.UUID, error) {
	stringId := ctx.Param("id")

	id, err := uuid.Parse(stringId)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

// localhost:8080/students/dissertation/:id
