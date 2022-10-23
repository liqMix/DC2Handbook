package components

import (
	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type LinkItem struct {
	app.Compo

	label string
	href  string
}

func (l *LinkItem) Render() app.UI {
	return app.A().Body(
		app.Text(l.label),
	).Href(l.href)
}

func LinkItemFromItem(i Item, path string) LinkItem {
	if path == "" {
		path = "/items#"
	}

	linkItem := LinkItem{
		label: i.Name,
		href:  path + i.ID,
	}
	return linkItem
}

func LinkItemFromPhoto(p Photo, path string) LinkItem {
	if path == "" {
		path = "/photos#"
	}

	name := p.Name
	if p.IsScoop {
		name += " (S)"
	}

	linkItem := LinkItem{
		label: name,
		href:  path + p.ID,
	}
	return linkItem
}

func LinkItemFromInvention(i Invention, path string) LinkItem {
	if path == "" {
		path = "/inventions#"
	}

	linkItem := LinkItem{
		label: i.Name,
		href:  path + i.ID,
	}
	return linkItem
}

func LinkItemFromStrings(label string, href string, path string) LinkItem {
	linkItem := LinkItem{
		label: label,
		href:  path + href,
	}
	return linkItem
}
