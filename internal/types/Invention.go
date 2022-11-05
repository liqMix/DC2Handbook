package types

type RecipeItem struct {
	Item  *Item
	Count uint8
}

type Invention struct {
	Base
	ImageURL    string
	Chapter     string
	Description string
	Photos      []Photo
	Recipe      []RecipeItem
}

func (i *Invention) ToLinkItem() *LinkItem {
	return &LinkItem{
		label: i.Name,
		href:  "/inventions#" + i.ID,
	}
}
