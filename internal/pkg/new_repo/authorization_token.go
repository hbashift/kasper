package new_repo

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type TokenRepository struct{}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}

func (r *TokenRepository) InsertTokenTx(ctx context.Context, tx *pgxpool.Tx, token *model.AuthorizationToken) error {
	stmt, args := table.AuthorizationToken.
		INSERT(table.AuthorizationToken.AllColumns).
		MODEL(token).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertTokenTx()")
	}

	return nil
}

func (r *TokenRepository) GetUsersActiveTokensTx(ctx context.Context, tx *pgxpool.Tx, userID uuid.UUID) ([]model.AuthorizationToken, error) {
	stmt, args := table.AuthorizationToken.
		SELECT(table.AuthorizationToken.AllColumns).
		WHERE(
			table.AuthorizationToken.UserID.EQ(postgres.UUID(userID)).
				AND(table.AuthorizationToken.IsActive.EQ(postgres.Bool(true))),
		).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetUsersActiveTokensTx()")
	}

	tokens := make([]model.AuthorizationToken, 0)

	for rows.Next() {
		token := model.AuthorizationToken{}

		if err := scanToken(rows, &token); err != nil {
			return nil, errors.Wrap(err, "GetUsersActiveTokensTx(): scanning row")
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func scanToken(row pgx.Row, target *model.AuthorizationToken) error {
	return row.Scan(
		&target.TokenID,
		&target.UserID,
		&target.IsActive,
		&target.TokenNumber,
		&target.CreatedAt,
		&target.UpdatedAt,
	)
}
