package administator_handler

import (
	"net/http"

	"uir_draft/internal/handlers/administator_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

type ReqStruct struct {
	Students []request_models.StudentsToNewSemester `json:"students"`
}

// StudentsToNewSemester
//
//	@Summary		Отправление студентов на следующий семестр
//
//	@Description	Отправление студентов на следующий семестр
//
//	@Tags			NEW 2
//	@Accept			json
//	@Param			input body request_models.DeleteByUUIDRequest true "Данные"
//	@Success		200
//	@Param			token	path		string	true	"Токен пользователя"
//	@Failure		400		{string}	string	"Неверный формат данных"
//	@Failure		401		{string}	string	"Токен протух"
//	@Failure		204		{string}	string	"Нет записей в БД"
//	@Failure		500		{string}	string	"Ошибка на стороне сервера"
//	@Router			/administrator/student/new-semester/{token} [post]
func (h *AdministratorHandler) StudentsToNewSemester(ctx *gin.Context) {
	_, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := ReqStruct{}
	if err = ctx.ShouldBind(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = h.user.StudentsToNewSemester(ctx, reqBody.Students); err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	ctx.Status(http.StatusOK)
}
