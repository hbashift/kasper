package mapping

import (
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
)

func MapScientificWorkToDomain(work *ScientificWork, session *model.AuthorizationToken) *model.ScientificWork {
	res := &model.ScientificWork{
		StudentID:  session.KasperID,
		Semester:   work.Semester,
		Name:       work.Name,
		State:      work.Name,
		Impact:     work.Impact,
		OutputData: work.OutputData,
		CoAuthors:  work.CoAuthors,
		WorkType:   work.WorkType,
	}

	if work.WorkID != nil {
		res.WorkID = *work.WorkID
	}

	return res
}
