package student

import (
	"context"
	"time"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func (s *Service) GetAllMarks(ctx context.Context, studentID uuid.UUID) (models.AllMarks, error) {
	exams := make([]model.Exams, 0)
	attestationMarks := make([]model.Marks, 0)
	supervisorMarks := make([]model.SupervisorMarks, 0)

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		var err error
		attestationMarks, err = s.marksRepo.GetStudentsAttestationMarksTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		supervisorMarks, err = s.marksRepo.GetStudentsSupervisorMarks(ctx, tx, studentID)
		if err != nil {
			return err
		}

		exams, err = s.marksRepo.GetStudentsExamResults(ctx, tx, studentID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return models.AllMarks{}, errors.Wrap(err, "GetAllMarks()")
	}

	respExams := make([]models.Exam, 0)
	for _, exam := range exams {
		respExam := models.Exam{
			ExamID:    exam.ExamID,
			StudentID: exam.StudentID,
			ExamType:  exam.ExamType,
			Semester:  exam.Semester,
			Mark:      exam.Mark,
			SetAt:     exam.SetAt,
		}

		respExams = append(respExams, respExam)
	}

	respSupMarks := make([]models.SupervisorMark, 0)
	for _, mark := range supervisorMarks {
		respMark := models.SupervisorMark{
			MarkID:       mark.MarkID,
			StudentID:    mark.StudentID,
			Mark:         mark.Mark,
			Semester:     mark.Semester,
			SupervisorID: mark.SupervisorID,
		}

		respSupMarks = append(respSupMarks, respMark)
	}

	respAttMarks := make([]models.AttestationMark, 0)
	for _, mark := range attestationMarks {
		respMark := models.AttestationMark{
			StudentID: mark.StudentID,
			Mark:      mark.Mark,
			Semester:  mark.Semester,
		}

		respAttMarks = append(respAttMarks, respMark)
	}

	return models.AllMarks{
		Exams:            respExams,
		SupervisorMarks:  respSupMarks,
		AttestationMarks: respAttMarks,
	}, nil
}

func (s *Service) UpsertExamResults(ctx context.Context, studentID uuid.UUID, exams []models.ExamRequest) error {
	dExams := make([]model.Exams, 0, len(exams))

	for _, exam := range exams {
		dExam := model.Exams{
			ExamID:    uuid.New(),
			StudentID: studentID,
			ExamType:  exam.ExamType,
			Semester:  exam.Semester,
			Mark:      exam.Mark,
			SetAt:     lo.ToPtr(time.Now()),
		}

		dExams = append(dExams, dExam)
	}

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		err := s.marksRepo.UpsertExamResults(ctx, tx, dExams)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "UpsertExamResults()")
	}

	return nil
}
