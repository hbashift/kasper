package mapping

import (
	"uir_draft/internal/generated/kasper/uir_draft/public/model"

	"github.com/google/uuid"
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
		Volume:     &work.Volume,
	}

	if work.WorkID != nil {
		res.WorkID = *work.WorkID
	}

	return res
}

func MapScientificWorkFromDomain(work *model.ScientificWork) *ScientificWork {
	res := &ScientificWork{
		Semester:   work.Semester,
		Name:       work.Name,
		State:      work.Name,
		Impact:     work.Impact,
		OutputData: work.OutputData,
		CoAuthors:  work.CoAuthors,
		WorkType:   work.WorkType,
		WorkID:     &work.WorkID,
	}

	if work.Volume != nil {
		res.Volume = *work.Volume
	}

	return res
}

func MapWorkIDsToDomain(ids *DeleteWorkIDs) ([]*uuid.UUID, error) {
	var UUIDs []*uuid.UUID

	for _, id := range ids.IDs {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		UUIDs = append(UUIDs, &uid)
	}

	return UUIDs, nil
}
