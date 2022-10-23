package types

type RecipeItem struct {
	Item *Item
	Count uint8
}

type Invention struct {
	ID string
	Name string
	ImageURL string
	Chapter uint8
	Description string
	Photos []*Photo
	Recipe []RecipeItem
}