package page

import (
	"hexarch/cmd/frontend/component"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

type Error struct {
	vecty.Core
	state *component.GlobalState
}

func (p *Error) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&mdc.Typography{
			Root: elem.Heading1(
				vecty.Text("An error ocurred"),
			),
		},
		vecty.Text(p.state.Auth.Error()),
	)
}

func NewError(state *component.GlobalState) vecty.Component {
	return &Error{state: state}
}
