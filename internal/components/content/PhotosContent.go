package content

import (
	"strconv"

	. "github.com/liqMix/DC2Photobook/internal/data"
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PhotosContent struct {
	app.Compo
}

func (pc *PhotosContent) Render() app.UI {
	photoList := ApplicationData().Photos
	userData := ApplicationData().UserData
	return app.Div().Body(
		app.Range(photoList).Slice(func(i int) app.UI {
			var photo = photoList[i]
			return renderPhoto(photo, userData.HasPhoto(photo.ID))
		}),
	)
}

func ToggleHasPhoto(photoID string) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		ApplicationData().UserData.TogglePhoto(photoID)
	}
}

func renderPhoto(p Photo, hasPhoto bool) app.UI {
	return app.Div().ID(p.ID).Body(
		app.H3().Body(
			app.Text(p.Name),
			app.Input().Type("checkbox").Checked(hasPhoto).OnChange(ToggleHasPhoto(p.ID)),
		),
		app.Br(),
		app.H4().Body(app.Text("Chapter "+strconv.Itoa(int(p.Chapter)))),
		app.Br(),
		app.Text(p.Location),
		app.Br(),
		app.Hr(),
	)
}
