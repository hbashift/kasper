package student_handler

import (
	"context"
	"net/http"
	"time"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type StudentService interface {
	GetDissertationPage(ctx context.Context, token string) (*models.DissertationPage, error)
	UpsertSemesterPlan(ctx context.Context, progress []*model.SemesterProgress) error
}

type studentHandler struct {
	service StudentService
}

type SemesterProgress struct {
	SemesterProgressID int32      `json:"semesterProgressID,omitempty"`
	StudentID          uuid.UUID  `json:"studentID,omitempty"`
	First              bool       `json:"first,omitempty"`
	Second             bool       `json:"second,omitempty"`
	Third              bool       `json:"third,omitempty"`
	Forth              bool       `json:"forth,omitempty"`
	Fifth              *bool      `json:"fifth,omitempty"`
	Sixth              *bool      `json:"sixth,omitempty"`
	ProgressName       string     `json:"progressName,omitempty"`
	LastUpdated        *time.Time `json:"lastUpdated,omitempty"`
	ClientID           uuid.UUID  `json:"clientID,omitempty"`
}

type Progress struct {
	Progress []SemesterProgress `json:"progress,omitempty"`
}

func NewStudentHandler(service StudentService) *studentHandler {
	return &studentHandler{service: service}
}

func (h *studentHandler) UpsertSemesterProgress(ctx *gin.Context) {
	res := &Progress{}
	err := ctx.ShouldBind(res)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var domainProgress []*model.SemesterProgress

	for _, prog := range res.Progress {
		p := &model.SemesterProgress{
			SemesterProgressID: prog.SemesterProgressID,
			StudentID:          prog.StudentID,
			First:              prog.First,
			Second:             prog.Second,
			Third:              prog.Third,
			Forth:              prog.Forth,
			Fifth:              nil,
			Sixth:              nil,
			ProgressName:       model.ProgressType(prog.ProgressName),
			LastUpdated:        lo.ToPtr(time.Now()),
			ClientID:           prog.ClientID,
		}

		domainProgress = append(domainProgress, p)
	}

	if err := h.service.UpsertSemesterPlan(ctx, domainProgress); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "everything is ok")
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
