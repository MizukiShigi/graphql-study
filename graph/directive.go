package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MizukiShigi/graphql-study/middlewares/auth"
)

var Directive DirectiveRoot = DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
	if _, ok := auth.GetUserName(ctx); !ok {
		return nil, errors.New("Unauthorized")
	}
	return next(ctx)
}
