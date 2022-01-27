package page

import (
	"context"
	"fmt"
	"hexarch/cmd/frontend/component"
	"hexarch/cmd/frontend/gql"
	"log"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/soypat/mdc"
)

type Profile struct {
	vecty.Core
	gState  *component.GlobalState `vecty:"prop"`
	watcher *component.Watcher
	res     *gql.GetGreetings
	err     error
}

func (p *Profile) Render() vecty.ComponentOrHTML {
	log.Println("Render Profile")

	p.watcher.Watch(1)

	greetings := vecty.List{}
	res := p.res
	if p.err != nil {
		greetings = append(greetings, &mdc.ListItem{Label: vecty.Text(fmt.Sprint(p.err))})
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
	ret := &Profile{
		gState: state,
	}
	ret.watcher = component.UseWatcher(
		ret,
		time.Second*3,
		func(ctx context.Context) error {
			log.Println("Request greetings")
			ret.res, ret.err = state.Client.GetGreetings(ctx)
			return nil
		},
		1,
	)

	return ret
}
