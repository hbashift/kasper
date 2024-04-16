package student_handler

import (
	"net/http"

	"uir_draft/internal/handlers/student_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpdateProgressiveness
//
//	@Summary		Обновление прогресса написания диссертации
//
//	@Description	Обновление прогресса написания диссертации
//
//	@Tags			Student.Dissertation
//	@Accept			json
//	@Param			input body request_models.UpdateProgressivenessRequest true "Данные"
//	@Success		200
//	@Param			token	path		string	true	"Токен пользователя"
//	@Failure		400		{string}	string	"Неверный формат данных"
//	@Failure		401		{string}	string	"Токен протух"
//	@Failure		204		{string}	string	"Нет записей в БД"
//	@Failure		500		{string}	string	"Ошибка на стороне сервера"
//	@Router			/students/dissertation/progress/percent/{token} [post]
func (h *StudentHandler) UpdateProgressiveness(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.UpdateProgressivenessRequest{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = h.student.UpdateStudentsProgressiveness(ctx, user.KasperID, reqBody.Progress); err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}