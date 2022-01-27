package page

import (
	"context"
	"hexarch/cmd/frontend/component"
	"hexarch/cmd/frontend/gql"
	"log"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/style"
	"github.com/soypat/mdc"
)

type Home struct {
	vecty.Core
	gState   *component.GlobalState
	name     string          `vecty:"prop"`
	greeting component.State `vecty:"prop"`
}

func (p *Home) onGreet(e *vecty.Event) {
	log.Println("Submit ", p.name)
	component.UseEffect(p, time.Second*3, func(ctx context.Context) (interface{}, error) {
		val, err := p.gState.Client.HelloWorld(ctx, gql.Input{Name: p.name})
		if err != nil {
			return nil, err
		}

		p.greeting.Set(val.HelloWorld)
		return val.HelloWorld, nil
	})
}

func (p *Home) onGreetingInput(e *vecty.Event) {
	p.name = e.Target.Get("value").String()
}

func (p *Home) Render() vecty.ComponentOrHTML {
	log.Println("Render Home")
	p.greeting.Component = p

	return elem.Div(
		vecty.If(p.greeting.String() == "", &mdc.Typography{Root: vecty.Text("Enter your name to receive a random greeting!")}),
		vecty.If(p.greeting.String() != "", &mdc.Typography{Root: vecty.Text(p.greeting.String())}),
		elem.Form(
			vecty.Markup(
				style.Margin(style.Px(0)),
				event.Submit(p.onGreet).PreventDefault(),
			),

			&component.TextField{
				ID:          "name",
				Label:       "Name",
				Value:       p.name,
				OnInput:     p.onGreetingInput,
				Autofocus:   true,
				Placeholder: "What's your name?",
			},

			&mdc.Button{
				Label: vecty.Text("Submit"),
			},
		),
	)
}

func NewHome(state *component.GlobalState) vecty.Component {
	return &Home{gState: state}
}
