package components

import (
	"strings"

	"github.com/liqMix/DC2Photobook/internal/components/content"
	err "github.com/liqMix/DC2Photobook/internal/components/error"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Content struct {
	app.Compo

	rootPath string
	subPath  string
}

func (c *Content) HandleNavChange() {
	path := app.Window().URL().Path
	c.rootPath = strings.Replace(path, "/", "", 1)

	// var paths []string = strings.Split(path, "/")
	// if (len(paths) > 1) {
	// 	tempPath := strings.Split(paths[1], "#")
	// 	app.Log(tempPath)
	// 	c.rootPath = tempPath[0]
	// 	if len(tempPath) > 1 {
	// 		c.hashPath = tempPath[1]
	// 	}
	// 	if (len(paths) > 2) {
	// 		c.subPath = paths[2]
	// 	}
	// }
}

func (c *Content) Render() app.UI {
	app.Log("rendering content page")
	app.Log(c.rootPath)
	app.Log(c.subPath)
	return app.Div().Class("content").Body(
		app.H1().Class("title fit center").Body(
			app.Text(c.rootPath+"\n"),
			app.Text(c.subPath+"\n"),
		),
		app.Article().Body(
			app.If(c.rootPath == "photos",
				&content.PhotosContent{},
			).ElseIf(c.rootPath == "items",
				&content.ItemsContent{},
			).ElseIf(c.rootPath == "inventions",
				&content.InventionsContent{},
			).ElseIf(c.rootPath == "",
				&content.InventionsContent{},
			).Else(
				&err.PageNotFound{},
			),
		),
	)
}
