package components

import (
	"github.com/liqMix/DC2Photobook/internal/utils"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const stylePrefix string = "goapp-"

type MainMenu struct {
	app.Compo

	selected string
}

func (m *MainMenu) renderOption(path, label string) app.UI {
	className := "link heading fit"
	if path == m.selected {
		className += " vignette"
	}
	return app.A().Href(path).Class(className).Body(
		app.Text(label),
	)
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
		m.renderOption("/", "Home"),
		m.renderOption("/about", "About"),
		m.renderOption("/user", "User"),
	)
}

func (m *MainMenu) renderAppOptions() app.UI {
	return app.Div().Class("vspace-top").Body(
		m.renderOption("/items", "Items"),
		m.renderOption("/photos", "Photos"),
		m.renderOption("/inventions", "Inventions"),
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
	m.selected = utils.GetPath().Root
	return app.Div().Body(
		m.renderTitle(),
		m.renderOptions(),
	).Class("fill unselectable")
}
