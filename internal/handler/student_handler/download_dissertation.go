package student_handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"uir_draft/internal/pkg/helpers"
	"uir_draft/internal/pkg/service/student/mapping"
)

func (h *studentHandler) DownloadDissertation(ctx *gin.Context) {
	token, err := helpers.GetUUID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reqBody := mapping.DownloadDissertation{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dissertation, err := h.service.DownloadDissertation(ctx, token.String(), &reqBody)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dst := fmt.Sprintf("./dissertations/%s/semester%d/%s",
		dissertation.StudentID.String(), dissertation.Semester, dissertation.Name)

	_, err = os.Stat(dst)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.File(dst)
}
