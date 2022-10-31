package content

import (
	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Inventions struct {
	app.Compo
}

func (ic *Inventions) Render() app.UI {
	inventionList := GetAppData().Inventions
	return app.Div().Body(
		app.Range(inventionList).Slice(func(i int) app.UI {
			var inv = inventionList[i]
			return renderInvention(inv, GetUserData().HasInvention(inv.ID))
		}),
	)
}

func ToggleHasInvention(iID string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		GetUserData().ToggleInvention(iID)
	}
}

func renderInvention(inv Invention, hasInv bool) app.UI {
	return app.Div().ID(inv.ID).Body(
		app.H3().
			Body(
				app.Text(inv.Name),
				app.Input().Type("checkbox").Checked(hasInv).OnChange(ToggleHasInvention((inv.ID))),
			),
		app.Br(),
		app.Range(inv.Photos).Slice(func(i int) app.UI {
			return renderInventionPhoto(inv.Photos[i])
		}),
		app.Hr(),
	)
}

func renderInventionPhoto(p *Photo) app.UI {
	return app.A().Body(
		app.Text(p.Name),
		app.Br(),
	).Href("/photos#" + p.ID)
}
