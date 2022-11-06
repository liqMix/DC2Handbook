package types

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type IBase interface {
	ToLinkItem(s Status) LinkItem
	ToOption(selectedID string) app.UI
}

type Base struct {
	ID      string
	Name    string
	Chapter string
}

func (b *Base) ToLinkItem(s Status) *LinkItem {
	return &LinkItem{
		label: b.Name,
		href:  b.ID,
		class: s.ToClass(),
	}
}

func (b *Base) ToOption(selectedID string) app.UI {
	isSelected := b.ID == selectedID
	return app.Option().Text(b.Name).Value(b.ID).
		Selected(isSelected).
		Disabled(isSelected)
}
