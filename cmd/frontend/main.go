package main

//go:generate gqlgenc

import (
	"hexarch/cmd/frontend/component"
	"log"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

const (
	title = "U-Rule"
	motto = "You are the best. Always have been."
)

func main() {
	mdc.SetDefaultViewport()
	mdc.AddDefaultStyles()
	mdc.AddDefaultScripts()

	body := &Body{}
	body.state.GlobalListener = func() {
		vecty.Rerender(body)
	}
	body.state.API.Init("http://localhost:8080/api", "")
	vecty.RenderBody(body)
}

type Body struct {
	vecty.Core
	state component.GlobalState `vecty:"prop"`
}

func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(NewApp(&b.state, siteMap))
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
