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
		app.Log("setting chapter")
		app.Log(newChapter)
		userData.SetChapter(newChapter)
	}
}

func (user *User) renderChapterSettings(userData *data.UserData) app.UI {
	currentChapter, err := data.GetChapter(userData.Chapter)
	chapters := data.GetAppData().Chapters
	if err != nil {
		app.Log(chapters)
		currentChapter = &chapters[0]
	}
	return app.Div().Body(
		app.H1().Text("Chapter: "+currentChapter.Name),
		app.Br(),
		app.Select().
			Body(
				app.Range(chapters).Slice(func(i int) app.UI {
					item := chapters[i]
					return item.ToOption(userData.Chapter)
				}),
			).
			OnChange(handleChapterUpdate()),
	)
}

func (user *User) renderPhotoSettings(userData *data.UserData) app.UI {
	photos := data.GetAppData().Photos
	return app.Div().Body(
		app.H1().Text("Photos"),
		app.Select().
			Body(
				app.Range(photos).Slice(func(i int) app.UI {
					item := photos[i]
					return item.ToOption(userData.Chapter)
				}),
			).
			OnChange(handleChapterUpdate()),
	)
}

func (user *User) renderInventionSettings(userData *data.UserData) app.UI {
	inventions := data.GetAppData().Inventions
	return app.Div().Body(
		app.H1().Text("Inventions"),
		app.Select().
			Body(
				app.Range(inventions).Slice(func(i int) app.UI {
					item := inventions[i]
					return item.ToOption(userData.Chapter)
				}),
			).
			OnChange(handleChapterUpdate()),
	)
}

func (user *User) Render() app.UI {
	userData := data.GetUserData()
	return app.Div().Body(
		user.renderChapterSettings(userData),
		user.renderPhotoSettings(userData),
		user.renderInventionSettings(userData),
	)
}
