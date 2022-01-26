package component

import (
	"net/http"

	"hexarch/cmd/frontend/gql"
)

type NavigationState struct {
	SideBarDismissed bool
}

type AuthState struct {
	Token         string
	Authenticated bool
	Scope         string
}

type API struct {
	Client *gql.Client
}

type GlobalState struct {
	NavigationState
	AuthState
	API
	GlobalListener func()
}

func (a *API) Init(path, token string) {
	if token != "" {
		// TODO
		a.Client = gql.NewClient(http.DefaultClient, path)
	} else {
		a.Client = gql.NewClient(http.DefaultClient, path)
	}
}
