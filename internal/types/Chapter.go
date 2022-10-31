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
