package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/jtyler258/mbp/api/gqlgen"
)

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
