package components

import (
	"github.com/liqMix/DC2Photobook/internal/components/content"
	err "github.com/liqMix/DC2Photobook/internal/components/error"
	"github.com/liqMix/DC2Photobook/internal/utils"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Content struct {
	app.Compo
}

func (c *Content) Render() app.UI {
	path := utils.GetPath()
	return app.Div().Class("content").Body(
		app.H1().Class("title fit center").Body(
			app.Text(path.Root+"\n"),
			app.Text(path.Sub+"\n"),
		),
		app.Article().Body(
			app.If(path.Root == "photos",
				&content.PhotosContent{},
			).ElseIf(path.Root == "items",
				&content.ItemsContent{},
			).ElseIf(path.Root == "inventions",
				&content.InventionsContent{},
			).ElseIf(path.Root == "",
				&content.InventionsContent{},
			).Else(
				&err.PageNotFound{},
			),
		),
	)
}
