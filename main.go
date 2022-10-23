package main

import (
	"log"
	"net/http"

	. "github.com/liqMix/DC2Photobook/internal/components"
	"github.com/liqMix/DC2Photobook/internal/data"
	"github.com/liqMix/DC2Photobook/internal/utils"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const stylePrefix string = "goapp-"

type Application struct {
	app.Compo

	path    *utils.Path
	Menu    MainMenu
	Submenu Submenu
	Content Content
}

func (a *Application) OnNav(ctx app.Context) {
	a.path = utils.GetPath()
	a.Submenu.HandleNavChange()
}

func (a *Application) OnMount(ctx app.Context) {
	data.InitAppData(&ctx)
	a.path = utils.GetPath()
}

func (a *Application) Render() app.UI {
	return ui.Shell().
		Class(stylePrefix+"app-info background").
		Menu(
			a.Menu.Render(),
		).
		Index(
			a.Submenu.Render(),
		).
		Content(
			app.Div().Class(stylePrefix+"stack header"),
			a.Content.Render(),
		)
}

func main() {
	// Route all pages to Application
	app.RouteWithRegexp("/*", &Application{})

	// app.RouteWithRegexp("/items*", &Content{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "DC2 Handbook",
		Title:       "DC2 Handbook",
		Description: "A tool for managing Dark Cloud 2 photos and recipes",
		Icon: app.Icon{
			Default: "/web/img/logo.png",
		},
		Styles: []string{
			"/web/css/docs.css",
			"/web/css/app.css",
		},
	})

	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal(err)
	}
}
