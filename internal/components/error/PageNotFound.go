package error

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PageNotFound struct {
	app.Compo
}

func (pnf *PageNotFound) Render() app.UI {
	return app.H3().Class("title fit center").Body(
		app.Text("Page Not Found!"),
	)
}
