package student

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/handlers/student_handler/request_models"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) GetReportComments(ctx context.Context, studentID uuid.UUID) (models.ReportComments, error) {
	var (
		planComments = make([]model.DissertationPlans, 0)
		dissComments = make([]model.DissertationCommentary, 0)
	)

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		var err error
		dissComments, err = s.commentRepo.GetDissertationComments(ctx, tx, studentID)
		if err != nil {
			return err
		}

		planComments, err = s.commentRepo.GetPlanComments(ctx, tx, studentID)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return models.ReportComments{}, errors.Wrap(err, "GetReportComments()")
	}

	respPlans := make([]models.DissertationPlan, 0)
	for _, plan := range planComments {
		respPlan := models.DissertationPlan{
			PlanID:    plan.PlanID,
			StudentID: plan.StudentID,
			Semester:  plan.Semester,
			PlanText:  plan.PlanText,
		}

		respPlans = append(respPlans, respPlan)
	}

	respComments := make([]models.DissertationComment, 0)
	for _, comment := range dissComments {
		respComment := models.DissertationComment{
			CommentaryID: comment.CommentaryID,
			StudentID:    comment.StudentID,
			Semester:     comment.Semester,
			Commentary:   comment.Commentary,
		}

		respComments = append(respComments, respComment)
	}

	return models.ReportComments{
		DissertationComments: respComments,
		DissertationPlans:    respPlans,
	}, nil
}

func (s *Service) UpsertReportComments(ctx context.Context, studentID uuid.UUID, req request_models.UpsertReportCommentsRequest) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		student, err := s.studRepo.GetStudentTx(ctx, tx, studentID)
		if err != nil {
			return err
		}

		if student.Status == model.ApprovalStatus_OnReview || student.Status == model.ApprovalStatus_Approved {
			return models.ErrNonMutableStatus
		}

		if err := s.commentRepo.UpsertDissertationComment(ctx, tx, model.DissertationCommentary{
			CommentaryID: uuid.New(),
			StudentID:    studentID,
			Semester:     req.Semester,
			Commentary:   req.DissertationComment.Commentary,
		}); err != nil {
			return err
		}

		if err := s.commentRepo.UpsertPlanComment(ctx, tx, model.DissertationPlans{
			PlanID:    uuid.New(),
			StudentID: studentID,
			Semester:  req.Semester,
			PlanText:  req.DissertationPlan.PlanText,
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "UpsertReportComments()")
	}

	return nil
}
