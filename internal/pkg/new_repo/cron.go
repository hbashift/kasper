package new_repo

import (
	"time"
)

const (
	jobInterval = 1 * time.Hour
)

//
//func CronUpdateSemester(ctx context.Context, tx *pgxpool.Pool) {
//	ticker := time.NewTicker(jobInterval)
//
//	go func() {
//		defer func() {
//			if p := recover(); p != nil {
//				log.Errorf(ctx, "recovered from %v", p)
//			}
//		}()
//
//		for {
//			select {
//			case <-ticker.C:
//				update(ctx, tx)
//			case <-ctx.Done():
//				ticker.Stop()
//				return
//			}
//		}
//	}()
//}
//
//func update(ctx context.Context, tx *pgxpool.Pool) {
//	err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
//		table.Students.
//			UPDATE(
//				table.Students.ActualSemester,
//				table.Students.Status,
//			).
//			SET(
//				table.Students.ActualSemester.ADD(postgres.Int32(1)),
//				model.ApprovalStatus_Empty,
//			).
//			WHERE(table.Students.StudyingStatus.EQ(postgres.String(model.StudentStatus_Studying.String()))).
//			Sql()
//		return nil
//	})
//	if err != nil {
//		log.Errorf(ctx, "updating students: %v", err.Error())
//	}
//}
