package models

import (
	"github.com/google/uuid"
)

type ChangeSupervisor struct {
	StudentID    uuid.UUID `json:"student_id" binding:"required"`
	SupervisorID uuid.UUID `json:"supervisor_id" binding:"required"`
}

type SetStudentsFlags struct {
	StudentID      uuid.UUID `json:"student_id,omitempty"`
	StudyingStatus string    `json:"studying_status,omitempty" enums:"academic,graduated,studying,expelled"`
	CanEdit        bool      `json:"can_edit"`
}

type UsersCredentials struct {
	Email    string
	Password string
}

type SupervisorStatus struct {
	SupervisorID uuid.UUID `json:"supervisor_id,omitempty"`
	Archived     bool      `json:"archived,omitempty"`
}
