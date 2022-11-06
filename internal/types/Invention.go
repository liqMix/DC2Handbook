package types

type RecipeItem struct {
	ItemID string
	Count  int
}

type Invention struct {
	Base
	ImageURL    string
	Chapter     string
	Description string
	Photos      []Photo
	Recipe      []RecipeItem
}

func (i *Invention) ToLinkItem(s Status) LinkItem {
	class := s.ToClass()
	return LinkItem{
		label:   i.Name,
		href:    "/inventions#" + i.ID,
		class:   class,
		Chapter: i.Chapter,
	}
}
