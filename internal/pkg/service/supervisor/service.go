package supervisor

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	TokenRepository interface {
		GetUserIDByTokenTx(ctx context.Context, tx pgx.Tx, token string) (uuid.UUID, error)
	}

	UsersRepository interface {
		GetUserTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (model.Users, error)
	}

	FeedbackRepository interface {
		UpsertFeedbackTx(ctx context.Context, tx pgx.Tx, feedback model.Feedback) error
	}

	ClientRepository interface {
		GetSupervisorsStudentsTx(ctx context.Context, tx pgx.Tx, supervisorID uuid.UUID) ([]models.Student, error)
	}

	MarksRepository interface {
		UpsertMarkTx(ctx context.Context, tx pgx.Tx, model model.Marks) error
	}
)

type Service struct {
	dissertationRepo FeedbackRepository
	tokenRepo        TokenRepository
	userRepo         UsersRepository
	client           ClientRepository
	db               *pgxpool.Pool
}

func NewService(dissertationRepo FeedbackRepository, tokenRepo TokenRepository, userRepo UsersRepository, client ClientRepository, db *pgxpool.Pool) *Service {
	return &Service{dissertationRepo: dissertationRepo, tokenRepo: tokenRepo, userRepo: userRepo, client: client, db: db}
}
