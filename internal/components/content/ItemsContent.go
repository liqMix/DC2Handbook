package content

import (
	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ItemsContent struct {
	app.Compo
}

func (pc *ItemsContent) Render() app.UI {
	itemList := ApplicationData().Items
	return app.Div().Body(
		app.Range(itemList).Slice(func(i int) app.UI {
			return renderItem(itemList[i])
		}),
	)
}

func renderItem(i Item) app.UI {
	return app.Div().ID(i.ID).Body(
		app.H3().Body(app.Text(i.Name)),
		app.Br(),
		app.Text(i.Value+" Gilda"),
		app.Hr(),
	)
}
