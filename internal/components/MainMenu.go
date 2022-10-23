package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const stylePrefix string = "goapp-"

type MainMenu struct {
	app.Compo
}

func (m *MainMenu) renderTitle() app.UI {
	return app.Div().Body(
		app.Header().Body(
			app.A().Body(
				app.Text("DC2 Handbook"),
			).Href("/").Class("app-title"),
		),
	).Class(stylePrefix + "stack header hspace-out")
}

func (m *MainMenu) renderHomeOptions() app.UI {
	return app.Div().Class("vspace-top").Body(
		app.A().Href("/").Class("link heading fit").Body(
			app.Text("Home"),
		),
		app.A().Href("/about").Class("link heading fit").Body(
			app.Text("About"),
		),
		app.A().Href("/config").Class("link heading fit").Body(
			app.Text("Config"),
		),
	)
}

func (m *MainMenu) renderAppOptions() app.UI {
	return app.Div().Class("vspace-top").Body(
		app.A().Href("/items").Class("link heading fit").Body(
			app.Text("Items"),
		),
		app.A().Href("/photos").Class("link heading fit").Body(
			app.Text("Photos"),
		),
		app.A().Href("/inventions").Class("link heading fit").Body(
			app.Text("Inventions"),
		),
	)
}

func (m *MainMenu) renderOptions() app.UI {
	return app.Nav().Body(
		app.Div().Body(
			m.renderHomeOptions(),
			m.renderAppOptions(),
		).Class("hspace-out"),
	).Class("content")
}

func (m *MainMenu) Render() app.UI {
	return app.Div().Body(
		m.renderTitle(),
		m.renderOptions(),
	).Class("fill unselectable")
}
