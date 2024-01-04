package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
)

type SupervisorRepository struct {
}

func NewSupervisorRepository() *SupervisorRepository {
	return &SupervisorRepository{}
}

func (r *SupervisorRepository) GetSupervisors(ctx context.Context, tx *pgxpool.Pool) ([]*model.Supervisors, error) {
	stmt, args := table.Supervisors.SELECT(table.Supervisors.AllColumns).Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetSupervisors()")
	}

	var supervisors []*model.Supervisors

	for rows.Next() {
		supervisor := &model.Supervisors{}

		if err = rows.Scan(
			&supervisor.ClientID,
			&supervisor.FullName,
			&supervisor.SupervisorID,
		); err != nil {
			return nil, errors.Wrap(err, "GetSupervisors()")
		}

		supervisors = append(supervisors, supervisor)
	}

	return supervisors, nil
}
