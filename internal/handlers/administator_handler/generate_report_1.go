package administator_handler

import (
	"net/http"

	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

// GenerateReportOne
//
//	@Summary		Сгенерировать отчет 1
//
//	@Description	Сгенерировать отчет 1
//
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Param			token	path		string	true	"Токен пользователя"
//	@Failure		400		{string}	string	"Неверный формат данных"
//	@Failure		401		{string}	string	"Токен протух"
//	@Failure		204		{string}	string	"Нет записей в БД"
//	@Failure		500		{string}	string	"Ошибка на стороне сервера"
//	@Router			/administrator/generate/report-one/{token} [get]
func (h *AdministratorHandler) GenerateReportOne(ctx *gin.Context) {
	_, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	groups, err := h.user.GenerateReportOne(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, groups)
}
