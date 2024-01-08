package student

import (
	"context"
	"time"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/pkg/service/student/mapping"

	"github.com/pkg/errors"
)

type StudentDissertationPlan struct {
	First   bool `json:"id1,omitempty"`
	Second  bool `json:"id2,omitempty"`
	Third   bool `json:"id3,omitempty"`
	Forth   bool `json:"id4,omitempty"`
	Fifth   bool `json:"id5,omitempty"`
	Sixth   bool `json:"id6,omitempty"`
	Seventh bool `json:"id7,omitempty"`
	Eighth  bool `json:"id8,omitempty"`
}

type StudCommonInfo struct {
	DissertationTitle     string  `db:"dissertation_title" json:"theme,omitempty"`
	SupervisorName        string  `db:"supervisor_name" json:"teacherFullName,omitempty"`
	EnrollmentOrderNumber string  `db:"enrollment_order_number" json:"numberOfOrderOfStatement,omitempty"`
	StudyingStartDate     string  `db:"studying_start_date" json:"dateOfOrderOfStatement"`
	Semester              int32   `db:"semester_number" json:"actualSemestr,omitempty"`
	Feedback              *string `db:"feedback" json:"feedback,omitempty"`
	TitlePageURL          string  `db:"title_page_url" json:"titlePageURL,omitempty"`
	ExplanatoryNoteURL    string  `db:"explanatory_note_url" json:"explanatoryNoteURL,omitempty"`
	StudentName           string  `db:"student_name"`
	NumberOfYears         int32   `db:"number_of_years" json:"number_of_years"`
}

type DissertationPage struct {
	DissertationPlan     map[string]*StudentDissertationPlan `json:"dissertationPlan"`
	CommonInfo           *StudCommonInfo                     `json:"commonInfo"`
	IDs                  []*mapping.DissertationIDs          `json:"ids"`
	DissertationStatuses []*mapping.DissertationStatus       `json:"statuses"`
}

func (s *Service) GetDissertationPage(ctx context.Context, token string) (*DissertationPage, error) {
	session, err := s.tokenRepo.Authenticate(ctx, token, s.db)
	if err != nil {
		return nil, errors.Wrap(err, "authentication error")
	}

	if session.TokenStatus != model.TokenStatus_Active {
		return nil, ErrNonValidToken
	}

	commonInfo, err := s.studRepo.GetStudentCommonInfo(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentCommonInfo()")
	}

	plans, err := s.semesterRepo.GetStudentDissertationPlan(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, errors.Wrap(err, "GetStudentDissertationPlan()")
	}

	domainIDs, err := s.dRepo.GetDissertationIDs(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, err
	}

	ids := mapping.MapIDsFromDomain(domainIDs)

	planMap := make(map[string]*StudentDissertationPlan, len(plans))

	for _, semester := range plans {
		plan := &StudentDissertationPlan{
			First:  semester.First,
			Second: semester.Second,
			Third:  semester.Third,
			Forth:  semester.Forth,
			Fifth:  semester.Fifth,
			Sixth:  semester.Sixth,
		}
		planMap[semester.Name] = plan
	}

	statuses, err := s.dRepo.GetStatuses(ctx, s.db, session.KasperID)
	if err != nil {
		return nil, err
	}

	startDate := commonInfo.StudyingStartDate.Format(time.DateOnly)

	info := StudCommonInfo{
		DissertationTitle:     commonInfo.DissertationTitle,
		SupervisorName:        commonInfo.SupervisorName,
		EnrollmentOrderNumber: commonInfo.EnrollmentOrderNumber,
		StudyingStartDate:     startDate,
		Semester:              commonInfo.Semester,
		Feedback:              commonInfo.Feedback,
		TitlePageURL:          commonInfo.TitlePageURL,
		ExplanatoryNoteURL:    commonInfo.ExplanatoryNoteURL,
		StudentName:           commonInfo.StudentName,
		NumberOfYears:         commonInfo.NumberOfYears,
	}

	return &DissertationPage{
		DissertationPlan:     planMap,
		CommonInfo:           &info,
		IDs:                  ids,
		DissertationStatuses: statuses,
	}, nil
}
