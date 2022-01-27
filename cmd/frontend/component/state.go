package component

import (
	"fmt"
	"net/http"

	"hexarch/cmd/frontend/gql"

	"github.com/soypat/mdc"
)

type NavigationState struct {
	SideBarDismissed bool
}

type AuthState interface {
	Error() string
	Token() string
	Authenticated() bool
	Scopes() []string
	LoginUrl(redirect string) string
	SignupUrl(redirect string) string
	LogoutUrl(redirect string) string
	Handle(queryStr, hashStr string) (string, error)
}

type API struct {
	token  string
	Client *gql.Client
}

type GlobalState struct {
	NavigationState
	API
	Auth           AuthState
	GlobalListener func()
	NavBar         *mdc.Navbar
}

func (a *API) Init(path, token string) {
	if token == a.token && a.Client != nil {
		return
	}
	if token != "" {
		a.Client = gql.NewClient(http.DefaultClient, path, func(req *http.Request) {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		})
	} else {
		a.Client = gql.NewClient(http.DefaultClient, path)
	}
	a.token = token
}
