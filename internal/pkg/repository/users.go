package repository

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/generated/new_kasper/new_uir/public/table"
	"uir_draft/internal/pkg/models"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) InsertUsersTx(ctx context.Context, tx pgx.Tx, users []model.Users) error {
	stmt, args := table.Users.
		INSERT(table.Users.AllColumns.Except(table.Users.Registered)).
		MODELS(users).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "InsertUsersTx()")
	}

	return nil
}

func (r *UsersRepository) GetUserTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (model.Users, error) {
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

func (r *UsersRepository) SetUserRegisteredTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) error {
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

func (r *UsersRepository) ChangeUsersPasswordTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID, password string) error {
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

func (r *UsersRepository) DeleteUserCascadeTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) error {
	stmt, args := table.Users.
		DELETE().
		WHERE(table.Users.UserID.EQ(postgres.UUID(userID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "DeleteUserCascadeTx()")
	}

	return nil
}

func (r *UsersRepository) GetUserByEmailTx(ctx context.Context, tx pgx.Tx, email string) (model.Users, error) {
	stmt, args := table.Users.
		SELECT(table.Users.AllColumns).
		WHERE(table.Users.Email.EQ(postgres.String(email))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)

	user := model.Users{}
	if err := scanUser(row, &user); err != nil {
		return model.Users{}, errors.Wrap(err, "GetUserByEmailTx()")
	}

	return user, nil
}

func (r *UsersRepository) GetUserByKasperIDTx(ctx context.Context, tx pgx.Tx, kasperID uuid.UUID) (model.Users, error) {
	stmt, args := table.Users.
		SELECT(table.Users.AllColumns).
		WHERE(table.Users.KasperID.EQ(postgres.UUID(kasperID))).
		Sql()

	row := tx.QueryRow(ctx, stmt, args...)

	user := model.Users{}
	if err := scanUser(row, &user); err != nil {
		return model.Users{}, errors.Wrap(err, "GetUserByKasperIDTx()")
	}

	return user, nil
}

func (r *UsersRepository) ChangeUsersEmail(ctx context.Context, tx pgx.Tx, userID uuid.UUID, email string) error {
	stmt, args := table.Users.
		UPDATE(table.Users.Email).
		SET(email).
		WHERE(table.Users.UserID.EQ(postgres.UUID(userID))).
		Sql()

	if _, err := tx.Exec(ctx, stmt, args...); err != nil {
		return errors.Wrap(err, "ChangeUsersEmail()")
	}

	return nil
}

func (r *UsersRepository) GetNotRegisteredUsers(ctx context.Context, tx pgx.Tx) ([]models.UserInfo, error) {
	stmt, args := table.Users.
		SELECT(table.Users.UserID, table.Users.Email, table.Users.Registered, table.Users.UserType).
		WHERE(table.Users.Registered.EQ(postgres.Bool(false))).
		Sql()

	rows, err := tx.Query(ctx, stmt, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetNotRegisteredUsers()")
	}

	users := make([]models.UserInfo, 0)

	for rows.Next() {
		user := models.UserInfo{}
		if err := scanUserInfo(rows, &user); err != nil {
			return nil, errors.Wrap(err, "GetNotRegisteredUsers(): scanning row")
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UsersRepository) DeleteNotRegisteredUsers(ctx context.Context, tx pgx.Tx, userIDs []uuid.UUID) error {
	idExpressions := make([]postgres.Expression, 0, len(userIDs))

	for _, id := range userIDs {
		idExp := postgres.UUID(id)

		idExpressions = append(idExpressions, idExp)
	}

	stmt, args := table.Users.
		DELETE().
		WHERE(table.Users.UserID.IN(idExpressions...).AND(table.Users.Registered.EQ(postgres.Bool(false)))).
		Sql()

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return errors.Wrap(err, "DeleteNotRegisteredUsers()")
	}

	return nil
}

func scanUserInfo(row pgx.Row, target *models.UserInfo) error {
	return row.Scan(
		&target.UserID,
		&target.Email,
		&target.Registered,
		&target.UserType,
	)
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
