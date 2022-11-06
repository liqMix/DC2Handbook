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
			var photo = photoList[i]
			return renderPhoto(&photo, userData)
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
		ToUI(userData.Chapter, p.Chapter, "📸").
		OnClick(
			ToggleHasPhoto(p.ID),
		)
	return app.Div().Class("list-item").ID(p.ID).Body(
		app.Hr(),
		app.H3().Class("no-margin-top").Body(
			app.Text(p.Name),
		),
		app.H4().Body(
			app.Text("Chapter "+p.Chapter),
			app.Br(),
			app.Div().Class("clickable").Body(statusUI),
			app.Br(),
		),
		app.Br(),
		app.Text(p.Location),
		app.Img().Src(p.ImageURL),
	)
}
