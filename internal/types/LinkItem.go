package types

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type LinkItem struct {
	app.Compo

	label string
	href  string
	class string
}

func (l *LinkItem) Render() app.UI {
	return app.A().
		Body(
			app.Text(l.label),
		).
		Href(l.href).
		Class(l.class)
}

func (l *LinkItem) GetLabel() string {
	return l.label
}
func (l *LinkItem) IsSelected(selectedLabel string) bool {
	return l.label == selectedLabel
}
