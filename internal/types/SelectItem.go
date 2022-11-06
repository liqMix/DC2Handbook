package types

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type SelectItem struct {
	Label string
	Value interface{}
	Class string
}

func NewSelectItem(label string, value interface{}, class string) SelectItem {
	// if no value, use the label
	if value == nil {
		value = label
	}
	return SelectItem{
		label,
		value,
		class,
	}
}

func (sl SelectItem) ToOption(selected interface{}) app.HTMLOption {
	isSelected := selected != nil && sl.Value == selected
	return app.Option().Text(sl.Label).Value(sl.Value).
		Class(sl.Class).
		Selected(isSelected).
		Disabled(isSelected)
}

func CreateSelectInput(items []SelectItem, selected interface{}) app.HTMLSelect {
	return app.Select().Body(
		app.Range(items).Slice(func(i int) app.UI {
			return items[i].ToOption(selected)
		}),
	)
}
