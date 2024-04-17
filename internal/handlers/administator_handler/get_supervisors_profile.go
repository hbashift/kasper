package administator_handler

import (
	"net/http"

	"uir_draft/internal/handlers/administator_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetSupervisorProfileResponse struct {
	// ID научного руководителя
	SupervisorID uuid.UUID `db:"supervisor_id" json:"supervisor_id" format:"uuid"`
	// Полное имя руководителя
	FullName   string  `db:"full_name" json:"full_name"`
	Faculty    *string `db:"faculty" json:"faculty"`
	Department *string `db:"department" json:"department"`
	Degree     *string `db:"degree" json:"degree"`
	Email      string  `json:"email"`
}

// GetSupervisorProfile
//
//	@Summary		Получение профиля научного руководителя
//	@Tags			Admin
//	@Description	Получение профиля научного руководителя
//	@Success		200		{object}	GetSupervisorProfileResponse		"Данные"
//
//	@Param			input	body		request_models.GetBySupervisorID	true	"Данные"
//	@Produce		json
//	@Param			token	path		string	true	"Токен пользователя"
//
//	@Failure		401		{string}	string	"Токен протух"
//	@Failure		204		{string}	string	"Нет записей в БД"
//	@Failure		500		{string}	string	"Ошибка на стороне сервера"
//	@Router			/administrator/supervisors/profile/{token} [put]
func (h *AdministratorHandler) GetSupervisorProfile(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	reqBody := request_models.GetBySupervisorID{}
	if err = ctx.ShouldBind(&reqBody); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	super, err := h.supervisor.GetSupervisorProfile(ctx, reqBody.SupervisorID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp := GetSupervisorProfileResponse{
		SupervisorID: super.SupervisorID,
		FullName:     super.FullName,
		Faculty:      super.Faculty,
		Department:   super.Department,
		Degree:       super.Degree,
		Email:        user.Email,
	}

	ctx.JSON(http.StatusOK, resp)
}
