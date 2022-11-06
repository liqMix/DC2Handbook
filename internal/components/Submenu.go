package components

import (
	"sort"
	"strconv"

	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/liqMix/DC2Photobook/internal/utils"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Submenu struct {
	app.Compo

	sortBy         string
	desc           bool
	hideTaken      bool
	selected       string
	subMenuItems   []LinkItem
	selectedStatus Status
}

func (s *Submenu) append(li LinkItem) {
	s.subMenuItems = append(s.subMenuItems, li)
}

func (s *Submenu) RefreshSubItems() {
	path := utils.GetPath()
	appData := GetAppData()
	userData := GetUserData()

	s.subMenuItems = make([]LinkItem, 0)
	hasSelectedStatus := s.selectedStatus != NA && s.selectedStatus != ""
	switch path.Root {
	case "/photos":
		for i := range appData.Photos {
			photo := &appData.Photos[i]
			pStatus := userData.GetPhotoStatus(photo)
			hasP := userData.HasPhoto(photo.ID)
			if !hasSelectedStatus || pStatus == s.selectedStatus {
				if !s.hideTaken || !hasP {
					s.append(photo.ToLinkItem(pStatus))
				}
			}
		}

	case "/inventions":
		for i := range appData.Inventions {
			inv := &appData.Inventions[i]
			iStatus := userData.GetInventionStatus(inv)
			hasI := userData.HasInvention(inv.ID)
			if !hasSelectedStatus || iStatus == s.selectedStatus {
				if !s.hideTaken || !hasI {
					s.append(inv.ToLinkItem(iStatus))
				}
			}
		}

	case "/items":
		for i := range appData.Items {
			item := &appData.Items[i]
			iStatus := userData.GetItemStatus(item)
			if !hasSelectedStatus || iStatus == s.selectedStatus {
				s.append(item.ToLinkItem(iStatus))
			}
		}

	default:
		s.subMenuItems = nil
	}

	if s.subMenuItems != nil {
		sort.Slice(s.subMenuItems, func(x, y int) bool {
			a := s.subMenuItems[x].GetLabel()
			b := s.subMenuItems[y].GetLabel()
			desc := s.desc
			switch s.sortBy {
			case "chapter":
				aC := s.subMenuItems[x].Chapter
				bC := s.subMenuItems[y].Chapter
				if aC == bC {
					desc = false
				} else {
					a = aC
					b = bC
				}
			default:
			}

			if desc {
				return a > b
			}
			return a < b
		})
	}
}

func (s *Submenu) selectSort(newSort string) {
	if s.sortBy == newSort {
		s.desc = !s.desc
	} else {
		s.desc = false
		s.sortBy = newSort
	}
	s.RefreshSubItems()
}

func (s *Submenu) renderHeader() app.UI {
	return app.Header().Body(
		app.Text(utils.GetPath().Title),
	).Class("h2 sub-menu_title")
}

func (s *Submenu) renderMenuItems() app.UI {
	if len(s.subMenuItems) == 0 {
		return app.Dl().Body(
			app.Dd().Class("sub-menu_link-item deemphasize").Body(
				app.Text("Nothing here!"),
			),
		)
	}

	return app.Dl().Body(
		app.Range(s.subMenuItems).Slice(func(i int) app.UI {
			menuItem := &s.subMenuItems[i]
			if menuItem == nil {
				return nil
			}

			return app.Dd().Class().Body(
				s.subMenuItems[i].Render(s.selected),
			).OnClick(s.selectSubMenuItem(*menuItem))
		}),
	)
}

func (s *Submenu) selectSubMenuItem(item LinkItem) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		s.selected = item.GetLabel()
	}
}

func (s *Submenu) renderHideButton() app.UI {
	show := "HIDE"
	if s.hideTaken {
		show = "SHOW"
	}
	return app.Div().Class("sub-menu_select margin-bottom").Body(
		app.Button().Class("button toggled-" + strconv.FormatBool(s.hideTaken)).
			Body(app.Text(show + " TAKEN")).
			OnClick(func(ctx app.Context, e app.Event) {
				s.hideTaken = !s.hideTaken
				s.RefreshSubItems()
			}),
	)
}
func (s *Submenu) renderMenuContainer() app.UI {
	return app.Div().Class("vspace-top").Body(
		app.Div().Class("godoc-index").Body(
			app.Div().Class("manual-nav").Body(
				CreateStatusSelect(s.selectedStatus).
					Class("sub-menu_select").
					OnChange(func(ctx app.Context, e app.Event) {
						status := StatusFromString(ctx.JSSrc().Get("value").String())
						s.desc = false
						s.sortBy = ""
						s.selectedStatus = status
						s.RefreshSubItems()
					}),
				s.renderHideButton(),
				app.Div().Class("margin-bottom").Body(
					app.Dd().Class("sub-menu_link-item").Body(
						app.Div().Body(app.Text("Name")).Class("clickable").
							OnClick(func(ctx app.Context, e app.Event) {
								s.selectSort("")
							}),
						app.Div().Body(app.Text("Chapter")).Class("clickable").
							OnClick(func(ctx app.Context, e app.Event) {
								s.selectSort("chapter")
							}),
					),

					s.renderMenuItems(),
				),
			),
		),
	)
}

func (s *Submenu) Render() app.UI {
	return app.If(s.subMenuItems != nil,
		app.Nav().Class("content").Body(
			app.Div().Class("margin-top").Body(
				s.renderHeader(),
				s.renderMenuContainer(),
			),
		))
}
