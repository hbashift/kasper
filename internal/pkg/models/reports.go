package models

import (
	"github.com/google/uuid"
)

type ReportComments struct {
	DissertationComments []DissertationComment `json:"dissertation_comments,omitempty"`
	DissertationPlans    []DissertationPlan    `json:"dissertation_plans,omitempty"`
}

type DissertationComment struct {
	CommentaryID uuid.UUID `json:"commentary_id,omitempty"`
	StudentID    uuid.UUID `json:"student_id,omitempty"`
	Semester     int32     `json:"semester,omitempty"`
	Commentary   *string   `json:"commentary,omitempty"`
}

type DissertationPlan struct {
	PlanID    uuid.UUID `json:"plan_id,omitempty"`
	StudentID uuid.UUID `json:"student_id,omitempty"`
	Semester  int32     `json:"semester,omitempty"`
	PlanText  *string   `json:"plan_text,omitempty"`
}

type DissertationCommentRequest struct {
	Semester   int32   `json:"semester,omitempty"`
	Commentary *string `json:"commentary,omitempty"`
}

type DissertationPlanRequest struct {
	Semester int32   `json:"semester,omitempty"`
	PlanText *string `json:"plan_text,omitempty"`
}
