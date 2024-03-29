package content

import (
	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Photos struct {
	app.Compo
}

func (pc *Photos) Render() app.UI {
	photoList := GetAppData().Photos
	userData := GetUserData()
	return app.Div().Body(
		app.Range(photoList).Slice(func(i int) app.UI {
			return renderPhoto(&photoList[i], userData)
		}),
	)
}

func ToggleHasPhoto(photoID string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		GetUserData().TogglePhoto(photoID)
	}
}

func renderPhoto(p *Photo, userData *UserData) app.UI {
	statusUI := userData.GetPhotoStatus(p).
		ToUI(userData.Chapter, p.Chapter, "📸", true).
		OnClick(
			ToggleHasPhoto(p.ID),
		)
	return app.Div().Class("list-item").ID(p.ID).Body(
		app.Hr(),
		app.H2().Class("list-item_header no-margin-top").Body(
			app.Div().Body(
				app.Text(p.Name),
				app.H5().Body(
					app.Text("Chapter "+p.Chapter),
					app.Br(),
					app.Div().Class("list-item_status").Body(statusUI),
					app.Br(),
				)),
			app.Div().Class("list-item_image_container").Body(
				app.Img().Class("list-item_image").Src(PLACEHOLDER_IMAGE),
			),
			app.Div().Class("list-item_header_icon").Body(
				app.If(
					p.IsScoop,
					app.Div().Class("icon").Body(app.Text("🥄")).Title("Scoop"),
				),
				app.If(
					p.Missable,
					app.Div().Class("icon").Body(app.Text("🙈")).Title("Missable"),
				),
			),
		),
		app.Br(),
		app.Text(p.Location),
		app.Img().Src(p.ImageURL),
	)
}
