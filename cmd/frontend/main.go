package main

//go:generate gqlgenc

import (
	"hexarch/cmd/frontend/auth0"
	"hexarch/cmd/frontend/component"
	"log"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

const (
	title        = "U-Rule"
	motto        = "You are the best. Always have been."
	AuthDomain   = "dev-itohio.eu.auth0.com"
	AuthClientID = "WyLy3saoCf12ArfKrPsGnQlUurVxzqVT"
	AuthAudience = "http://localhost:8080/api"
	AuthBaseUrl  = "http://localhost:3000/"
	AuthScope    = ""
)

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()
	mdc.AddDefaultScripts()

	body := &Body{state: &globalState}
	body.state.GlobalListener = func() {
		vecty.Rerender(body)
	}
	auth, err := auth0.New(
		auth0.WithDomain(AuthDomain),
		auth0.WithClientID(AuthClientID),
		auth0.WithAudience(AuthAudience),
		auth0.WithRedirectUri(AuthBaseUrl),
		auth0.WithScope(AuthScope),
	)
	if err != nil {
		panic(err)
	}
	body.state.Auth = auth
	vecty.RenderBody(body)
}

type Body struct {
	vecty.Core
	state *component.GlobalState `vecty:"prop"`
}

func (b *Body) Render() vecty.ComponentOrHTML {
	log.Println("Body render ", b.state.Auth.Authenticated(), b.state.Auth.Token())
	authPre, authPost := component.NewAuth(b.state, AuthAudience)
	return elem.Body(
		authPre,
		NewApp(b.state, siteMap),
		authPost,
	)
}

func NewApp(state *component.GlobalState, siteMap component.SiteMap) *mdc.SPA {
	log.Println("NewApp")
	spa := &mdc.SPA{
		FullHeightDrawer: false,
		Drawer: &mdc.Leftbar{
			Title:     vecty.Text(title),
			Subtitle:  vecty.Text(motto),
			Variant:   mdc.VariantDismissableLeftbar,
			Dismissed: state.SideBarDismissed,
			List: &mdc.List{
				ID:       "list-111",
				ListElem: mdc.ElementNavigationList,
				Items:    siteMap.SideBar(state),
			},
		},
		Content: siteMap.Content(state),
		Navbar:  component.NewNavbar(state, siteMap),
	}

	return spa
}
