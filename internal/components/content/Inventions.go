package content

import (
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
		app.H3().
			Body(
				app.Text(inv.Name),
				app.Br(),
			),
		app.H4().Body(
			app.Text("Chapter "+string(inv.Chapter)),
			app.Br(),
			statusUI,
			app.Br(),
		),
		app.Div().Class("list-item_sub").Body(
			app.Range(inv.Photos).Slice(func(i int) app.UI {
				return renderInventionPhoto(&inv.Photos[i])
			}),
		),
		app.Br(),
		app.Div().Class("list-item_description").Body(
			app.Text(inv.Description),
		),
		app.Div().Class("list-item_sub").Body(
			app.Range(inv.Recipe).Slice(func(i int) app.UI {
				return renderInventionRecipeItem(&inv.Recipe[i])
			}),
		),
		app.Hr(),
	)
}

func renderInventionPhoto(p *Photo) app.UI {
	return app.A().Class("list-item_sub_item").Body(
		app.Text(p.Name),
		app.Br(),
	).Href("/photos#" + p.ID)
}

func renderInventionRecipeItem(ri *RecipeItem) app.UI {
	return app.A().Class("list-item_sub_item").Body(
		// app.Text(ri.Item),
		// app.Text(":"+string(ri.Count)),
		app.Br(),
	)
	// Href("/items#" + ri.Item.ID)
}
