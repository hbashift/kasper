package repositories

import (
	"context"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
)

type ClientUserRepository struct {
}

func NewClientUserRepository() *ClientUserRepository {
	return &ClientUserRepository{}
}

func (c *ClientUserRepository) GetClient(ctx context.Context, tx *pgxpool.Pool, email string) (*model.ClientUser, error) {
	var clientUser *model.ClientUser

	err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		stmt, args := table.ClientUser.
			SELECT(table.ClientUser.AllColumns).
			WHERE(table.ClientUser.Email.EQ(postgres.String(email))).
			Sql()

		row := tx.QueryRow(ctx, stmt, args...)

		return scanClientUser(row, clientUser)
	})
	if err != nil {
		return nil, err
	}

	return clientUser, err
}

func scanClientUser(row pgx.Row, target *model.ClientUser) error {
	return row.Scan(
		&target.Email,
		&target.ClientType,
		&target.Password,
		&target.ClientID,
	)
}
