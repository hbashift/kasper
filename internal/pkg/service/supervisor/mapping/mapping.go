package mapping

import (
	"github.com/samber/lo"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
)

func MapStudentListFromDomain(domain []*model.Students) *ListOfStudents {
	list := &ListOfStudents{}

	for _, domainStudent := range domain {
		student := StudentCommonInfo{
			FullName:        domainStudent.FullName,
			Group:           lo.FromPtr(domainStudent.GroupNumber),
			Topic:           domainStudent.DissertationTitle,
			EnrollmentOrder: domainStudent.EnrollmentOrder,
			DateOfStatement: lo.FromPtr(domainStudent.StartDate),
			StudentID:       domainStudent.StudentID,
		}

		list.Array = append(list.Array, student)
	}

	return list
}
