package mapping

import (
	"github.com/google/uuid"
)

type ListOfStudents struct {
	Array []StudentCommonInfo `json:"array,omitempty"`
}

type StudentCommonInfo struct {
	FullName        string    `json:"fullName,omitempty"`
	Group           string    `json:"group,omitempty"`
	Topic           string    `json:"topic,omitempty"`
	EnrollmentOrder string    `json:"numberOfOrderOfStatement,omitempty"`
	DateOfStatement string    `json:"dateOfStatement,omitempty"`
	StudentID       uuid.UUID `json:"studentID,omitempty"`
}

type DownloadDissertation struct {
	Semester  int32     `json:"semester,omitempty"`
	StudentID uuid.UUID `json:"studentID,omitempty"`
}

type SetStatus struct {
	Semester int32  `json:"semester,omitempty"`
	Status   string `json:"status"`
}
