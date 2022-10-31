package types

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Item struct {
	Base
	ImageURL string
	Value    string
}

func (i *Item) ToLinkItem() *LinkItem {
	return &LinkItem{
		label: i.Name,
		href:  "/items#" + i.ID,
	}
}

func (i *Item) ToOption(selectedID string) app.UI {
	isSelected := i.ID == selectedID
	return app.Option().Text(i.Name).Value(i.ID).
		Selected(isSelected).
		Disabled(isSelected)
}
