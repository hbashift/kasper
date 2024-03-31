package models

import (
	"time"

	"github.com/google/uuid"
)

type ClassroomLoad struct {
	// ID аудиторной нагрузки
	LoadID *uuid.UUID `json:"load_id,omitempty" db:"classroom_load.load_id" format:"uuid"`
	// Кол-во часов
	Hours *int32 `json:"hours,omitempty" db:"classroom_load.hours"`
	// Тип аудиторной нагрузки
	LoadType *string `json:"load_type,omitempty" db:"classroom_load.load_type" enums:"practice,lectures,laboratory,exam"`
	// Основное учитель
	MainTeacher *string `json:"main_teacher,omitempty" db:"classroom_load.main_teacher"`
	// Название группы
	GroupName *string `json:"group_name,omitempty" db:"classroom_load.group_name"`
	// Название предмета
	SubjectName *string `json:"subject_name,omitempty" db:"classroom_load.subject_name"`
}

type IndividualStudentsLoad struct {
	// ID индивидуальной работы со студентами
	LoadID *uuid.UUID `json:"load_id,omitempty" db:"individual_students_load.load_id" format:"uuid"`
	// Количество студентов
	StudentsAmount *int32 `json:"students_amount,omitempty" db:"individual_students_load.students_amount"`
	// Тип индивидуальной работы
	LoadType *string `json:"load_type" db:"individual_students_load.load_type" enums:"project practice,bachelor,masters"`
	// Комментарий
	Comment *string `json:"comment,omitempty" db:"individual_students_load.comment"`
}

type AdditionalLoad struct {
	// ID дополнительной нагрузки
	LoadID *uuid.UUID `json:"load_id,omitempty" db:"additional_load.load_id" format:"uuid"`
	// Название нагрузки
	Name *string `json:"name,omitempty" db:"additional_load.name"`
	// Объем
	Volume *string `json:"volume,omitempty" db:"additional_load.volume"`
	// Комментарий
	Comment *string `json:"comment,omitempty" db:"additional_load.comment"`
}

type TeachingLoad struct {
	// ID совокупности нагрузок за семестр
	LoadsID uuid.UUID `json:"loads_id,omitempty" db:"teaching_load.loads_id"`
	// ID студента
	StudentID uuid.UUID `json:"student_id,omitempty" db:"teaching_load.student_id"`
	// Семестр
	Semester int `json:"semester,omitempty" db:"teaching_load.semester"`
	// Статус проверки и подтверждения
	ApprovalStatus string `json:"approval_status,omitempty" db:"teaching_load.approval_status" enums:"todo,approved,on review,in progress,empty,failed"`
	// Дата последнего обновления
	UpdatedAt time.Time `json:"updated_at" db:"teaching_load.updated_at"`
	// Дата принятия научным руководителем
	AcceptedAt *time.Time `json:"accepted_at,omitempty" db:"teaching_load.accepted_at"`
	// Объект, описывающий аудиторную нагрузку
	ClassroomLoad ClassroomLoad `json:"classroom_load"`
	// Объект, описывающий индивидуальную работу со студентами
	IndividualStudentsLoad IndividualStudentsLoad `json:"individual_students_load"`
	// Объект, описывающий дополнительную нагрузку
	AdditionalLoad AdditionalLoad `json:"additional_load"`
}
