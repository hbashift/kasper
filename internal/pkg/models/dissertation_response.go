package models

import (
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
)

type StudentDissertationPlan struct {
	Name    string `db:"name" json:"name,omitempty"`
	First   bool   `db:"first" json:"id1,omitempty"`
	Second  bool   `db:"second" json:"id2,omitempty"`
	Third   bool   `db:"third" json:"id3,omitempty"`
	Forth   bool   `db:"forth" json:"id4,omitempty"`
	Fifth   bool   `db:"fifth" json:"id5,omitempty"`
	Sixth   bool   `db:"sixth" json:"id6,omitempty"`
	Seventh bool   `db:"seventh" json:"id7,omitempty"`
	Eighth  bool   `db:"eighth" json:"id8,omitempty"`
}

type DissertationPageResponse struct {
	SemesterProgress      []SemesterProgressResponse   `json:"semester_progress,omitempty"`
	DissertationsStatuses []DissertationsResponse      `json:"dissertations_statuses,omitempty"`
	DissertationTitles    []DissertationTitlesResponse `json:"dissertation_titles,omitempty"`
	Feedback              []FeedbackResponse           `json:"feedback,omitempty"`
}

type SemesterProgressResponse struct {
	ProgressID   uuid.UUID            `json:"progress_id,omitempty"`
	StudentID    uuid.UUID            `json:"student_id,omitempty"`
	ProgressType model.ProgressType   `json:"progress_type,omitempty"`
	First        bool                 `json:"first,omitempty"`
	Second       bool                 `json:"second,omitempty"`
	Third        bool                 `json:"third,omitempty"`
	Forth        bool                 `json:"forth,omitempty"`
	Fifth        bool                 `json:"fifth,omitempty"`
	Sixth        bool                 `json:"sixth,omitempty"`
	Seventh      bool                 `json:"seventh,omitempty"`
	Eighth       bool                 `json:"eighth,omitempty"`
	UpdatedAt    time.Time            `json:"updated_at,omitempty"`
	Status       model.ApprovalStatus `json:"status,omitempty"`
	AcceptedAt   *time.Time           `json:"accepted_at,omitempty"`
}

func (s *SemesterProgressResponse) SetDomainData(progress model.SemesterProgress) {
	s.ProgressID = progress.ProgressID
	s.StudentID = progress.StudentID
	s.ProgressType = progress.ProgressType
	s.First = progress.First
	s.Second = progress.Second
	s.Third = progress.Third
	s.Forth = progress.Forth
	s.Fifth = progress.Fifth
	s.Sixth = progress.Sixth
	s.Seventh = progress.Seventh
	s.Eighth = progress.Eighth
	s.UpdatedAt = progress.UpdatedAt
	s.Status = progress.Status
	s.AcceptedAt = progress.AcceptedAt
}

func (s *SemesterProgressResponse) ToDomain() model.SemesterProgress {
	return model.SemesterProgress{
		ProgressID:   s.ProgressID,
		StudentID:    s.StudentID,
		ProgressType: s.ProgressType,
		First:        s.First,
		Second:       s.Second,
		Third:        s.Third,
		Forth:        s.Forth,
		Fifth:        s.Fifth,
		Sixth:        s.Sixth,
		Seventh:      s.Seventh,
		Eighth:       s.Eighth,
		UpdatedAt:    s.UpdatedAt,
		Status:       s.Status,
		AcceptedAt:   s.AcceptedAt,
	}
}

type DissertationsResponse struct {
	DissertationID uuid.UUID            `json:"dissertation_id,omitempty"`
	StudentID      uuid.UUID            `json:"student_id,omitempty"`
	Status         model.ApprovalStatus `json:"status,omitempty"`
	CreatedAt      *time.Time           `json:"created_at,omitempty"`
	UpdatedAt      *time.Time           `json:"updated_at,omitempty"`
	Semester       int32                `json:"semester,omitempty"`
}

func (d *DissertationsResponse) SetDomainData(dissertations model.Dissertations) {
	d.DissertationID = dissertations.DissertationID
	d.StudentID = dissertations.StudentID
	d.Status = dissertations.Status
	d.CreatedAt = dissertations.CreatedAt
	d.UpdatedAt = dissertations.UpdatedAt
	d.Semester = dissertations.Semester
}

func (d *DissertationsResponse) ToDomain() model.Dissertations {
	return model.Dissertations{
		DissertationID: d.DissertationID,
		StudentID:      d.StudentID,
		Status:         d.Status,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
		Semester:       d.Semester,
	}
}

type DissertationTitlesResponse struct {
	TitleID    uuid.UUID            `json:"title_id,omitempty"`
	StudentID  uuid.UUID            `json:"student_id,omitempty"`
	Title      string               `json:"title,omitempty"`
	CreatedAt  time.Time            `json:"created_at"`
	Status     model.ApprovalStatus `json:"status,omitempty"`
	AcceptedAt *time.Time           `json:"accepted_at,omitempty"`
	Semester   int32                `json:"semester,omitempty"`
}

func (d *DissertationTitlesResponse) SetDomainData(titles model.DissertationTitles) {
	d.TitleID = titles.TitleID
	d.StudentID = titles.StudentID
	d.Title = titles.Title
	d.CreatedAt = titles.CreatedAt
	d.Status = titles.Status
	d.AcceptedAt = titles.AcceptedAt
	d.Semester = titles.Semester
}

func (d *DissertationTitlesResponse) ToDomain() model.DissertationTitles {
	return model.DissertationTitles{
		TitleID:    d.TitleID,
		StudentID:  d.StudentID,
		Title:      d.Title,
		CreatedAt:  d.CreatedAt,
		Status:     d.Status,
		AcceptedAt: d.AcceptedAt,
		Semester:   d.Semester,
	}
}

type FeedbackResponse struct {
	FeedbackID     uuid.UUID `json:"feedback_id,omitempty"`
	StudentID      uuid.UUID `json:"student_id,omitempty"`
	DissertationID uuid.UUID `json:"dissertation_id,omitempty"`
	Feedback       *string   `json:"feedback,omitempty"`
	Semester       int32     `json:"semester,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
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

func (f *FeedbackResponse) ToDomain() model.Feedback {
	return model.Feedback{
		FeedbackID:     f.FeedbackID,
		StudentID:      f.StudentID,
		DissertationID: f.DissertationID,
		Feedback:       f.Feedback,
		Semester:       f.Semester,
		CreatedAt:      f.CreatedAt,
		UpdatedAt:      f.UpdatedAt,
	}
}

type UploadDissertation struct {
	Semester Semester `form:"semester"`
}

type Semester struct {
	SemesterNumber int32 `json:"semester"`
}
