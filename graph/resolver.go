package graph

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/MizukiShigi/graphql-study/graph/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services service.Services
	*Loaders
}
