package repositories

import (
	"context"
	"strings"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"
)

type ClientUserRepository struct {
}

func NewClientUserRepository() *ClientUserRepository {
	return &ClientUserRepository{}
}

func (r *ClientUserRepository) GetClient(ctx context.Context, tx *pgxpool.Pool, email string) (*model.ClientUser, error) {
	clientUser := &model.ClientUser{}

	err := tx.BeginFunc(ctx, func(tx pgx.Tx) error {
		stmt, args := table.ClientUser.
			SELECT(table.ClientUser.AllColumns).
			WHERE(table.ClientUser.Email.EQ(postgres.String(strings.TrimSpace(email)))).
			Sql()

		row := tx.QueryRow(ctx, stmt, args...)

		err := scanClientUser(row, clientUser)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return clientUser, err
}

func (r *ClientUserRepository) ChangePassword(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID, password string) error {
	stmt, args := table.ClientUser.
		UPDATE(table.ClientUser.Password).
		SET(password).
		WHERE(table.ClientUser.ClientID.EQ(postgres.UUID(clientID))).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "ChangePassword()")
	}

	return nil
}

func (r *ClientUserRepository) GetClientByClientID(ctx context.Context, tx *pgxpool.Pool, clientID uuid.UUID) (*model.ClientUser, error) {
	stmt, args := table.ClientUser.
		SELECT(table.ClientUser.AllColumns).
		WHERE(table.ClientUser.ClientID.EQ(postgres.UUID(clientID))).
		Sql()
	clientUser := &model.ClientUser{}
	row := tx.QueryRow(ctx, stmt, args...)

	err := scanClientUser(row, clientUser)
	if err != nil {
		return nil, err
	}

	return clientUser, nil

}

func scanClientUser(row pgx.Row, target *model.ClientUser) error {
	return row.Scan(
		&target.Email,
		&target.ClientType,
		&target.Password,
		&target.ClientID,
		&target.KasperID,
		&target.Registered,
	)
}
