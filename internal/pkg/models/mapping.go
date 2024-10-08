package models

import (
	"strings"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func MapDissertationPageFromDomain(
	dSemesterProgresses []model.SemesterProgress,
	dDissertations []model.Dissertations,
	dTitles []model.DissertationTitles,
	dFeedbacks []model.Feedback,
	dComments []model.StudentsCommentary,
	dProgresses []model.Progressiveness,
) DissertationPageResponse {
	semesterProgresses := make([]SemesterProgressResponse, 0, len(dSemesterProgresses))
	for _, dSemesterProgress := range dSemesterProgresses {
		semesterProgress := SemesterProgressResponse{}
		semesterProgress.SetDomainData(dSemesterProgress)

		semesterProgresses = append(semesterProgresses, semesterProgress)
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

	comments := make([]StudentComment, 0, len(dComments))
	for _, dComment := range dComments {
		comment := StudentComment{}
		comment.SetDomainData(dComment)

		comments = append(comments, comment)
	}

	progresses := make([]Progressiveness, 0, len(dProgresses))
	for _, dProgress := range dProgresses {
		progress := Progressiveness{}
		progress.SetDomainData(dProgress)

		progresses = append(progresses, progress)
	}

	return DissertationPageResponse{
		SemesterProgress:      semesterProgresses,
		DissertationsStatuses: dissertations,
		DissertationTitles:    titles,
		Feedback:              feedbacks,
		StudentsComments:      comments,
		Progresses:            progresses,
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

func MapPublicationsToDomain(publications []Publication, worksID uuid.UUID) (dPublicationsInsert, dPublicationsUpdate []model.Publications, err error) {
	for _, publication := range publications {
		var status model.PublicationStatus = ""
		if err := status.Scan(strings.TrimSpace(lo.FromPtr(publication.Status))); err != nil {
			return nil, nil, ErrInvalidEnumValue
		}

		dPublication := model.Publications{
			WorksID:    worksID,
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

func MapConferencesToDomain(conferences []Conference, worksID uuid.UUID) (dConferencesInsert, dConferencesUpdate []model.Conferences, err error) {
	for _, conf := range conferences {
		var status model.ConferenceStatus = ""
		if err := status.Scan(strings.TrimSpace(lo.FromPtr(conf.Status))); err != nil {
			return nil, nil, ErrInvalidEnumValue
		}

		dConf := model.Conferences{
			WorksID:        worksID,
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

func MapResearchProjectToDomain(projects []ResearchProject, worksID uuid.UUID) (dResearchInsert, dResearchUpdate []model.ResearchProjects) {
	for _, proj := range projects {
		dProj := model.ResearchProjects{
			WorksID:     worksID,
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

func MapPatentsToDomain(patents []Patent, worksID uuid.UUID) (dPatentsInsert, dPatentsUpdate []model.Patents, err error) {
	for _, patent := range patents {
		var patentType model.PatentType = ""
		if err := patentType.Scan(strings.TrimSpace(patent.Type)); err != nil {
			return nil, nil, errors.Wrap(err, ErrInvalidEnumValue.Error())
		}

		dPatent := model.Patents{
			WorksID:          worksID,
			Name:             patent.Name,
			RegistrationDate: patent.RegistrationDate,
			Type:             patentType,
			AddInfo:          patent.AddInfo,
		}

		if lo.FromPtr(patent.PatentID) == uuid.Nil {
			dPatent.PatentID = uuid.New()
			dPatentsInsert = append(dPatentsInsert, dPatent)
		} else {
			dPatent.PatentID = lo.FromPtr(patent.PatentID)
			dPatentsUpdate = append(dPatentsUpdate, dPatent)
		}
	}

	return dPatentsInsert, dPatentsUpdate, nil
}

func MapClassroomLoadToDomain(loads []ClassroomLoad, tLoadID uuid.UUID) (dLoadInsert, dLoadUpdate []model.ClassroomLoad, err error) {
	for _, load := range loads {
		var loadType model.ClassroomLoadType
		if err := loadType.Scan(strings.TrimSpace(lo.FromPtr(load.LoadType))); err != nil {
			return nil, nil, ErrInvalidEnumValue
		}

		dLoad := model.ClassroomLoad{
			TLoadID:     tLoadID,
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

func MapIndividualWorkToDomain(loads []IndividualStudentsLoad, tLoadID uuid.UUID) (dLoadInsert, dLoadUpdate []model.IndividualStudentsLoad, err error) {
	for _, load := range loads {
		var loadType model.IndividualStudentsLoadType
		if err = loadType.Scan(strings.TrimSpace(lo.FromPtr(load.LoadType))); err != nil {
			return nil, nil, errors.Wrap(err, "MapIndividualWorkToDomain()")
		}

		dLoad := model.IndividualStudentsLoad{
			TLoadID:        tLoadID,
			StudentsAmount: lo.FromPtr(load.StudentsAmount),
			Comment:        load.Comment,
			LoadType:       loadType,
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

func MapAdditionalLoadToDomain(loads []AdditionalLoad, tLoadID uuid.UUID) (dLoadInsert, dLoadUpdate []model.AdditionalLoad) {
	for _, load := range loads {
		dLoad := model.AdditionalLoad{
			TLoadID: tLoadID,
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

func MapFeedbackToDomain(request FeedbackRequest, studentID, dissertationID uuid.UUID) model.Feedback {
	return model.Feedback{
		FeedbackID:     uuid.New(),
		StudentID:      studentID,
		DissertationID: dissertationID,
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

func MapPublicationsFromDomain(dPublications []model.Publications) []Publication {
	publications := make([]Publication, 0, len(dPublications))

	for _, dPublication := range dPublications {
		publication := Publication{
			WorksID:       dPublication.WorksID,
			PublicationID: lo.ToPtr(dPublication.PublicationID),
			Name:          lo.ToPtr(dPublication.Name),
			Scopus:        lo.ToPtr(dPublication.Scopus),
			Rinc:          lo.ToPtr(dPublication.Rinc),
			Wac:           lo.ToPtr(dPublication.Wac),
			Wos:           lo.ToPtr(dPublication.Wos),
			Impact:        lo.ToPtr(dPublication.Impact),
			Status:        lo.ToPtr(dPublication.Status.String()),
			OutputData:    dPublication.OutputData,
			CoAuthors:     dPublication.CoAuthors,
			Volume:        dPublication.Volume,
		}

		publications = append(publications, publication)
	}

	return publications
}

func MapConferencesFromDomain(dConferences []model.Conferences) []Conference {
	conferences := make([]Conference, 0, len(dConferences))

	for _, dConf := range dConferences {
		conf := Conference{
			WorksID:        dConf.WorksID,
			ConferenceID:   lo.ToPtr(dConf.ConferenceID),
			Status:         lo.ToPtr(dConf.Status.String()),
			Scopus:         lo.ToPtr(dConf.Scopus),
			Rinc:           lo.ToPtr(dConf.Rinc),
			Wac:            lo.ToPtr(dConf.Wac),
			Wos:            lo.ToPtr(dConf.Wos),
			ConferenceName: lo.ToPtr(dConf.ConferenceName),
			ReportName:     lo.ToPtr(dConf.ReportName),
			Location:       lo.ToPtr(dConf.Location),
			ReportedAt:     lo.ToPtr(dConf.ReportedAt),
		}

		conferences = append(conferences, conf)
	}

	return conferences
}

func MapResearchProjectFromDomain(dProjects []model.ResearchProjects) []ResearchProject {
	projects := make([]ResearchProject, 0, len(dProjects))

	for _, dProj := range dProjects {
		proj := ResearchProject{
			WorksID:     dProj.WorksID,
			ProjectID:   lo.ToPtr(dProj.ProjectID),
			ProjectName: lo.ToPtr(dProj.ProjectName),
			StartAt:     lo.ToPtr(dProj.StartAt),
			EndAt:       lo.ToPtr(dProj.EndAt),
			AddInfo:     dProj.AddInfo,
			Grantee:     dProj.Grantee,
		}

		projects = append(projects, proj)
	}

	return projects
}

func MapPatentsFromDomain(dPatents []model.Patents) []Patent {
	patents := make([]Patent, 0, len(dPatents))

	for _, dPatent := range dPatents {
		patent := Patent{
			WorksID:          dPatent.WorksID,
			PatentID:         lo.ToPtr(dPatent.PatentID),
			Name:             dPatent.Name,
			RegistrationDate: dPatent.RegistrationDate,
			Type:             dPatent.Type.String(),
			AddInfo:          dPatent.AddInfo,
		}

		patents = append(patents, patent)
	}

	return patents
}

func MapClassroomLoadFromDomain(dLoads []model.ClassroomLoad) []ClassroomLoad {
	loads := make([]ClassroomLoad, 0, len(dLoads))

	for _, dLoad := range dLoads {
		load := ClassroomLoad{
			TLoadID:     dLoad.TLoadID,
			LoadID:      lo.ToPtr(dLoad.LoadID),
			Hours:       lo.ToPtr(dLoad.Hours),
			LoadType:    lo.ToPtr(dLoad.LoadType.String()),
			MainTeacher: lo.ToPtr(dLoad.MainTeacher),
			GroupName:   lo.ToPtr(dLoad.GroupName),
			SubjectName: lo.ToPtr(dLoad.SubjectName),
		}

		loads = append(loads, load)
	}

	return loads
}

func MapIndividualWorkFromDomain(dLoads []model.IndividualStudentsLoad) []IndividualStudentsLoad {
	loads := make([]IndividualStudentsLoad, 0, len(dLoads))

	for _, dLoad := range dLoads {
		load := IndividualStudentsLoad{
			TLoadID:        dLoad.TLoadID,
			LoadID:         lo.ToPtr(dLoad.LoadID),
			StudentsAmount: lo.ToPtr(dLoad.StudentsAmount),
			LoadType:       lo.ToPtr(dLoad.LoadType.String()),
			Comment:        dLoad.Comment,
		}

		loads = append(loads, load)
	}

	return loads
}

func MapAdditionalLoadFromDomain(dLoads []model.AdditionalLoad) []AdditionalLoad {
	loads := make([]AdditionalLoad, 0, len(dLoads))

	for _, dLoad := range dLoads {
		load := AdditionalLoad{
			TLoadID: dLoad.TLoadID,
			LoadID:  lo.ToPtr(dLoad.LoadID),
			Name:    lo.ToPtr(dLoad.Name),
			Volume:  dLoad.Volume,
			Comment: dLoad.Comment,
		}

		loads = append(loads, load)
	}

	return loads
}

func ConvertScientificWorksToResponse(
	studentID uuid.UUID,
	dWorks []model.ScientificWorksStatus,
	publications []Publication,
	conferences []Conference,
	projects []ResearchProject,
	patents []Patent,
) []ScientificWork {
	scientificWorks := make([]ScientificWork, 0, len(dWorks))

	for i, work := range dWorks {
		scientificWorks = append(scientificWorks, ScientificWork{
			WorksID:          work.WorksID,
			Semester:         int(work.Semester),
			StudentID:        studentID,
			ApprovalStatus:   work.Status.String(),
			UpdatedAt:        work.UpdatedAt,
			AcceptedAt:       work.AcceptedAt,
			Publications:     []Publication{},
			Conferences:      []Conference{},
			ResearchProjects: []ResearchProject{},
		})

		for _, publication := range publications {
			if publication.WorksID == work.WorksID {
				scientificWorks[i].Publications = append(scientificWorks[i].Publications, publication)
			}
		}

		for _, conference := range conferences {
			if conference.WorksID == work.WorksID {
				scientificWorks[i].Conferences = append(scientificWorks[i].Conferences, conference)
			}
		}

		for _, project := range projects {
			if project.WorksID == work.WorksID {
				scientificWorks[i].ResearchProjects = append(scientificWorks[i].ResearchProjects, project)
			}
		}

		for _, patent := range patents {
			if patent.WorksID == work.WorksID {
				scientificWorks[i].Patents = append(scientificWorks[i].Patents, patent)
			}
		}
	}
	return scientificWorks
}

func ConvertTeachingLoadsToResponse(
	studentID uuid.UUID,
	dLoads []model.TeachingLoadStatus,
	classroom []ClassroomLoad,
	additional []AdditionalLoad,
	individual []IndividualStudentsLoad,
) []TeachingLoad {
	teachingLoads := make([]TeachingLoad, 0, len(dLoads))

	for i, load := range dLoads {
		teachingLoads = append(teachingLoads, TeachingLoad{
			TLoadID:                 load.LoadsID,
			StudentID:               studentID,
			Semester:                int(load.Semester),
			ApprovalStatus:          load.Status.String(),
			UpdatedAt:               load.UpdatedAt,
			AcceptedAt:              load.AcceptedAt,
			ClassroomLoads:          []ClassroomLoad{},
			IndividualStudentsLoads: []IndividualStudentsLoad{},
			AdditionalLoads:         []AdditionalLoad{},
		})

		for _, class := range classroom {
			if class.TLoadID == load.LoadsID {
				teachingLoads[i].ClassroomLoads = append(teachingLoads[i].ClassroomLoads, class)
			}
		}

		for _, add := range additional {
			if add.TLoadID == load.LoadsID {
				teachingLoads[i].AdditionalLoads = append(teachingLoads[i].AdditionalLoads, add)
			}
		}

		for _, ind := range individual {
			if ind.TLoadID == load.LoadsID {
				teachingLoads[i].IndividualStudentsLoads = append(teachingLoads[i].IndividualStudentsLoads, ind)
			}
		}
	}
	return teachingLoads
}
