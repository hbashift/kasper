package mapping

import "github.com/google/uuid"

type SetAcademicLeave struct {
	StudentID       uuid.UUID `json:"studentID,omitempty"`
	IsAcademicLeave bool      `json:"isAcademicLeave,omitempty"`
}

type ChangeSupervisor struct {
	StudentID    uuid.UUID `json:"studentID,omitempty"`
	SupervisorID uuid.UUID `json:"supervisorID,omitempty"`
}
