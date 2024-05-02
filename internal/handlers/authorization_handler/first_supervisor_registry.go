package authorization_handler

import (
	"log"
	"net/http"

	"uir_draft/internal/handlers/authorization_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

// FirstSupervisorRegistry
//
//	@Summary		Первичная регистрация научного руководителя
//
//	@Description	Первичная регистрация научного руководителя
//
//	@Tags			NEW
//	@Accept			json
//
//	@Produce		json
//
//	@Success		200
//	@Param			token	path		string									true	"Токен пользователя"
//	@Param			input	body		request_models.FirstSupervisorRegistry	true	"Данные"
//	@Failure		400		{string}	string									"Неверный формат данных"
//	@Failure		401		{string}	string									"Токен протух"
//	@Failure		204		{string}	string									"Нет записей в БД"
//	@Failure		500		{string}	string									"Ошибка на стороне сервера"
//	@Router			/authorize/registration/supervisor/{token} [post]
func (h *AuthorizationHandler) FirstSupervisorRegistry(ctx *gin.Context) {
	user, err := h.authenticateSupervisor(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.FirstSupervisorRegistry{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("first_student_registry request body: %v", reqBody)

	if err = h.supervisor.InitSupervisor(ctx, *user, reqBody); err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}
