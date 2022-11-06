package components

import (
	"github.com/liqMix/DC2Photobook/internal/utils"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type MainMenu struct {
	app.Compo

	selected string
}

func (m *MainMenu) renderOption(path, label string) app.UI {
	className := "link"
	if path == m.selected {
		className += " vignette"
	}
	return app.A().Href(path).Class(className).Body(
		app.Text(label),
	)
}

func (m *MainMenu) renderTitle() app.UI {
	return app.Div().Body(
		app.Img().Src("/web/img/logo.png"),
	).Class("main-menu_title")
}

func (m *MainMenu) renderHomeOptions() app.UI {
	return app.Div().Class("main-menu_section").Body(
		m.renderOption("/", "Home"),
		m.renderOption("/about", "About"),
		m.renderOption("/user", "User"),
	)
}

func (m *MainMenu) renderAppOptions() app.UI {
	return app.Div().Class("main-menu_section").Body(
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
		).Class(),
	).Class("content main-menu_options")
}

func (m *MainMenu) Render() app.UI {
	m.selected = utils.GetPath().Root
	return app.Div().Body(
		m.renderTitle(),
		m.renderOptions(),
		app.Div().Class("main-menu_border"),
	).Class("main-menu")
}
