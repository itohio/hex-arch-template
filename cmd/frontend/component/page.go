package component

import (
	"github.com/hexops/vecty"
)

type Page struct {
	vecty.Core
	state   *GlobalState
	builder func(state *GlobalState) vecty.Component
}

func (p *Page) Render() vecty.ComponentOrHTML {
	// FIXME: cause the login/signin/logout to update
	// vecty.Rerender(p.state.NavBar.SectionEnd)
	return p.builder(p.state)
}
