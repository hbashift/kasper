package models

import (
	"strings"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func MapDissertationPageFromDomain(dProgresses []model.SemesterProgress, dDissertations []model.Dissertations, dTitles []model.DissertationTitles, dFeedbacks []model.Feedback) DissertationPageResponse {
	progresses := make([]SemesterProgressResponse, 0, len(dProgresses))
	for _, dProgress := range dProgresses {
		progress := SemesterProgressResponse{}
		progress.SetDomainData(dProgress)

		progresses = append(progresses, progress)
	}

	dissertations := make([]DissertationsResponse, 0, len(dDissertations))
	for _, dDissertation := range dDissertations {
		dissertation := DissertationsResponse{}
		dissertation.SetDomainData(dDissertation)

		dissertations = append(dissertations, dissertation)
	}

	titles := make([]DissertationTitlesResponse, 0, len(dTitles))
	for _, dTitle := range dTitles {
		title := DissertationTitlesResponse{}
		title.SetDomainData(dTitle)

		titles = append(titles, title)
	}

	feedbacks := make([]FeedbackResponse, 0, len(dFeedbacks))
	for _, dFeedback := range dFeedbacks {
		feedback := FeedbackResponse{}
		feedback.SetDomainData(dFeedback)

		feedbacks = append(feedbacks, feedback)
	}

	return DissertationPageResponse{
		SemesterProgress:      progresses,
		DissertationsStatuses: dissertations,
		DissertationTitles:    titles,
		Feedback:              feedbacks,
	}
}

func MapSemesterProgressToDomain(progresses []SemesterProgressRequest, status model.ApprovalStatus, studentID uuid.UUID) ([]model.SemesterProgress, error) {
	updatedAt := time.Now()

	var acceptedAt *time.Time
	acceptedAt = nil
	if status == model.ApprovalStatus_Approved {
		acceptedAt = lo.ToPtr(time.Now())
	}

	dProgresses := make([]model.SemesterProgress, 0, len(progresses))

	for _, progress := range progresses {
		dProgress, err := progress.ToDomain()
		if err != nil {
			return nil, errors.Wrap(err, "MapSemesterProgressToDomain()")
		}

		dProgress.ProgressID = uuid.New()
		dProgress.StudentID = studentID
		dProgress.UpdatedAt = updatedAt
		dProgress.Status = status
		dProgress.AcceptedAt = acceptedAt

		dProgresses = append(dProgresses, dProgress)
	}

	return dProgresses, nil
}

func MapPublicationsToDomain(publications []Publication) (dPublicationsInsert, dPublicationsUpdate []model.Publications, err error) {
	for _, publication := range publications {
		var status model.PublicationStatus = ""
		if err := status.Scan(strings.TrimSpace(lo.FromPtr(publication.Status))); err != nil {
			return nil, nil, ErrInvalidEnumValue
		}

		dPublication := model.Publications{
			WorksID:    publication.WorksID,
			Name:       lo.FromPtr(publication.Name),
			Scopus:     lo.FromPtr(publication.Scopus),
			Rinc:       lo.FromPtr(publication.Rinc),
			Wac:        lo.FromPtr(publication.Wac),
			Wos:        lo.FromPtr(publication.Wos),
			Impact:     lo.FromPtr(publication.Impact),
			Status:     status,
			OutputData: publication.OutputData,
			CoAuthors:  publication.CoAuthors,
			Volume:     publication.Volume,
		}

		if lo.FromPtr(publication.PublicationID) == uuid.Nil {
			dPublication.PublicationID = uuid.New()
			dPublicationsInsert = append(dPublicationsInsert, dPublication)
		} else {
			dPublication.PublicationID = lo.FromPtr(publication.PublicationID)
			dPublicationsUpdate = append(dPublicationsUpdate, dPublication)
		}
	}

	return dPublicationsInsert, dPublicationsUpdate, nil
}

func MapConferencesToDomain(conferences []Conference) (dConferencesInsert, dConferencesUpdate []model.Conferences, err error) {
	for _, conf := range conferences {
		var status model.ConferenceStatus = ""
		if err := status.Scan(strings.TrimSpace(lo.FromPtr(conf.Status))); err != nil {
			return nil, nil, ErrInvalidEnumValue
		}

		dConf := model.Conferences{
			WorksID:        conf.WorksID,
			Status:         status,
			Scopus:         lo.FromPtr(conf.Scopus),
			Rinc:           lo.FromPtr(conf.Rinc),
			Wac:            lo.FromPtr(conf.Wac),
			Wos:            lo.FromPtr(conf.Wos),
			ConferenceName: lo.FromPtr(conf.ConferenceName),
			ReportName:     lo.FromPtr(conf.ReportName),
			Location:       lo.FromPtr(conf.Location),
			ReportedAt:     lo.FromPtr(conf.ReportedAt),
		}

		if lo.FromPtr(conf.ConferenceID) == uuid.Nil {
			dConf.ConferenceID = uuid.New()
			dConferencesInsert = append(dConferencesInsert, dConf)
		} else {
			dConf.ConferenceID = lo.FromPtr(conf.ConferenceID)
			dConferencesUpdate = append(dConferencesUpdate, dConf)
		}
	}

	return dConferencesInsert, dConferencesUpdate, nil
}

func MapResearchProjectToDomain(projects []ResearchProject) (dResearchInsert, dResearchUpdate []model.ResearchProjects) {
	for _, proj := range projects {
		dProj := model.ResearchProjects{
			WorksID:     proj.WorksID,
			ProjectName: lo.FromPtr(proj.ProjectName),
			StartAt:     lo.FromPtr(proj.StartAt),
			EndAt:       lo.FromPtr(proj.EndAt),
			AddInfo:     proj.AddInfo,
			Grantee:     proj.Grantee,
		}

		if lo.FromPtr(proj.ProjectID) == uuid.Nil {
			dProj.ProjectID = uuid.New()
			dResearchInsert = append(dResearchInsert, dProj)
		} else {
			dProj.ProjectID = lo.FromPtr(proj.ProjectID)
			dResearchUpdate = append(dResearchUpdate, dProj)
		}
	}

	return dResearchInsert, dResearchUpdate
}

func MapClassroomLoadToDomain(loads []ClassroomLoad) (dLoadInsert, dLoadUpdate []model.ClassroomLoad, err error) {
	for _, load := range loads {
		var loadType model.ClassroomLoadType
		if err := loadType.Scan(strings.TrimSpace(lo.FromPtr(load.LoadType))); err != nil {
			return nil, nil, ErrInvalidEnumValue
		}

		dLoad := model.ClassroomLoad{
			TLoadID:     load.TLoadID,
			Hours:       lo.FromPtr(load.Hours),
			LoadType:    loadType,
			MainTeacher: lo.FromPtr(load.MainTeacher),
			GroupName:   lo.FromPtr(load.GroupName),
			SubjectName: lo.FromPtr(load.SubjectName),
		}
		if lo.FromPtr(load.LoadID) == uuid.Nil {
			dLoad.LoadID = uuid.New()
			dLoadInsert = append(dLoadInsert, dLoad)
		} else {
			dLoad.LoadID = lo.FromPtr(load.LoadID)
			dLoadUpdate = append(dLoadUpdate, dLoad)
		}
	}

	return dLoadInsert, dLoadUpdate, nil
}

func MapIndividualWorkToDomain(loads []IndividualStudentsLoad) (dLoadInsert, dLoadUpdate []model.IndividualStudentsLoad, err error) {
	for _, load := range loads {
		dLoad := model.IndividualStudentsLoad{
			TLoadID:        load.TLoadID,
			StudentsAmount: lo.FromPtr(load.StudentsAmount),
			Comment:        load.Comment,
		}

		var loadType model.ApprovalStatus
		if err = loadType.Scan(load.LoadType); err != nil {
			return nil, nil, errors.Wrap(err, "MapIndividualWorkToDomain()")
		}

		if lo.FromPtr(load.LoadID) == uuid.Nil {
			dLoad.LoadID = uuid.New()
			dLoadInsert = append(dLoadInsert, dLoad)
		} else {
			dLoad.LoadID = lo.FromPtr(load.LoadID)
			dLoadUpdate = append(dLoadUpdate, dLoad)
		}
	}

	return dLoadInsert, dLoadUpdate, nil
}

func MapAdditionalLoadToDomain(loads []AdditionalLoad) (dLoadInsert, dLoadUpdate []model.AdditionalLoad) {
	for _, load := range loads {
		dLoad := model.AdditionalLoad{
			TLoadID: load.TLoadID,
			Name:    lo.FromPtr(load.Name),
			Volume:  load.Volume,
			Comment: load.Comment,
		}

		if lo.FromPtr(load.LoadID) == uuid.Nil {
			dLoad.LoadID = uuid.New()
			dLoadInsert = append(dLoadInsert, dLoad)
		} else {
			dLoad.LoadID = lo.FromPtr(load.LoadID)
			dLoadUpdate = append(dLoadUpdate, dLoad)
		}
	}

	return dLoadInsert, dLoadUpdate
}

func MapFeedbackToDomain(request FeedbackRequest, studentID uuid.UUID) model.Feedback {
	return model.Feedback{
		FeedbackID:     uuid.New(),
		StudentID:      studentID,
		DissertationID: request.DissertationID,
		Feedback:       request.Feedback,
		Semester:       request.Semester,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func MapApprovalStatusToDomain(status string) (model.ApprovalStatus, error) {
	status = strings.TrimSpace(strings.ToLower(status))

	switch {
	case status == model.ApprovalStatus_OnReview.String():
		return model.ApprovalStatus_OnReview, nil
	case status == model.ApprovalStatus_InProgress.String():
		return model.ApprovalStatus_InProgress, nil
	case status == model.ApprovalStatus_Approved.String():
		return model.ApprovalStatus_Approved, nil
	case status == model.ApprovalStatus_Todo.String():
		return model.ApprovalStatus_Todo, nil
	case status == model.ApprovalStatus_Failed.String():
		return model.ApprovalStatus_Failed, nil
	case status == "":
		return model.ApprovalStatus_Empty, nil
	default:
		return "", ErrUnknownApprovalStatus
	}
}
