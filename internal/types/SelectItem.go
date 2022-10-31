package types

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type SelectItem struct {
	Label string
	Value interface{}
}

func NewSelectItem(label string, value interface{}) SelectItem {
	return SelectItem{
		label,
		value,
	}
}

func (sl *SelectItem) ToOption(selectedValue interface{}) app.UI {
	isSelected := sl.Value == selectedValue
	return app.Option().Text(sl.Label).Value(sl.Value).
		Selected(isSelected).
		Disabled(isSelected)
}
