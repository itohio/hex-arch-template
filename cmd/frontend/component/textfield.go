package component

import (
	"fmt"
	"log"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

type TextField struct {
	vecty.Core

	ID          string             `vecty:"prop"`
	Label       string             `vecty:"prop"`
	Value       string             `vecty:"prop"`
	Placeholder string             `vecty:"prop"`
	Autofocus   bool               `vecty:"prop"`
	Disabled    bool               `vecty:"prop"`
	NoRipple    bool               `vecty:"prop"`
	OnInput     func(*vecty.Event) `vecty:"prop"`
}

// elem.Input(
// 	vecty.Markup(
// 		vecty.Class("mdc-text-field"),
// 		prop.Placeholder("What's your name?"),
// 		prop.Autofocus(true),
// 		prop.Value(p.name),
// 		event.Input(p.onGreetingInput),
// 	),
// ),

func (cb *TextField) Render() vecty.ComponentOrHTML {
	if cb.ID == "" {
		panic(fmt.Errorf("Form ID must be provided"))
	}
	log.Println("Value is ", cb.Value)
	return elem.Div(
		vecty.Markup(vecty.Class("mdc-form-field")),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-text-field", "mdc-text-field--filled"),
				vecty.MarkupIf(cb.Disabled, vecty.Class("mdc-text-field--disabled")),
			),
			vecty.If(!cb.NoRipple, elem.Div(vecty.Markup(vecty.Class("mdc-text-field__ripple")))),
			elem.Label(
				vecty.Markup(
					vecty.Class("mdc-floating-label--float-above"),
					prop.For(cb.ID),
				),
				vecty.Text(cb.Label),
			),
			elem.Input(
				vecty.Markup(
					prop.ID(cb.ID),
					prop.Type(prop.TypeText),
					vecty.Class("mdc-text-field__input"),
					vecty.Attribute("aria-labelledby", cb.ID),
					prop.Disabled(cb.Disabled),
					prop.Value(cb.Value),
					event.Input(cb.OnInput),
					prop.Placeholder("What's your name?"),
				),
			),
			vecty.If(!cb.NoRipple, elem.Div(vecty.Markup(vecty.Class("mdc-line__ripple")))),
		),
	)

	// return elem.Div(vecty.Markup(vecty.Class("mdc-form-field")),
	// 	elem.Div(vecty.Markup(
	// 		vecty.Class("mdc-text-field"),
	// 		vecty.MarkupIf(cb.Disabled, vecty.Class("mdc-text-field--disabled")),
	// 	),
	// 		elem.Input(vecty.Markup(
	// 			prop.Type(prop.TypeText),
	// 			prop.ID(cb.ID),
	// 			vecty.Class("mdc-text-field__native-control"),
	// 			// TODO(soypat): How to implement data determination.
	// 			vecty.MarkupIf(false, vecty.Attribute("data-indeterminate", "true")),
	// 		)),
	// 		// elem.Div(
	// 		// 	vecty.Markup(vecty.Class("mdc-checkbox__background"), vecty.UnsafeHTML(svgCheckbox)),
	// 		// ),
	// 		elem.Div(vecty.Markup(vecty.MarkupIf(!cb.NoRipple), vecty.Class("mdc-text-field__ripple"))),
	// 	),
	// 	elem.Label(vecty.Markup(prop.For(cb.ID)), cb.Label),
	// )
}
