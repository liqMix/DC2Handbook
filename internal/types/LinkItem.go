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

func (l LinkItem) Render() app.UI {
	return app.A().
		Body(
			app.Div().Body(app.Text(l.label)),
			app.Div().Body(app.If(l.Chapter != "", app.Text(l.Chapter))),
		).
		Href(l.href).
		Class(l.class + " sub-menu_link-item")
}

func (l LinkItem) GetLabel() string {
	return l.label
}
func (l LinkItem) IsSelected(selectedLabel string) bool {
	return l.label == selectedLabel
}
