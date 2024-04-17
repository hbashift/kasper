package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
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
	Progressiveness int32 `db:"students.progressiveness" json:"progress"`
}

type Supervisor struct {
	// ID научного руководителя
	SupervisorID uuid.UUID `db:"supervisor_id" json:"supervisor_id" format:"uuid"`
	// Полное имя руководителя
	FullName   string `db:"full_name" json:"full_name"`
	Faculty    string `db:"faculty" json:"faculty"`
	Department string `db:"department" json:"department"`
	Degree     string `db:"degree" json:"degree"`
}

type SupervisorFull struct {
	// ID научного руководителя
	SupervisorID uuid.UUID `db:"supervisor_id" json:"supervisor_id" format:"uuid"`
	// Полное имя руководителя
	FullName string `db:"full_name" json:"full_name"`
	// Дата начала
	StartAt time.Time `db:"start_at" json:"start_at" format:"date-time"`
	// Дата окончания (пустое, если руководитель актуальный)
	EndAt *time.Time `db:"end_at" json:"end_at,omitempty" format:"date-time"`
}

type StudentSupervisorPair struct {
	// Информация о студенте в паре
	Student Student `json:"student"`
	// Информация о научном руководителе в паре
	Supervisor Supervisor `json:"supervisor"`
}
