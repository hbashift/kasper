package mapping

import (
	"errors"

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
	}

	if work.Volume > 0 {
		res.Volume = &work.Volume
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

func MapTeachingLoadFromDomain(domainLoads []*model.TeachingLoad) TeachingLoad {
	var loads []SingleLoad

	for _, domainLoad := range domainLoads {
		if domainLoad == nil {
			continue
		}

		loadType, check := TeachingLoadTypeMapFromDomain[domainLoad.LoadType]
		if !check {
			loadType = TeachingLoadType_UNKNOWN
		}

		load := SingleLoad{
			StudentID:      domainLoad.StudentID,
			Semester:       domainLoad.Semester,
			Hours:          domainLoad.Hours,
			AdditionalLoad: domainLoad.AdditionalLoad,
			LoadType:       loadType,
			MainTeacher:    domainLoad.MainTeacher,
			GroupName:      domainLoad.GroupName,
			SubjectName:    domainLoad.SubjectName,
		}

		loads = append(loads, load)
	}

	return TeachingLoad{Array: loads}
}

func MapTeachingLoadToDomain(loads *TeachingLoad, session *model.AuthorizationToken) ([]*model.TeachingLoad, error) {
	var domainLoads []*model.TeachingLoad

	for _, load := range loads.Array {
		loadType, check := TeachingLoadTypeMapToDomain[load.LoadType]
		if !check {
			return nil, errors.New("unknown teaching_load_type")
		}

		domainLoad := model.TeachingLoad{
			LoadID:         uuid.New(),
			StudentID:      session.KasperID,
			Semester:       load.Semester,
			Hours:          load.Hours,
			AdditionalLoad: load.AdditionalLoad,
			LoadType:       loadType,
			MainTeacher:    load.MainTeacher,
			GroupName:      load.GroupName,
			SubjectName:    load.SubjectName,
		}

		domainLoads = append(domainLoads, &domainLoad)
	}

	return domainLoads, nil
}
