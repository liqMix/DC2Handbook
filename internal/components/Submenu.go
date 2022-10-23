package components

import (
	"strings"

	. "github.com/liqMix/DC2Photobook/internal/data"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Submenu struct {
	app.Compo

	selected     string
	rootPath     string
	sorter       Sorter
	subMenuItems []LinkItem
}

func (s *Submenu) HandleNavChange() {
	path := app.Window().URL().Path
	path = strings.Replace(path, "/", "", 1)
	var subMenuItems []LinkItem
	if s.rootPath != path {
		s.rootPath = path
		s.selected = ""
	}
	// var paths = strings.Split(path, "/")
	// if s.rootPath != paths[1] {
	// 	s.rootPath = paths[1]
	// 	s.selected = ""
	// }

	switch s.rootPath {
	case "config":
		configItems := [][]string{{"test", "test"}}
		for _, v := range configItems {
			subMenuItems = append(subMenuItems, LinkItemFromStrings(v[0], v[1], s.rootPath))
		}
	case "items":
		items := ApplicationData().Items
		subMenuItems = make([]LinkItem, len(items))
		for i, v := range items {
			subMenuItems[i] = LinkItemFromItem(v, "")
		}
	case "photos":
		items := ApplicationData().Photos
		subMenuItems = make([]LinkItem, len(items))
		for i, v := range items {
			subMenuItems[i] = LinkItemFromPhoto(v, "")
		}
	case "inventions":
		items := ApplicationData().Inventions
		subMenuItems = make([]LinkItem, len(items))
		for i, v := range items {
			subMenuItems[i] = LinkItemFromInvention(v, "")
		}
	default:
	}

	s.subMenuItems = subMenuItems
}

func (s *Submenu) renderHeader() app.UI {
	return app.Header().Body(
		app.Text("Index"),
		s.sorter.Render(),
	).Class("h2 vspace-top")
}

func (s *Submenu) renderMenuItems() app.UI {
	return app.Dl().Body(
		app.Range(s.subMenuItems).Slice(func(i int) app.UI {
			className := ""
			if s.subMenuItems[i].label == s.selected {
				className = "vignette"
			}
			return app.Dd().Class(className).Body(
				s.subMenuItems[i].Render(),
			).OnClick(s.selectSubMenuItem(s.subMenuItems[i]))
		}),
	)
}

func (s *Submenu) selectSubMenuItem(item LinkItem) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		s.selected = item.label
	}
}
func (s *Submenu) renderMenuContainer() app.UI {
	return app.Div().Class("vspace-top").Body(
		app.Div().Class("godoc-index").Body(
			app.Div().Class("manual-nav").Body(
				s.renderMenuItems(),
			),
		),
	)
}

func (s *Submenu) Render() app.UI {
	return app.If(s.subMenuItems != nil && len(s.subMenuItems) > 0,
		app.Nav().Class("header-out content unselectable").Body(
			app.Div().Class("hspace-out").Body(
				s.renderHeader(),
				s.renderMenuContainer(),
			),
		))
}
