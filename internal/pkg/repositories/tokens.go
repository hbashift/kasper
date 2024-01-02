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

func (r *tokenRepository) Authenticate(ctx context.Context, token string, tx *pgxpool.Pool) (*model.AuthorizationToken, error) {
	stmt, args := table.AuthorizationToken.
		SELECT(table.AuthorizationToken.AllColumns).
		WHERE(table.AuthorizationToken.TokenNumber.EQ(postgres.String(token))).Sql()

	var session model.AuthorizationToken

	row := tx.QueryRow(ctx, stmt, args...)

	if err := scanAuthorizationToken(row, &session); err != nil {
		return nil, errors.Wrap(err, "mapping authorization token")
	}

	return &session, nil
}

func (r *tokenRepository) InsertToken(ctx context.Context, tx *pgxpool.Pool, token *model.AuthorizationToken) error {
	stmt, args := table.AuthorizationToken.
		INSERT(table.AuthorizationToken.AllColumns).
		MODEL(token).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *tokenRepository) GetUserTypeByToken(ctx context.Context, tx *pgxpool.Pool, token string) (*model.UserType, error) {
	stmt, args := table.AuthorizationToken.
		INNER_JOIN(table.ClientUser, table.AuthorizationToken.KasperID.EQ(table.ClientUser.KasperID)).
		SELECT(table.ClientUser.ClientType).
		WHERE(table.AuthorizationToken.TokenNumber.EQ(postgres.String(token))).Sql()

	row := tx.QueryRow(ctx, stmt, args...)

	var clientType model.UserType
	if err := row.Scan(&clientType); err != nil {
		return nil, errors.Wrap(err, "GetUserTypeByToken()")
	}

	return &clientType, nil
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
	)
}
