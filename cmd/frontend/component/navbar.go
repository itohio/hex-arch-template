package component

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/event"
	"github.com/soypat/mdc"
	"github.com/soypat/mdc/icons"
	router "marwan.io/vecty-router"
)

func NewNavbar(state *GlobalState, siteMap SiteMap) *mdc.Navbar {
	dehaze := &mdc.Button{
		Icon: icons.Dehaze,
		Listeners: []*vecty.EventListener{event.Click(func(e *vecty.Event) {
			state.SideBarDismissed = !state.SideBarDismissed
			state.GlobalListener()
		})},
	}

	loginBtns := vecty.List{}

	if !state.Authenticated {
		loginBtns = vecty.List{
			&mdc.ListItem{Label: vecty.Text("Login"), ListItemElem: mdc.ElementAnchorListItem, Href: "/login"},
			&mdc.ListItem{Label: vecty.Text("Signup"), ListItemElem: mdc.ElementAnchorListItem, Href: "/signin"},
		}
	} else {
		loginBtns = vecty.List{
			&mdc.ListItem{Label: vecty.Text("Logout"), ListItemElem: mdc.ElementAnchorListItem, Href: "/logout"},
		}
	}

	return &mdc.Navbar{
		SectionStart: vecty.List{
			dehaze,
			&mdc.Typography{
				Root: router.Link(
					siteMap[0].Route,
					siteMap[0].Label,
					router.LinkOptions{Class: "mdc-list-item"},
				),
				Style: mdc.Headline6,
			},
		},
		SectionCenter: siteMap.TopMenu(state),
		SectionEnd:    loginBtns,
	}
}
