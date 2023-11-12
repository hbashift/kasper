package repositories

import (
	"context"

	"uir_draft/internal/generated/kasper/uir_draft/public/model"
	"uir_draft/internal/generated/kasper/uir_draft/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type tokenRepository struct {
	postgres *pgxpool.Pool
}

func NewTokenRepository(postgres *pgxpool.Pool) *tokenRepository {
	return &tokenRepository{postgres: postgres}
}

func (r *tokenRepository) Authenticate(ctx context.Context, token string) (*model.AuthorizationToken, error) {
	stmt, args := table.AuthorizationToken.
		SELECT(table.AuthorizationToken.AllColumns).
		WHERE(table.AuthorizationToken.TokenNumber.EQ(postgres.String(token))).Sql()

	var session model.AuthorizationToken

	row := r.postgres.QueryRow(ctx, stmt, args...)

	if err := scanAuthorizationToken(row, &session); err != nil {
		return nil, errors.Wrap(err, "mapping authorization token")
	}

	return &session, nil
}

func scanAuthorizationToken(row pgx.Row, target *model.AuthorizationToken) error {
	return row.Scan(
		&target.TokenID,
		&target.ClientID,
		&target.KasperID,
		&target.TokenStatus,
		&target.TokenNumber,
		&target.CreatedAt,
		&target.UpdatedAt,
		&target.ExpirationDate,
	)
}
