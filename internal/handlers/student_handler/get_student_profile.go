package student_handler

import (
	"net/http"
	"time"

	"uir_draft/internal/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetStudentProfileResponse struct {
	// ID студента
	StudentID uuid.UUID `db:"students.student_id" json:"student_id,omitempty" format:"uuid"`
	// Полное имя
	FullName string `db:"students.full_name" json:"full_name,omitempty"`
	// Актуальный семестр
	ActualSemester int32 `db:"students.actual_semester" json:"actual_semester,omitempty"`
	// Количество лет обучения
	Years int32 `db:"students.years" json:"years,omitempty"`
	// Дата начала обучения
	StartDate time.Time `db:"students.start_date" json:"start_date"`
	// Статус обучения
	StudyingStatus string `db:"students.studying_status" json:"studying_status,omitempty" enums:"academic,graduated,studying,expelled"`
	// Статус проверки и подтверждения
	Status string `db:"students.status" json:"status,omitempty" enums:"todo,approved,on review,in progress,empty,failed"`
	// Специализация
	Specialization string `db:"specializations.title" json:"specialization,omitempty"`
	// Название группы
	GroupName string `db:"groups.group_name" json:"group_name,omitempty"`
	// Флаг о возможности редактировать всю информацию
	CanEdit bool `db:"students.can_edit" json:"can_edit,omitempty"`
	// Процент выполнения диссертации
	Progressiveness int32  `db:"students.progressiveness" json:"progress"`
	Email           string `json:"email"`
}

// GetStudentProfile
//
//	@Summary		Получение списка всех групп
//
//	@Description	Получение списка всех групп
//
//	@Tags			Student
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	GetStudentProfileResponse	"Данные"
//	@Param			token	path		string					true	"Токен пользователя"
//	@Failure		400		{string}	string					"Неверный формат данных"
//	@Failure		401		{string}	string					"Токен протух"
//	@Failure		204		{string}	string					"Нет записей в БД"
//	@Failure		500		{string}	string					"Ошибка на стороне сервера"
//	@Router			/student/profile/{token} [get]
func (h *StudentHandler) GetStudentProfile(ctx *gin.Context) {
	user, err := h.authenticate(ctx)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	student, err := h.student.GetStudentStatus(ctx, user.KasperID)
	if err != nil {
		ctx.AbortWithError(models.MapErrorToCode(err), err)
		return
	}

	resp := GetStudentProfileResponse{
		StudentID:       student.StudentID,
		FullName:        student.FullName,
		ActualSemester:  student.ActualSemester,
		Years:           student.Years,
		StartDate:       student.StartDate,
		StudyingStatus:  student.StudyingStatus,
		Status:          student.Status,
		Specialization:  student.Specialization,
		GroupName:       student.GroupName,
		CanEdit:         student.CanEdit,
		Progressiveness: student.Progressiveness,
		Email:           user.Email,
	}

	ctx.JSON(http.StatusOK, resp)
}
