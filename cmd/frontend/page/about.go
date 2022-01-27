package page

import (
	"hexarch/cmd/frontend/component"
	"log"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

type About struct {
	vecty.Core
}

func (p *About) Render() vecty.ComponentOrHTML {
	log.Println("Render About")
	return elem.Div(
		&mdc.Typography{Root: vecty.Text("About stuf.")},
	)
}

func NewAbout(state *component.GlobalState) vecty.Component {
	return &About{}
}
