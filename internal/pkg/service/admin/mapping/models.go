package mapping

import (
	"time"

	"github.com/google/uuid"
)

type SetAcademicLeave struct {
	StudentID       uuid.UUID `json:"studentID,omitempty"`
	IsAcademicLeave bool      `json:"isAcademicLeave,omitempty"`
}

type ChangeSupervisor struct {
	StudentID    uuid.UUID `json:"studentID,omitempty"`
	SupervisorID uuid.UUID `json:"supervisorID,omitempty"`
}

type UpdateStudentsCommonInfo struct {
	StudentID       uuid.UUID `json:"studentID,omitempty"`
	EnrollmentOrder string    `json:"enrollmentOrder,omitempty"`
	StartDate       time.Time `json:"startDate"`
}

type StudentSupervisorPair struct {
	StudentName    string    `json:"studentFullName,omitempty" db:"student_name"`
	StudentID      uuid.UUID `json:"studentId,omitempty" db:"student_id"`
	SupervisorName string    `json:"teacherFullName,omitempty" db:"supervisor_name"`
	SupervisorID   uuid.UUID `json:"teacherId,omitempty" db:"supervisor_id"`
}

type SupervisorInfo struct {
	SupervisorName string    `json:"teacherFullName,omitempty"`
	SupervisorID   uuid.UUID `json:"teacherId,omitempty"`
}

type GetStudSupPairs struct {
	Pairs       []*StudentSupervisorPair `json:"pairs,omitempty"`
	Supervisors []*SupervisorInfo        `json:"supervisors,omitempty"`
}
