package models

import (
	"time"

	"github.com/google/uuid"
)

type Publication struct {
	// ID Совокупности научных работ за семестр
	WorksID uuid.UUID `json:"works_id,omitempty" db:"publications.works_id" format:"uuid"`
	// ID научной публикации
	PublicationID *uuid.UUID `json:"publication_id,omitempty" db:"publications.publication_id" format:"uuid"`
	// Название научной публикации
	Name *string `json:"name,omitempty" db:"publications.name"`
	// Индекс РИНЦ
	Rinc *bool `json:"rinc" db:"publications.rinc"`
	// Индекс Scopus
	Scopus *bool `json:"scopus" db:"publications.scopus"`
	// Индекс WaC
	Wac *bool `json:"wac" db:"publications.wac"`
	// Индекс WoS
	Wos *bool `json:"wos" db:"publications.wos"`
	// Влияние
	Impact *float64 `json:"impact,omitempty" db:"publications.impact" format:"float"`
	// Статус прогресса публикации
	Status *string `json:"status,omitempty" db:"publications.status" enums:"to print,published,in progress"`
	// Выходные данные
	OutputData *string `json:"output_data,omitempty" db:"publications.output_data"`
	// Со-Авторы
	CoAuthors *string `json:"co_authors,omitempty" db:"publications.co_authors"`
	// Объем написанной работы
	Volume *int32 `json:"volume,omitempty" db:"publications.volume"`
}

type Conference struct {
	// ID Совокупности научных работ за семестр
	WorksID uuid.UUID `json:"works_id,omitempty" db:"conferences.works_id" format:"uuid"`
	// ID конференции
	ConferenceID *uuid.UUID `json:"conference_id,omitempty" db:"conferences.conference_id" format:"uuid"`
	// Статус прогресса научной конференции
	Status *string `json:"status,omitempty" db:"conferences.status" enums:"registered,performed"`
	// Индекс РИНЦ
	Rinc *bool `json:"rinc" db:"conferences.rinc"`
	// Индекс Scopus
	Scopus *bool `json:"scopus" db:"conferences.scopus"`
	// Индекс WaC
	Wac *bool `json:"wac" db:"conferences.wac"`
	// Индекс WoS
	Wos *bool `json:"wos" db:"conferences.wos"`
	// Название конференции
	ConferenceName *string `json:"conference_name,omitempty" db:"conferences.conference_name"`
	// Название доклада
	ReportName *string `json:"report_name,omitempty" db:"conferences.report_name"`
	// Место проведения
	Location *string `json:"location,omitempty" db:"conferences.location"`
	// Дата доклада
	ReportedAt *time.Time `json:"reported_at" db:"conferences.reported_at" format:"date-time"`
}

type ResearchProject struct {
	// ID Совокупности научных работ за семестр
	WorksID uuid.UUID `json:"works_id,omitempty" db:"research_projects.works_id" format:"uuid"`
	// ID научного проекта
	ProjectID *uuid.UUID `json:"project_id,omitempty" db:"research_projects.project_id" format:"uuid"`
	// Название проекта
	ProjectName *string `json:"project_name,omitempty" db:"research_projects.project_name"`
	// Дата начала проекта
	StartAt *time.Time `json:"start_at" db:"research_projects.start_at" format:"date-time"`
	// Дата окончание
	EndAt *time.Time `json:"end_at" db:"research_projects.end_at" format:"date-time"`
	// Дополнительная информация
	AddInfo *string `json:"add_info,omitempty" db:"research_projects.add_info"`
	// Грантодатель
	Grantee *string `json:"grantee,omitempty" db:"research_projects.grantee"`
}

type ScientificWork struct {
	// Семестр, за который присылаются научные работы
	Semester int `json:"semester,omitempty" db:"scientific_works.semester"`
	// ID студента
	StudentID uuid.UUID `json:"student_id,omitempty" db:"scientific_works.student_id" format:"uuid"`
	// Статус проверки и подтверждения
	ApprovalStatus string `json:"works_status,omitempty" db:"scientific_works.approval_status" enums:"todo,approved,on review,in progress,empty,failed"`
	// Дата последнего обновления
	UpdatedAt time.Time `json:"updated_at" db:"scientific_works.updated_at" format:"date-time"`
	// Дата принятия научным руководителем
	AcceptedAt *time.Time `json:"accepted_at,omitempty" db:"scientific_works.accepted_at" format:"date-time"`
	// Объект, описывающий научную публикацию
	Publications []Publication `json:"publications"`
	// Объект, описывающий научную конференцию
	Conferences []Conference `json:"conferences"`
	// Объект, описывающий научно-исследовательский проект
	ResearchProjects []ResearchProject `json:"research_projects"`
}
