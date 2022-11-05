package components

import (
	"strings"

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
	rootTitle := strings.Replace(path.Root, "/", "", 1)
	return app.Div().Class("content").Body(
		app.H1().Class("title fit center").Body(
			app.Text(rootTitle+"\n"),
			app.Text(path.Sub+"\n"),
		),
		app.Article().Body(
			app.If(path.Root == "/photos",
				&content.Photos{},
			).ElseIf(path.Root == "/items",
				&content.Items{},
			).ElseIf(path.Root == "/inventions",
				&content.Inventions{},
			).ElseIf(path.Root == "/user",
				&content.User{},
			).ElseIf(path.Root == "/",
				&content.Home{},
			).Else(
				&err.PageNotFound{},
			),
		),
	)
}
