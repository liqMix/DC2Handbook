package content

import (
	"github.com/liqMix/DC2Photobook/internal/data"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type User struct {
	app.Compo
}

func handleChapterUpdate() app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		userData := data.GetUserData()
		newChapter := ctx.JSSrc().Get("value").String()
		userData.SetChapter(newChapter)
	}
}

func (user *User) renderChapterSettings(userData *data.UserData) app.UI {
	currentChapter, err := data.GetChapter(userData.Chapter)
	chapters := data.GetAppData().Chapters
	if err != nil {
		currentChapter = chapters[0]
	}
	label := "Chapter " + currentChapter.ID + ": " + currentChapter.Name
	return app.Div().Body(
		app.H1().Text(label),
		app.Br(),
		app.Select().
			Body(
				app.Range(chapters).Slice(func(i int) app.UI {
					item := chapters[i]
					item.Name = label
					return item.ToOption(userData.Chapter)
				}),
			).
			OnChange(handleChapterUpdate()),
	)
}

func (user *User) renderPhotos(ud *data.UserData) app.UI {
	return app.Div().Body(
		app.H1().Text("Photos"),
		app.Dl().Body(
			app.Range(ud.Photos).Map(func(id string) app.UI {
				photo, err := data.GetPhoto(id)
				if err != nil || !ud.Photos[id] {
					return nil
				}
				return app.Dd().Body(
					app.Text(photo.Name),
					app.Div().Class("clickable").Body(
						app.Text("❌"),
					).OnClick(func(ctx app.Context, e app.Event) {
						ud.TogglePhoto(id)
					}),
				)
			}),
		),
	)
}

func (user *User) renderInventions(ud *data.UserData) app.UI {
	return app.Div().Body(
		app.H1().Text("Inventions"),
		app.Dl().Body(
			app.Range(ud.Inventions).Map(func(id string) app.UI {
				inv, err := data.GetInvention(id)
				if err != nil || !ud.Inventions[id] {
					return nil
				}
				return app.Dd().Body(
					app.Text(inv.Name),
					app.Div().Class("clickable").Body(
						app.Text("❌"),
					).OnClick(func(ctx app.Context, e app.Event) {
						ud.ToggleInvention(id)
					}),
				)
			}),
		),
	)
}

func (user *User) Render() app.UI {
	userData := data.GetUserData()
	return app.Div().Body(
		user.renderChapterSettings(userData),
		user.renderPhotos(userData),
		user.renderInventions(userData),
	)
}
