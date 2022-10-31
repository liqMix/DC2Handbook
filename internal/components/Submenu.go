package components

import (
	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/liqMix/DC2Photobook/internal/utils"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Submenu struct {
	app.Compo

	selected     string
	sorter       Sorter
	subMenuItems []*LinkItem
}

func (s *Submenu) HandleNavChange() {
	var subMenuItems []*LinkItem
	path := utils.GetPath()
	appData := GetAppData()
	var list []IBase
	switch path.Root {
	case "photos":
		list = make([]IBase, len(appData.Photos))
		for i := range appData.Photos {
			list[i] = &appData.Photos[i]
		}

	case "inventions":
		list = make([]IBase, len(appData.Inventions))
		for i := range appData.Inventions {
			list[i] = &appData.Inventions[i]
		}

	case "items":
		list = make([]IBase, len(appData.Items))
		for i := range appData.Items {
			list[i] = &appData.Items[i]
		}

	default:
	}
	subMenuItems = make([]*LinkItem, len(list))
	for i, v := range list {
		subMenuItems[i] = v.ToLinkItem()
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
			if s.subMenuItems[i].IsSelected(s.selected) {
				className = "vignette"
			}
			return app.Dd().Class(className).Body(
				s.subMenuItems[i].Render(),
			).OnClick(s.selectSubMenuItem(s.subMenuItems[i]))
		}),
	)
}

func (s *Submenu) selectSubMenuItem(item *LinkItem) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		s.selected = item.GetLabel()
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
