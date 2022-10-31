package content

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (home *Home) Render() app.UI {
	return app.Div().Body()
}
