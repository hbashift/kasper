package models

import (
	"github.com/google/uuid"
)

type ReportComments struct {
	DissertationComments []DissertationComment `json:"dissertation_comments"`
	DissertationPlans    []DissertationPlan    `json:"dissertation_plans"`
}

type DissertationComment struct {
	CommentaryID uuid.UUID `json:"commentary_id"`
	StudentID    uuid.UUID `json:"student_id"`
	Semester     int32     `json:"semester"`
	Commentary   *string   `json:"commentary,omitempty"`
}

type DissertationPlan struct {
	PlanID    uuid.UUID `json:"plan_id"`
	StudentID uuid.UUID `json:"student_id"`
	Semester  int32     `json:"semester"`
	PlanText  *string   `json:"plan_text,omitempty"`
}

type DissertationCommentRequest struct {
	Semester   int32   `json:"semester"`
	Commentary *string `json:"commentary,omitempty"`
}

type DissertationPlanRequest struct {
	Semester int32   `json:"semester"`
	PlanText *string `json:"plan_text,omitempty"`
}
