package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"hexarch/pkg/adapters/left/gql/graph/generated"
	"hexarch/pkg/adapters/left/gql/graph/model"
)

func (r *queryResolver) Setup(ctx context.Context) (*model.Setup, error) {
	return &model.Setup{
		Domain:   r.Cfg.Auth.Domain,
		Audience: r.Cfg.Auth.Audience,
		ClientID: r.Cfg.Auth.ClientID,
	}, nil
}

func (r *queryResolver) Greetings(ctx context.Context) ([]string, error) {
	return r.Db.GetGreetings(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
