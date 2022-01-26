package page

import (
	"hexarch/cmd/frontend/component"
	"log"

	"github.com/hexops/vecty"
	"github.com/soypat/mdc"
)

type Contact struct {
	vecty.Core
}

func (p *Contact) Render() vecty.ComponentOrHTML {
	log.Println("Render Contact")
	return &mdc.Typography{Root: vecty.Text("About Contact.")}
}

func NewContact(state *component.GlobalState) vecty.Component {
	log.Println("NewContact")
	return &Contact{}
}
