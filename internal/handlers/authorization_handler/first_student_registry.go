package authorization_handler

import (
	"net/http"

	"uir_draft/internal/handlers/authorization_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

// FirstStudentRegistry
//
//	@Summary		Первичная регистрация студента
//
//	@Description	Первичная регистрация студента
//
//	@Tags			Authorization
//	@Accept			json
//
//	@Produce		json
//
//	@Success		200
//	@Param			token	path		string								true	"Токен пользователя"
//	@Param			input	body		request_models.FirstStudentRegistry	true	"Данные"
//	@Failure		400		{string}	string								"Неверный формат данных"
//	@Failure		401		{string}	string								"Токен протух"
//	@Failure		204		{string}	string								"Нет записей в БД"
//	@Failure		500		{string}	string								"Ошибка на стороне сервера"
//	@Router			/authorize/registration/student/{token} [post]
func (h *AuthorizationHandler) FirstStudentRegistry(ctx *gin.Context) {
	user, err := h.authenticateStudent(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.FirstStudentRegistry{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = h.student.InitStudent(ctx, *user, reqBody); err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}
