package component

import (
	"github.com/hexops/vecty"
	"github.com/soypat/mdc"
	router "marwan.io/vecty-router"
)

type PageLink struct {
	Route        string
	Label        string
	Description  string
	Content      func(*GlobalState) vecty.Component
	Disabled     func(*GlobalState) bool
	Options      router.LinkOptions
	RouteOptions router.NewRouteOpts
	Protected    bool
	Sidebar      bool
	Scope        string
	Children     []PageLink
}

type SiteMap []PageLink

func (s SiteMap) TopMenu(state *GlobalState) vecty.List {
	mainMenu := make(vecty.List, 0, len(s))

	for _, item := range s[1:] {
		if item.Protected && !state.Authenticated {
			continue
		}
		mainMenu = append(mainMenu, &mdc.ListItem{
			Label: router.Link(
				item.Route,
				item.Label,
				router.LinkOptions{ID: item.Options.ID, Class: "mdc-list-item"},
			),
		})
	}

	return mainMenu
}

func (s SiteMap) SideBar(state *GlobalState) vecty.List {
	sideBar := make(vecty.List, 0, len(s))

	// TODO: nested

	for _, item := range s[1:] {
		if !item.Sidebar {
			continue
		}
		if item.Protected && !state.Authenticated {
			continue
		}
		// TODO: match scopes
		sideBar = append(sideBar, &mdc.ListItem{
			Label: router.Link(
				item.Route,
				item.Label,
				router.LinkOptions{ID: item.Options.ID, Class: "mdc-list-item"},
			),
			ListItemElem: mdc.ElementAnchorListItem,
			Href:         item.Route,
		})
	}
	return sideBar
}

func (s SiteMap) Content(state *GlobalState) vecty.List {
	content := make(vecty.List, 0, len(s))
	for _, item := range s {
		if item.Content == nil || (item.Disabled != nil && item.Disabled(state)) {
			continue
		}
		content = append(content, router.NewRoute(item.Route, item.Content(state), item.RouteOptions))
	}
	// TODO: nested routes

	content = append(content, router.NotFoundHandler(&mdc.Typography{Root: vecty.Text("404"), Style: mdc.Headline1}))
	return content
}
