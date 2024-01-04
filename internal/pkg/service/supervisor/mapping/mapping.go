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
			StudentID:       domainStudent.StudentID,
		}

		if domainStudent.StartDate != nil {
			student.DateOfStatement = domainStudent.StartDate.Format("02/01/2006")
		}

		list.Array = append(list.Array, student)
	}

	return list
}

func MapDissertationStatus(status string) model.DissertationStatus {
	switch {
	case status == "Не сдано":
		return model.DissertationStatus_Failed
	case status == "Принято":
		return model.DissertationStatus_Passed
	case status == "На доработку":
		return model.DissertationStatus_Todo
	default:
		return model.DissertationStatus_Empty
	}
}
