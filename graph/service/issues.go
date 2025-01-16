package service

import (
	"context"
	"log"

	"github.com/MizukiShigi/graphql-study/graph/db"
	"github.com/MizukiShigi/graphql-study/graph/model"
	"github.com/MizukiShigi/graphql-study/graph/utils/paginationutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type issueService struct {
	exec boil.ContextExecutor
}

var defaultColumns = []string{
	db.IssueTableColumns.ID,
	db.IssueTableColumns.URL,
	db.IssueTableColumns.Title,
	db.IssueColumns.Closed,
	db.IssueColumns.Number,
	db.IssueColumns.Repository,
	db.IssueColumns.Author,
}

func (i *issueService) GetIssuesByRepoID(ctx context.Context, repoID string, params *paginationutil.ListParams) (*model.IssueConnection, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	queryMods := []qm.QueryMod{
		qm.Select(defaultColumns...),
	}

	queryMods = append(queryMods, db.IssueWhere.Repository.EQ(repoID))

	if params.After != nil {
		queryMods = append(queryMods, qm.Where("id < ?", *params.After))
		queryMods = append(queryMods, qm.OrderBy("id DESC"))
	} else if params.Before != nil {
		queryMods = append(queryMods, qm.Where("id > ?", *params.Before))
		queryMods = append(queryMods, qm.OrderBy("id ASC"))
	} else {
		queryMods = append(queryMods, qm.OrderBy("id ASC"))
	}

	queryMods = append(queryMods, qm.Limit(params.GetLimit()))

	issues, err := db.Issues(
		queryMods...,
	).All(ctx, i.exec)
	if err != nil {
		return nil, err
	}

	nodes := make([]*model.Issue, len(issues))
	edges := make([]*model.IssueEdge, len(issues))
	for i, issue := range issues {
		nodes[i] = convertIssue(issue)
		edges[i] = &model.IssueEdge{
			Node:   nodes[i],
			Cursor: issue.ID,
		}
	}

	var startCursor, endCursor *string
	if len(edges) > 0 {
		start := edges[0].Cursor
		end := edges[len(edges)-1].Cursor
		startCursor = &start
		endCursor = &end
	}

	hasNext := false
	hasPrevious := false

	if len(issues) > 0 {
		if params.After != nil || params.Before != nil {
			hasNext = len(issues) == params.GetLimit()
			hasPrevious = true
		} else {
			hasNext = len(issues) == params.GetLimit()
		}
	}

	return &model.IssueConnection{
		Nodes: nodes,
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage:     hasNext,
			EndCursor:       endCursor,
			HasPreviousPage: hasPrevious,
			StartCursor:     startCursor,
		},
	}, nil
}

func (i *issueService) GetIssueByRepoNumber(ctx context.Context, repoID string, number int) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(defaultColumns...),
		db.IssueWhere.Repository.EQ(repoID),
		db.IssueWhere.Number.EQ(int64(number)),
	).One(ctx, i.exec)
	if err != nil {
		return nil, err
	}

	return convertIssue(issue), nil
}

func convertIssue(issue *db.Issue) *model.Issue {
	issueURL, err := model.UnmarshalURI(issue.URL)
	if err != nil {
		log.Println("invalid URI", issue.URL)
	}

	return &model.Issue{
		ID:         issue.ID,
		URL:        issueURL,
		Title:      issue.Title,
		Closed:     (issue.Closed == 1),
		Number:     int32(issue.Number),
		Repository: &model.Repository{ID: issue.Repository},
		Author:     &model.User{ID: issue.Author},
	}
}
