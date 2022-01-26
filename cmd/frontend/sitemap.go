package main

import (
	"hexarch/cmd/frontend/component"
	"hexarch/cmd/frontend/page"

	router "marwan.io/vecty-router"
)

var siteMap = []component.PageLink{
	{
		Label:        "Home",
		Route:        "/",
		Content:      page.NewHome,
		RouteOptions: router.NewRouteOpts{ExactMatch: true},
	},
	{
		Label:        "Profile",
		Route:        "/profile",
		Content:      page.NewProfile,
		Protected:    true,
		Sidebar:      true,
		RouteOptions: router.NewRouteOpts{ExactMatch: true},
	},
	{
		Label:        "About",
		Route:        "/about",
		Sidebar:      true,
		Content:      page.NewAbout,
		RouteOptions: router.NewRouteOpts{ExactMatch: true},
	},
	{
		Label:        "Contact",
		Route:        "/contact",
		Sidebar:      true,
		Content:      page.NewContact,
		RouteOptions: router.NewRouteOpts{ExactMatch: true},
	},
}
