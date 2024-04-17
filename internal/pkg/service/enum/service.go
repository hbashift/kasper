package enum

import (
	"context"

	"uir_draft/internal/generated/new_kasper/new_uir/public/model"
	"uir_draft/internal/pkg/models"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type (
	EnumRepository interface {
		GetSpecializationsTx(ctx context.Context, tx pgx.Tx) ([]model.Specializations, error)
		InsertSpecializationsTx(ctx context.Context, tx pgx.Tx, specs []model.Specializations) error
		UpdateSpecializationTx(ctx context.Context, tx pgx.Tx, spec model.Specializations) error
		ArchiveSpecializations(ctx context.Context, tx pgx.Tx, specsIDs []int32) error

		GetGroupsTx(ctx context.Context, tx pgx.Tx) ([]model.Groups, error)
		InsertGroupsTx(ctx context.Context, tx pgx.Tx, groups []model.Groups) error
		UpdateGroupTx(ctx context.Context, tx pgx.Tx, group model.Groups) error
		DeleteGroupsTx(ctx context.Context, tx pgx.Tx, groupsIDs []int32) error
	}
)

type Service struct {
	repo EnumRepository
	db   *pgxpool.Pool
}

func NewService(repo EnumRepository, db *pgxpool.Pool) *Service {
	return &Service{repo: repo, db: db}
}

func (s *Service) GetSpecializations(ctx context.Context) ([]models.Specialization, error) {
	domainSpecs := make([]model.Specializations, 0)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		specs, err := s.repo.GetSpecializationsTx(ctx, tx)
		if err != nil {
			return err
		}

		domainSpecs = specs
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetSpecializations()")
	}

	specs := make([]models.Specialization, 0, len(domainSpecs))
	for _, dSpec := range domainSpecs {
		spec := models.Specialization{
			SpecializationID: dSpec.SpecID,
			Name:             dSpec.Title,
		}

		specs = append(specs, spec)
	}

	return specs, nil
}

func (s *Service) GetGroups(ctx context.Context) ([]models.Group, error) {
	domainGroups := make([]model.Groups, 0)

	err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		groups, err := s.repo.GetGroupsTx(ctx, tx)
		if err != nil {
			return err
		}

		domainGroups = groups
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetGroups()")
	}

	groups := make([]models.Group, 0, len(domainGroups))
	for _, dSpec := range domainGroups {
		spec := models.Group{
			GroupID: dSpec.GroupID,
			Name:    dSpec.GroupName,
		}

		groups = append(groups, spec)
	}

	return groups, nil
}

func (s *Service) InsertGroups(ctx context.Context, groups []models.Group) error {
	domainGroups := make([]model.Groups, 0, len(groups))
	for _, group := range groups {
		dGroup := model.Groups{
			GroupName: group.Name,
		}

		domainGroups = append(domainGroups, dGroup)
	}

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		return s.repo.InsertGroupsTx(ctx, tx, domainGroups)
	}); err != nil {
		return errors.Wrap(err, "InsertGroups()")
	}

	return nil
}

func (s *Service) InsertSpecializations(ctx context.Context, specializations []models.Specialization) error {
	domainSpec := make([]model.Specializations, 0, len(specializations))
	for _, group := range specializations {
		dGroup := model.Specializations{
			SpecID: group.SpecializationID,
			Title:  group.Name,
		}

		domainSpec = append(domainSpec, dGroup)
	}

	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		return s.repo.InsertSpecializationsTx(ctx, tx, domainSpec)
	}); err != nil {
		return errors.Wrap(err, "InsertSpecializations()")
	}

	return nil
}

func (s *Service) DeleteGroups(ctx context.Context, groupIDs []int32) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		return s.repo.DeleteGroupsTx(ctx, tx, groupIDs)
	}); err != nil {
		return errors.Wrap(err, "DeleteGroups()")
	}

	return nil
}

func (s *Service) DeleteSpecializations(ctx context.Context, specIDs []int32) error {
	if err := s.db.BeginFunc(ctx, func(tx pgx.Tx) error {
		return s.repo.ArchiveSpecializations(ctx, tx, specIDs)
	}); err != nil {
		return errors.Wrap(err, "DeleteSpecializations()")
	}

	return nil
}
