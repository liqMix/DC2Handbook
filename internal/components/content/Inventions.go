package content

import (
	"strconv"

	"github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Inventions struct {
	app.Compo
}

func (ic *Inventions) Render() app.UI {
	inventionList := GetAppData().Inventions
	ud := GetUserData()
	return app.Div().Body(
		app.Range(inventionList).Slice(func(i int) app.UI {
			var inv = &inventionList[i]
			return renderInvention(inv, ud)
		}),
	)
}

func ToggleHasInvention(iID string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		GetUserData().ToggleInvention(iID)
	}
}

func renderInvention(inv *Invention, userData *data.UserData) app.UI {
	statusUI := userData.GetInventionStatus(inv).
		ToUI(userData.Chapter, inv.Chapter, "🔧").
		OnClick(
			ToggleHasInvention(inv.ID),
		)

	return app.Div().Class("list-item").ID(inv.ID).Body(
		app.Div().Class("list-item_header no-margin-top").
			Body(
				app.Div().Body(
					app.H2().Class("no-margin-top").Body(
						app.Text(inv.Name),
					),
					app.H4().Class("text-small").Body(
						app.Text("Chapter "+string(inv.Chapter)),
					),
					app.H5().Class("clickable").Body(
						statusUI,
					),
				),
				app.Div().Class("list-item_sub_section text-small").Body(
					app.H3().Class("no-margin-top").Body(app.Text("Photos")),
					app.Dl().Class("list-item_sub").Body(
						app.Range(inv.Photos).Slice(func(i int) app.UI {
							photo, err := GetPhoto(inv.Photos[i].ID)
							if err != nil {
								return nil
							}
							return renderInventionPhoto(&photo)
						}),
					),
					app.H3().Class("no-margin-top").Body(app.Text("Recipe")),
					app.Dl().Class("list-item_sub").Body(
						app.Range(inv.Recipe).Slice(func(i int) app.UI {
							return renderInventionRecipeItem(inv.Recipe[i])
						}),
					),
				),
			),
		app.Br(),
		app.Div().Class("list-item_description").Body(
			app.Text(inv.Description),
		),
		app.Hr(),
	)
}

func renderInventionPhoto(p *Photo) app.UI {
	class := GetUserData().GetPhotoStatus(p).ToClass()
	return app.Dd().Body(
		app.A().Class("list-item_sub_item clickable "+class).Body(
			app.Text(p.Name),
			app.Br(),
		).Href("/photos#" + p.ID))
}

func renderInventionRecipeItem(ri RecipeItem) app.UI {
	item, err := GetItem(ri.ItemID)
	if err != nil {
		return nil
	}
	return app.Dd().Body(
		app.A().Class("list-item_sub_item clickable").Body(
			app.Text(item.Name+" x "+strconv.Itoa(ri.Count)),
			app.Br(),
		).Href("/items#" + ri.ItemID),
	)
}
