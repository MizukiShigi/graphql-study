package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"strings"

	"github.com/MizukiShigi/graphql-study/graph/model"
	"github.com/MizukiShigi/graphql-study/graph/utils/paginationutil"
)

// Author is the resolver for the author field.
func (r *issueResolver) Author(ctx context.Context, obj *model.Issue) (*model.User, error) {
	thunk := r.Loaders.UserLoader.Load(ctx, obj.Author.ID)
	user, err := thunk()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error) {
	panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
}

// AddUserByID is the resolver for the addUserById field.
func (r *mutationResolver) AddUserByID(ctx context.Context, input model.AddUserByIDInput) (*model.AddUserByIDPayload, error) {
	return r.Services.AddUserById(ctx, input.ID, input.Name)
}

// Owner is the resolver for the owner field.
func (r *projectV2Resolver) Owner(ctx context.Context, obj *model.ProjectV2) (*model.User, error) {
	return r.Services.GetUserByID(ctx, obj.Owner.ID)
}

// Repository is the resolver for the repository field.
func (r *queryResolver) Repository(ctx context.Context, name string, owner string) (*model.Repository, error) {
	return r.Services.GetRepoByFullName(ctx, owner, name)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	return r.Services.GetUserByName(ctx, name)
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	switch {
	case strings.HasPrefix(id, "U_"):
		return r.Services.GetUserByID(ctx, id)
	case strings.HasPrefix(id, "REPO_"):
		return r.Services.GetRepoByID(ctx, id)
	default:
		return nil, fmt.Errorf("invalid id: %s", id)
	}
}

// Owner is the resolver for the owner field.
func (r *repositoryResolver) Owner(ctx context.Context, obj *model.Repository) (*model.User, error) {
	return r.Services.GetUserByID(ctx, obj.Owner.ID)
}

// Issue is the resolver for the issue field.
func (r *repositoryResolver) Issue(ctx context.Context, obj *model.Repository, number int32) (*model.Issue, error) {
	return r.Services.GetIssueByRepoNumber(ctx, obj.ID, int(number))
}

// Issues is the resolver for the issues field.
func (r *repositoryResolver) Issues(ctx context.Context, obj *model.Repository, after *string, before *string, first *int32, last *int32) (*model.IssueConnection, error) {
	return r.Services.GetIssuesByRepoID(ctx, obj.ID,
		&paginationutil.ListParams{After: after, Before: before, First: first, Last: last},
	)
}

// PullRequest is the resolver for the pullRequest field.
func (r *repositoryResolver) PullRequest(ctx context.Context, obj *model.Repository, number int32) (*model.PullRequest, error) {
	panic(fmt.Errorf("not implemented: PullRequest - pullRequest"))
}

// PullRequests is the resolver for the pullRequests field.
func (r *repositoryResolver) PullRequests(ctx context.Context, obj *model.Repository, after *string, before *string, first *int32, last *int32) (*model.PullRequestConnection, error) {
	panic(fmt.Errorf("not implemented: PullRequests - pullRequests"))
}

// Issue returns IssueResolver implementation.
func (r *Resolver) Issue() IssueResolver { return &issueResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// ProjectV2 returns ProjectV2Resolver implementation.
func (r *Resolver) ProjectV2() ProjectV2Resolver { return &projectV2Resolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Repository returns RepositoryResolver implementation.
func (r *Resolver) Repository() RepositoryResolver { return &repositoryResolver{r} }

type issueResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type projectV2Resolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type repositoryResolver struct{ *Resolver }
