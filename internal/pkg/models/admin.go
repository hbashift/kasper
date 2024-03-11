package models

import (
	"github.com/google/uuid"
)

type ChangeSupervisor struct {
	StudentID    uuid.UUID `json:"student_id" binding:"required"`
	SupervisorID uuid.UUID `json:"supervisor_id" binding:"required"`
}
