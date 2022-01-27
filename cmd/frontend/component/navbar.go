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

	if !state.Auth.Authenticated() {
		loginBtns = vecty.List{
			&mdc.ListItem{Label: vecty.Text("Login"), ListItemElem: mdc.ElementAnchorListItem, Href: state.Auth.LoginUrl(getpathname())},
			&mdc.ListItem{Label: vecty.Text("Signup"), ListItemElem: mdc.ElementAnchorListItem, Href: state.Auth.SignupUrl(getpathname())},
		}
	} else {
		loginBtns = vecty.List{
			&mdc.ListItem{Label: vecty.Text("Logout"), ListItemElem: mdc.ElementAnchorListItem, Href: state.Auth.LogoutUrl("#logout")},
		}
	}

	state.NavBar = &mdc.Navbar{
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
	return state.NavBar
}
