package service

import (
	"context"

	"github.com/MizukiShigi/graphql-study/graph/db"
	"github.com/MizukiShigi/graphql-study/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func (r *repositoryService) GetRepoByID(ctx context.Context, id string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryTableColumns.ID,
			db.RepositoryTableColumns.Name,
			db.RepositoryTableColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.ID.EQ(id),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}

func (r *repositoryService) GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryTableColumns.ID,
			db.RepositoryTableColumns.Name,
			db.RepositoryTableColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}

func convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Owner:     &model.User{ID: repo.Owner},
		Name:      repo.Name,
		CreatedAt: repo.CreatedAt,
	}
}
