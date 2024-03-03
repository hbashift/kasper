package request_models

import (
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
)

type DeleteIDs struct {
	IDs      []uuid.UUID `json:"ids,omitempty"`
	Semester int32       `json:"semester,omitempty"`
}

type UpsertAdditionalLoadRequest struct {
	// Дополнительные нагрузки, которые нужно добавить/изменить. Если запрос на добавление, то
	// ID дополнительной нагрузки (load_id) не заполнять (делать null), в противном случае - иначе
	Loads []models.AdditionalLoad `json:"loads,omitempty"`
	// ID совокупности всех нагрузок за семестр
	TLoadID uuid.UUID `json:"t_load_id,omitempty" format:"uuid"`
	// Семестр
	Semester int32 `json:"semester,omitempty"`
}

type UpsertClassroomLoadRequest struct {
	// Аудиторные нагрузки, которые нужно добавить/изменить. Если запрос на добавление, то
	// ID нагрузки (load_id) не заполнять (делать null), в противном случае - иначе
	Loads []models.ClassroomLoad `json:"loads,omitempty"`
	// ID совокупности всех нагрузок за семестр
	TLoadID uuid.UUID `json:"t_load_id,omitempty" format:"uuid"`
	// Семестр
	Semester int32 `json:"semester,omitempty"`
}

type UpsertConferencesRequest struct {
	// Научные конференции, которые нужно добавить/изменить. Если запрос на добавление, то
	// ID конференции (conference_id) не заполнять (делать null), в противном случае - иначе
	Conferences []models.Conference `json:"conferences,omitempty"`
	// ID совокупности всех научных работ за семестр
	WorkID uuid.UUID `json:"work_id,omitempty" format:"uuid"`
	// Семестр
	Semester int32 `json:"semester,omitempty"`
}

type UpsertIndividualLoadRequest struct {
	// Индивидуальные нагрузки, которые нужно добавить/изменить. Если запрос на добавление, то
	// ID индивидуальной нагрузки (load_id) не заполнять (делать null), в противном случае - иначе
	Loads    []models.IndividualStudentsLoad `json:"loads,omitempty"`
	TLoadID  uuid.UUID                       `json:"t_load_id,omitempty" format:"uuid"`
	Semester int32                           `json:"semester,omitempty"`
}

type UpsertPublicationsRequest struct {
	// Научные публикации, которые нужно добавить/изменить. Если запрос на добавление, то
	// ID публикации (publication_id) не заполнять (делать null), в противном случае - иначе
	Publications []models.Publication `json:"publications,omitempty"`
	// ID совокупности всех научных работ за семестр
	WorkID   uuid.UUID `json:"work_id,omitempty" format:"uuid"`
	Semester int32     `json:"semester,omitempty"`
}

type UpsertResearchProjectsRequest struct {
	// Научные проекты, которые нужно добавить/изменить. Если запрос на добавление, то
	// ID проекта (project_id) не заполнять (делать null), в противном случае - иначе
	Projects []models.ResearchProject `json:"projects,omitempty"`
	// ID совокупности всех научных работ за семестр
	WorkID   uuid.UUID `json:"work_id,omitempty" format:"uuid"`
	Semester int32     `json:"semester,omitempty"`
}

type UpsertProgressRequest struct {
	Progresses []models.SemesterProgressRequest `json:"progresses"`
}
