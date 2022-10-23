package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Sorter struct {
	app.Compo

	filter string
}

func (s *Sorter) Render() app.UI {
	return app.Select()
}
