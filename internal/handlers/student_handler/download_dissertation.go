package student_handler

import (
	"fmt"
	"net/http"
	"os"

	"uir_draft/internal/handlers/student_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

// DownloadDissertation
//
//	@Summary		Скачивание файла диссертации
//
//	@Description	Скачивание файла диссертации
//
//	@Tags			Student
//	@Accept			json
//
//	@Produce		json
//
//	@Success		200		"Файл"
//	@Param			token	path		string										true	"Токен пользователя"
//	@Param			input	body		request_models.DownloadDissertationRequest	true	"Данные"
//	@Failure		400		{string}	string										"Неверный формат данных"
//	@Failure		401		{string}	string										"Токен протух"
//	@Failure		204		{string}	string										"Нет записей в БД"
//	@Failure		500		{string}	string										"Ошибка на стороне сервера"
//	@Router			/students/dissertation/file/{token} [put]
func (h *StudentHandler) DownloadDissertation(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.DownloadDissertationRequest{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	dissertation, err := h.dissertation.GetDissertationData(ctx, user.KasperID, reqBody.Semester)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	if dissertation.FileName == nil {
		ctx.Status(http.StatusNoContent)
		return
	}

	dst := fmt.Sprintf("./dissertations/%s/semester%d/%s",
		dissertation.StudentID.String(), dissertation.Semester, lo.FromPtr(dissertation.FileName))

	_, err = os.Stat(dst)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Header("Content-Disposition", lo.FromPtr(dissertation.FileName))
	ctx.File(dst)
}
