package supervisor_handler

import (
	"net/http"

	"uir_draft/internal/handlers/supervisor_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

// UpsertSupervisorMark
//
//	@Summary		Проставление оценки научником
//	@Tags			NEW
//	@Description	Проставление оценки научником
//	@Success		200
//	@Param			token	path	string	true	"Токен пользователя"
//	@Accept			json
//
//	@Produce		json
//
//	@Param			input	body		request_models.UpsertSupervisorMarkRequest	true	"Запрос"
//
//	@Failure		401		{string}	string										"Токен протух"
//	@Failure		204		{string}	string										"Нет записей в БД"
//	@Failure		500		{string}	string										"Ошибка на стороне сервера"
//	@Router			/supervisors/student/marks/{token} [post]
func (h *SupervisorHandler) UpsertSupervisorMark(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.UpsertSupervisorMarkRequest{}
	if err = ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	err = h.supervisor.UpsertSupervisorMark(ctx, reqBody.StudentID, user.KasperID, reqBody.Semester, reqBody.Mark)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusOK)
}
