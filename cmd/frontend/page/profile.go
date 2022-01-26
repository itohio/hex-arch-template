package page

import (
	"context"
	"fmt"
	"hexarch/cmd/frontend/component"
	"log"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

type Profile struct {
	vecty.Core
	gState *component.GlobalState `vecty:"prop"`
}

func (p *Profile) Render() vecty.ComponentOrHTML {
	log.Println("Render Profile")

	greetings := vecty.List{}
	res, err := p.gState.Client.GetGreetings(context.Background())
	if err != nil {
		greetings = append(greetings, &mdc.ListItem{Label: vecty.Text(fmt.Sprint(err))})
	} else if res == nil {
		greetings = append(greetings, &mdc.ListItem{Label: vecty.Text("Result is nil")})
	} else {
		greetings = append(greetings, &mdc.ListItem{Label: vecty.Text("Available greetings:")})
		for _, g := range res.Greetings {
			greetings = append(greetings, &mdc.ListItem{Label: vecty.Text(g)})
		}
	}

	return elem.Div(
		&mdc.Typography{Root: vecty.Text("Profile stuff.")},
		greetings,
	)
}

func NewProfile(state *component.GlobalState) vecty.Component {
	log.Println("NewProfile")
	return &Profile{gState: state}
}
