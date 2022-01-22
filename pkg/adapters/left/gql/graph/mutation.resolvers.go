package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"hexarch/pkg/adapters/left/gql/graph/generated"
	"hexarch/pkg/adapters/left/gql/graph/model"
)

func (r *mutationResolver) HelloWorld(ctx context.Context, input model.Input) (string, error) {
	return r.App.SayHello(input.Name), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
