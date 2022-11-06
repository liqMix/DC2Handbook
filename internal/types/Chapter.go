package types

type Chapter struct {
	Base
}

func (c *Chapter) ToLinkItem() *LinkItem {
	return &LinkItem{
		label: c.Name,
		href:  "/chapters#" + c.ID,
	}
}

func (c *Chapter) ToLabel() string {
	label := c.Name
	if c.ID != "0" {
		label = "Chapter " + c.ID + ": " + c.Name
	}
	return label
}
