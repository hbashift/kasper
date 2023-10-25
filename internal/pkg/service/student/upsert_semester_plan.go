package student

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
)

func (s *Service) UpsertSemesterPlan(ctx context.Context, progress []*model.SemesterProgress) error {
	// TODO маппинг

	return s.semesterRepo.UpsertSemesterPlan(ctx, s.db, progress)
}
