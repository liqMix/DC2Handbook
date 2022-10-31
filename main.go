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

	loaded  bool
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
	ctx.SetState("loaded", false)
	data.InitAppData(&ctx)
	data.InitUserData(&ctx)
	a.path = utils.GetPath()
	ctx.ObserveState("loaded").
		While(func() bool {
			return !a.loaded
		}).
		OnChange(func() {
			a.loaded = data.GetAppData() != nil
		}).
		Value(&a.loaded)
}

func (a *Application) Render() app.UI {
	return ui.Shell().
		Class(stylePrefix + "app-info background").
		Menu(
			a.Menu.Render(),
		).
		Index(
			a.Submenu.Render(),
		).
		Content(
			app.If(!a.loaded,
				ui.Loader().Loading(true),
			).Else(
				app.Div().Class(stylePrefix+"stack header"),
				a.Content.Render(),
			),
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
		},
	})

	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal(err)
	}
}
