package administator_handler

import (
	"fmt"
	"os"

	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

// DownloadReportTwo
//
//	@Summary		Скачивание отчета 2
//
//	@Description	Скачивание отчета 2
//
//	@Tags			NEW 2
//	@Accept			json
//
//	@Produce		json
//
//	@Success		200		"Файл"
//	@Param			token	path		string	true	"Токен пользователя"
//	@Failure		400		{string}	string	"Неверный формат данных"
//	@Failure		401		{string}	string	"Токен протух"
//	@Failure		204		{string}	string	"Нет записей в БД"
//	@Failure		500		{string}	string	"Ошибка на стороне сервера"
//	@Router			/administrator/download/report-two/{token} [get]
func (h *AdministratorHandler) DownloadReportTwo(ctx *gin.Context) {
	_, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	dst := fmt.Sprint("./reports/report_2/output.xlsx")

	_, err = os.Stat(dst)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Header("Content-Disposition", "output2.xlsx")
	ctx.File(dst)
}
