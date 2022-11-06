package main

import (
	"log"
	"net/http"

	. "github.com/liqMix/DC2Photobook/internal/components"
	"github.com/liqMix/DC2Photobook/internal/data"
	"github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const stylePrefix string = "goapp-"

type Application struct {
	app.Compo

	loaded       bool
	currentFocus string
	rerender     bool
	Menu         MainMenu
	Submenu      Submenu
	Content      Content
}

// func (a *Application) focusElement(ctx app.Context) {
// 	a.currentFocus = a.path.Fragment
// 	link := app.Window().GetElementByID(a.currentFocus)
// 	if link.Truthy() {
// 		link.Set("autofocus", true)
// 	}
// }
func (a *Application) unfocusElement(ctx app.Context, e app.Event) {
	app.Log("onclick")
	// link := app.Window().GetElementByID(a.currentFocus)
	// if link.Truthy() {
	// 	link.Set("className", "")
	// }
}

func (a *Application) OnMount(ctx app.Context) {
	ctx.SetState(string(types.E_LOADED), false)
	data.InitAppData(&ctx)
	data.InitUserData(&ctx)
	app.Window().AddEventListener("click", a.unfocusElement)

	// When data has finished loading from JSON
	ctx.ObserveState(string(types.E_LOADED)).
		While(func() bool {
			return !a.loaded
		}).
		OnChange(func() {
			a.loaded = data.GetAppData() != nil
		}).
		Value(&a.loaded)

	// When user is saved
	ctx.ObserveState(types.E_USERUPDATE).
		OnChange(func() {
			a.Submenu.RefreshSubItems()
		}).
		Value(&a.rerender)
}

func (a *Application) OnNav(ctx app.Context) {
	a.Submenu.RefreshSubItems()
	// a.focusElement(ctx)
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
			"/web/css/app.css",
			"/web/css/docs.css",
		},
	})

	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal(err)
	}
}
