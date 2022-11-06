package types

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type LinkItem struct {
	app.Compo

	Chapter string

	label string
	href  string
	class string
}

func (l LinkItem) Render(selected string) app.UI {
	class := " sub-menu_link-item clickable"
	if l.IsSelected(selected) {
		class += " vignette"
	}
	return app.A().
		Body(
			app.Div().Body(app.Text(l.label)),
			app.Div().Body(app.If(l.Chapter != "", app.Text(l.Chapter))),
		).
		Href(l.href).
		Class(l.class + class)
}

func (l LinkItem) GetLabel() string {
	return l.label
}
func (l LinkItem) IsSelected(selectedLabel string) bool {
	return l.label == selectedLabel
}
