package models

import (
	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
)

type DissertationPageRequest struct {
	SemesterProgress      []SemesterProgressResponse
	DissertationsStatuses []DissertationsResponse
	DissertationTitles    []DissertationTitlesResponse
	Feedback              []FeedbackResponse
}

type SemesterProgressRequest struct {
	ProgressType model.ProgressType `json:"progress_type,omitempty"`
	First        bool               `json:"first,omitempty"`
	Second       bool               `json:"second,omitempty"`
	Third        bool               `json:"third,omitempty"`
	Forth        bool               `json:"forth,omitempty"`
	Fifth        bool               `json:"fifth,omitempty"`
	Sixth        bool               `json:"sixth,omitempty"`
	Seventh      bool               `json:"seventh,omitempty"`
	Eighth       bool               `json:"eighth,omitempty"`
}

func (s *SemesterProgressRequest) SetDomainData(progress model.SemesterProgress) {
	s.ProgressType = progress.ProgressType
	s.First = progress.First
	s.Second = progress.Second
	s.Third = progress.Third
	s.Forth = progress.Forth
	s.Fifth = progress.Fifth
	s.Sixth = progress.Sixth
	s.Seventh = progress.Seventh
	s.Eighth = progress.Eighth
}

func (s *SemesterProgressRequest) ToDomain() model.SemesterProgress {
	return model.SemesterProgress{
		ProgressType: s.ProgressType,
		First:        s.First,
		Second:       s.Second,
		Third:        s.Third,
		Forth:        s.Forth,
		Fifth:        s.Fifth,
		Sixth:        s.Sixth,
		Seventh:      s.Seventh,
		Eighth:       s.Eighth,
	}
}

type DissertationsRequest struct {
	Semester int32 `json:"semester,omitempty"`
}

func (d *DissertationsRequest) SetDomainData(dissertations model.Dissertations) {
	d.Semester = dissertations.Semester
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

func (d *DissertationTitlesRequest) SetDomainData(titles model.DissertationTitles) {
	d.Title = titles.Title
	d.Semester = titles.Semester
}

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

func (f *FeedbackRequest) SetDomainData(feedback model.Feedback) {
	f.DissertationID = feedback.DissertationID
	f.Feedback = feedback.Feedback
	f.Semester = feedback.Semester
}

func (f *FeedbackRequest) ToDomain() model.Feedback {
	return model.Feedback{
		DissertationID: f.DissertationID,
		Feedback:       f.Feedback,
		Semester:       f.Semester,
	}
}
