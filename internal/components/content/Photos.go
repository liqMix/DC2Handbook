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
	userHasPhoto := userData.HasPhoto(p.ID)
	statusUI := userData.GetPhotoStatus(p).ToUI(userData.Chapter, p.Chapter)

	return app.Div().ID(p.ID).Body(
		app.H3().Body(
			app.Text(p.Name),
			app.Input().Type("checkbox").Checked(userHasPhoto).OnChange(ToggleHasPhoto(p.ID)),
		),
		app.H4().Body(
			app.Text("Chapter "+p.Chapter),
			statusUI,
		),
		app.Br(),
		app.Text(p.Location),
		app.Img().Src(p.ImageURL),
		app.Hr(),
	)
}
