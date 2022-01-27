package main

import (
	"hexarch/cmd/frontend/component"
	"hexarch/cmd/frontend/page"

	router "marwan.io/vecty-router"
)

var globalState component.GlobalState

var siteMap = []component.PageLink{
	{
		Label:          "Home",
		Route:          "/",
		Content:        page.NewHome(&globalState),
		ContentBuilder: page.NewHome,
		RouteOptions:   router.NewRouteOpts{ExactMatch: true},
	},
	{
		Label:          "Profile",
		Route:          "/profile",
		Content:        page.NewProfile(&globalState),
		ContentBuilder: page.NewProfile,
		Protected:      true,
		Sidebar:        true,
		RouteOptions:   router.NewRouteOpts{ExactMatch: true},
	},
	{
		Label:          "About",
		Route:          "/about",
		Sidebar:        true,
		Content:        page.NewAbout(&globalState),
		ContentBuilder: page.NewAbout,
		RouteOptions:   router.NewRouteOpts{ExactMatch: true},
	},
	{
		Label:          "Contact",
		Route:          "/contact",
		Sidebar:        true,
		Content:        page.NewContact(&globalState),
		ContentBuilder: page.NewContact,
		RouteOptions:   router.NewRouteOpts{ExactMatch: true},
	},
	{
		Label:          "",
		Route:          "/error",
		Content:        page.NewError(&globalState),
		ContentBuilder: page.NewError,
		RouteOptions:   router.NewRouteOpts{},
	},
}
