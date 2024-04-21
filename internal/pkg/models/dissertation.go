package models

import (
	"strings"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type SemesterProgressRequest struct {
	// Тип прогресса написания диссертации
	ProgressType string `json:"progress_type,omitempty" enums:"intro,ch. 1,ch. 2,ch. 3,ch. 4,ch. 5,ch. 6,end,literature,abstract"`
	// Первый семестр
	First bool `json:"first,omitempty"`
	// Второй семестр
	Second bool `json:"second,omitempty"`
	// Третий семестр
	Third bool `json:"third,omitempty"`
	// Четвертый семестр
	Forth bool `json:"forth,omitempty"`
	// Пятый семестр
	Fifth bool `json:"fifth,omitempty"`
	// Шестой семестр
	Sixth bool `json:"sixth,omitempty"`
	// Седьмой семестр
	Seventh bool `json:"seventh,omitempty"`
	// Восьмой семестр
	Eighth bool `json:"eighth,omitempty"`
}

//func (s *SemesterProgressRequest) SetDomainData(progress model.SemesterProgress) {
//	s.ProgressType = progress.ProgressType.String()
//	s.First = progress.First
//	s.Second = progress.Second
//	s.Third = progress.Third
//	s.Forth = progress.Forth
//	s.Fifth = progress.Fifth
//	s.Sixth = progress.Sixth
//	s.Seventh = progress.Seventh
//	s.Eighth = progress.Eighth
//}

func (s *SemesterProgressRequest) ToDomain() (model.SemesterProgress, error) {
	var progressType model.ProgressType
	err := progressType.Scan(strings.TrimSpace(s.ProgressType))
	if err != nil {
		return model.SemesterProgress{}, errors.Wrap(err, "SemesterProgressRequest.ToDomain(): wrong progress_type")
	}

	return model.SemesterProgress{
		ProgressType: progressType,
		First:        s.First,
		Second:       s.Second,
		Third:        s.Third,
		Forth:        s.Forth,
		Fifth:        s.Fifth,
		Sixth:        s.Sixth,
		Seventh:      s.Seventh,
		Eighth:       s.Eighth,
	}, nil
}

type DissertationsRequest struct {
	Semester int32 `json:"semester,omitempty"`
}

func (d *DissertationsRequest) ToDomain() model.Dissertations {
	return model.Dissertations{
		Semester: d.Semester,
	}
}

type DissertationTitlesRequest struct {
	Title    string `json:"title,omitempty"`
	Semester int32  `json:"semester,omitempty"`
}

//func (d *DissertationTitlesRequest) SetDomainData(titles model.DissertationTitles) {
//	d.Title = titles.Title
//	d.Semester = titles.Semester
//}

func (d *DissertationTitlesRequest) ToDomain() model.DissertationTitles {
	return model.DissertationTitles{
		Title:    d.Title,
		Semester: d.Semester,
	}
}

type FeedbackRequest struct {
	DissertationID uuid.UUID `json:"dissertation_id,omitempty"`
	Feedback       *string   `json:"feedback,omitempty"`
	Semester       int32     `json:"semester,omitempty"`
}

//func (f *FeedbackRequest) SetDomainData(feedback model.Feedback) {
//	f.DissertationID = feedback.DissertationID
//	f.Feedback = feedback.Feedback
//	f.Semester = feedback.Semester
//}

func (f *FeedbackRequest) ToDomain() model.Feedback {
	return model.Feedback{
		DissertationID: f.DissertationID,
		Feedback:       f.Feedback,
		Semester:       f.Semester,
	}
}

type DissertationPageResponse struct {
	// Информация о студенте
	StudentStatus Student `json:"student_status"`
	// Прогресс написания диссертации
	SemesterProgress []SemesterProgressResponse `json:"semester_progress,omitempty" format:"array"`
	// Статусы всех диссертаций (файлов)
	DissertationsStatuses []DissertationsResponse `json:"dissertations_statuses,omitempty" format:"array"`
	// Названия диссертаций
	DissertationTitles []DissertationTitlesResponse `json:"dissertation_titles,omitempty" format:"array"`
	// Массив обратной связи по каждой из диссертаций (за каждый семестр)
	Feedback []FeedbackResponse `json:"feedback,omitempty" format:"array"`
	// Список научных руководителей
	Supervisors []SupervisorFull `json:"supervisors,omitempty" format:"array"`
	// Сопроводительные комментарии студента научному руководителю за каждый семестр
	StudentsComments []StudentComment `json:"students_comments" format:"array"`
}

type StudentComment struct {
	CommentaryID uuid.UUID `json:"commentary_id,omitempty"`
	StudentID    uuid.UUID `json:"student_id,omitempty"`
	Semester     int32     `json:"semester,omitempty"`
	Commentary   *string   `json:"commentary,omitempty"`
	CommentedAt  time.Time `json:"commented_at"`
}

func (s *StudentComment) SetDomainData(data model.StudentsCommentary) {
	s.CommentaryID = data.CommentaryID
	s.StudentID = data.StudentID
	s.Semester = data.Semester
	s.Commentary = data.Commentary
	s.CommentedAt = data.CommentedAt
}

type SemesterProgressResponse struct {
	// ID прогресса написания для одного типа прогресса
	ProgressID uuid.UUID `json:"progress_id,omitempty" format:"uuid"`
	// ID студента
	StudentID uuid.UUID `json:"student_id,omitempty" format:"uuid"`
	// Тип прогресса
	ProgressType string `json:"progress_type,omitempty" enums:"intro,ch. 1,ch. 2,ch. 3,ch. 4,ch. 5,ch. 6,end,literature,abstract"`
	First        bool   `json:"first,omitempty"`
	Second       bool   `json:"second,omitempty"`
	Third        bool   `json:"third,omitempty"`
	Forth        bool   `json:"forth,omitempty"`
	Fifth        bool   `json:"fifth,omitempty"`
	Sixth        bool   `json:"sixth,omitempty"`
	Seventh      bool   `json:"seventh,omitempty"`
	Eighth       bool   `json:"eighth,omitempty"`
	// Дата последнего обновления
	UpdatedAt time.Time `json:"updated_at,omitempty" format:"date-time"`
	// Статус проверки и подтверждения
	Status string `json:"status,omitempty" enums:"todo,approved,on review,in progress,empty,failed"`
	// Дата принятия научным руководителем
	AcceptedAt *time.Time `json:"accepted_at,omitempty" format:"date-time"`
}

func (s *SemesterProgressResponse) SetDomainData(progress model.SemesterProgress) {
	s.ProgressID = progress.ProgressID
	s.StudentID = progress.StudentID
	s.ProgressType = progress.ProgressType.String()
	s.First = progress.First
	s.Second = progress.Second
	s.Third = progress.Third
	s.Forth = progress.Forth
	s.Fifth = progress.Fifth
	s.Sixth = progress.Sixth
	s.Seventh = progress.Seventh
	s.Eighth = progress.Eighth
	s.UpdatedAt = progress.UpdatedAt
	s.Status = progress.Status.String()
	s.AcceptedAt = progress.AcceptedAt
}

//func (s *SemesterProgressResponse) ToDomain() model.SemesterProgress {
//	return model.SemesterProgress{
//		ProgressID:   s.ProgressID,
//		StudentID:    s.StudentID,
//		ProgressType: model.ProgressType(s.ProgressType),
//		First:        s.First,
//		Second:       s.Second,
//		Third:        s.Third,
//		Forth:        s.Forth,
//		Fifth:        s.Fifth,
//		Sixth:        s.Sixth,
//		Seventh:      s.Seventh,
//		Eighth:       s.Eighth,
//		UpdatedAt:    s.UpdatedAt,
//		Status:       model.ApprovalStatus(s.Status),
//		AcceptedAt:   s.AcceptedAt,
//	}
//}

type DissertationsResponse struct {
	// ID диссертации
	DissertationID uuid.UUID `json:"dissertation_id,omitempty" format:"uuid"`
	// ID студента
	StudentID uuid.UUID `json:"student_id,omitempty" format:"uuid"`
	// Статус проверки и подтверждения
	Status string `json:"status,omitempty" enums:"todo,approved,on review,in progress,empty,failed"`
	// Дата создания
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Дата последнего обновления
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Семестр
	Semester int32 `json:"semester,omitempty"`
}

func (d *DissertationsResponse) SetDomainData(dissertations model.Dissertations) {
	d.DissertationID = dissertations.DissertationID
	d.StudentID = dissertations.StudentID
	d.Status = dissertations.Status.String()
	d.CreatedAt = dissertations.CreatedAt
	d.UpdatedAt = dissertations.UpdatedAt
	d.Semester = dissertations.Semester
}

//func (d *DissertationsResponse) ToDomain() model.Dissertations {
//	return model.Dissertations{
//		DissertationID: d.DissertationID,
//		StudentID:      d.StudentID,
//		Status:         model.ApprovalStatus(d.Status),
//		CreatedAt:      d.CreatedAt,
//		UpdatedAt:      d.UpdatedAt,
//		Semester:       d.Semester,
//	}
//}

type DissertationTitlesResponse struct {
	// ID названия диссертации
	TitleID uuid.UUID `json:"title_id,omitempty" format:"uuid"`
	// ID студента
	StudentID uuid.UUID `json:"student_id,omitempty" format:"uuid"`
	// Название
	Title string `json:"title,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Статус проверки и подтверждения
	Status string `json:"status,omitempty" enums:"todo,approved,on review,in progress,empty,failed"`
	// Дата принятия
	AcceptedAt *time.Time `json:"accepted_at,omitempty" format:"date-time"`
	// Семестр
	Semester int32 `json:"semester,omitempty"`
	// Объект исследования
	ResearchObject string `json:"research_object"`
	// Предмет исследования
	ResearchSubject string `json:"research_subject"`
}

func (d *DissertationTitlesResponse) SetDomainData(titles model.DissertationTitles) {
	d.TitleID = titles.TitleID
	d.StudentID = titles.StudentID
	d.Title = titles.Title
	d.CreatedAt = titles.CreatedAt
	d.Status = titles.Status.String()
	d.AcceptedAt = titles.AcceptedAt
	d.Semester = titles.Semester
	d.ResearchObject = titles.ResearchObject
	d.ResearchSubject = titles.ResearchSubject
}

//func (d *DissertationTitlesResponse) ToDomain() model.DissertationTitles {
//	return model.DissertationTitles{
//		TitleID:    d.TitleID,
//		StudentID:  d.StudentID,
//		Title:      d.Title,
//		CreatedAt:  d.CreatedAt,
//		Status:     model.ApprovalStatus(d.Status),
//		AcceptedAt: d.AcceptedAt,
//		Semester:   d.Semester,
//	}
//}

type FeedbackResponse struct {
	// ID обратной связи
	FeedbackID uuid.UUID `json:"feedback_id,omitempty" format:"uuid"`
	// ID студента
	StudentID uuid.UUID `json:"student_id,omitempty" format:"uuid"`
	// ID диссертации к которой привязана обратная связь
	DissertationID uuid.UUID `json:"dissertation_id,omitempty" format:"uuid"`
	// Текст обратной связи
	Feedback *string `json:"feedback,omitempty"`
	// Семестр
	Semester int32 `json:"semester,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Дата последнего обновления
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
}

func (f *FeedbackResponse) SetDomainData(feedback model.Feedback) {
	f.FeedbackID = feedback.FeedbackID
	f.StudentID = feedback.StudentID
	f.DissertationID = feedback.DissertationID
	f.Feedback = feedback.Feedback
	f.Semester = feedback.Semester
	f.CreatedAt = feedback.CreatedAt
	f.UpdatedAt = feedback.UpdatedAt
}

//func (f *FeedbackResponse) ToDomain() model.Feedback {
//	return model.Feedback{
//		FeedbackID:     f.FeedbackID,
//		StudentID:      f.StudentID,
//		DissertationID: f.DissertationID,
//		Feedback:       f.Feedback,
//		Semester:       f.Semester,
//		CreatedAt:      f.CreatedAt,
//		UpdatedAt:      f.UpdatedAt,
//	}
//}

type UploadDissertation struct {
	Semester Semester `form:"semester"`
}

type Semester struct {
	SemesterNumber int32 `json:"semester"`
}
