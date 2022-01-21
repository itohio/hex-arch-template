package gql

import (
	"context"
	"fmt"
	"hexarch/pkg/adapters/left/auth"

	"github.com/99designs/gqlgen/graphql"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	_, err = auth.GetAuthorization(ctx)
	if err != nil {
		return
	}

	return next(ctx)
}
func HasScope(ctx context.Context, obj interface{}, next graphql.Resolver, scopes []*string) (res interface{}, err error) {
	token, err := auth.GetAuthorization(ctx)
	if err != nil {
		return
	}

	claimScopes, err := auth.GetScopes(token)

	// Doesn't need to be optimized as the number of scopes
	// should be kept to a minimum anyway.
	matches := 0
	for _, scope := range scopes {
		for _, claimScope := range claimScopes {
			if *scope == claimScope {
				matches++
				break
			}
		}
	}
	if matches != len(scopes) {
		return nil, fmt.Errorf("scope not found")
	}

	return next(ctx)
}
