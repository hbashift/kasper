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

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) InsertUsersTx(ctx context.Context, tx *pgxpool.Tx, users []model.Users) error {
	stmt, args := table.Users.
		INSERT().
		MODELS(users).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertUsersTx()")
	}

	return nil
}

func (r *UsersRepository) GetUserTx(ctx context.Context, tx *pgxpool.Tx, userID uuid.UUID) (model.Users, error) {
	stmt, args := table.Users.
		SELECT(table.Users.AllColumns).
		WHERE(table.Users.UserID.EQ(postgres.UUID(userID))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)

	user := model.Users{}
	if err := scanUser(row, &user); err != nil {
		return model.Users{}, errors.Wrap(err, "GetUserTx()")
	}

	return user, nil
}

func (r *UsersRepository) SetUserRegisteredTx(ctx context.Context, tx *pgxpool.Tx, userID uuid.UUID) error {
	stmt, args := table.Users.
		UPDATE(table.Users.Registered).
		SET(true).
		WHERE(table.Users.UserID.EQ(postgres.UUID(userID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "SetUserRegisteredTx()")
	}

	return nil
}

func (r *UsersRepository) ChangeUsersPasswordTx(ctx context.Context, tx *pgxpool.Tx, userID uuid.UUID, password string) error {
	stmt, args := table.Users.
		UPDATE(table.Users.Password).
		SET(password).
		WHERE(table.Users.UserID.EQ(postgres.UUID(userID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "ChangeUsersPasswordTx()")
	}

	return nil
}

func scanUser(row pgx.Row, target *model.Users) error {
	return row.Scan(
		&target.UserID,
		&target.Email,
		&target.Password,
		&target.KasperID,
		&target.UserType,
		&target.Registered,
	)
}
