package service

import (
	"context"

	"github.com/MizukiShigi/graphql-study/graph/model"
	"github.com/MizukiShigi/graphql-study/graph/utils/paginationutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	*repositoryService
	*issueService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

type UserService interface {
	AddUserById(ctx context.Context, id string, name string) (*model.AddUserByIDPayload, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUsersByIDs(ctx context.Context, IDs []string) ([]*model.User, error)
}

type IssueService interface {
	GetIssuesByRepoID(ctx context.Context, repoID string, params *paginationutil.ListParams) (*model.IssueConnection, error)
	GetIssueByRepoNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
}

type RepositoryService interface {
	GetRepoByID(ctx context.Context, id string) (*model.Repository, error)
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

func New(exec boil.ContextExecutor) Services {
	boil.DebugMode = true
	return &services{
		userService: &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
		issueService: &issueService{exec: exec},
	}
}
