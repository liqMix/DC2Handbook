package content

import (
	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Items struct {
	app.Compo
}

func (pc *Items) Render() app.UI {
	itemList := GetAppData().Items
	return app.Div().Body(
		app.Range(itemList).Slice(func(i int) app.UI {
			return renderItem(itemList[i])
		}),
	)
}

func renderItem(i Item) app.UI {
	return app.Div().Class("list-item").ID(i.ID).Body(
		app.H3().Body(app.Text(i.Name)),
		app.Br(),
		app.Text(i.Value+" Gilda"),
		app.Div().Class("list-item_image_container").Body(
			app.Img().Class("list-item_image").Src(PLACEHOLDER_IMAGE),
		),
		app.Hr(),
	)
}
