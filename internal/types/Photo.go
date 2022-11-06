package types

type Photo struct {
	Base
	ImageURL     string
	Chapter      string
	Missable     bool
	MissableNote string
	Location     string
	IsScoop      bool
	Memo         string
}

func (p *Photo) ToLinkItem(s Status) LinkItem {
	name := p.Name
	class := s.ToClass()
	if p.IsScoop {
		name += " (S)"
	}
	return LinkItem{
		label:   name,
		href:    "/photos#" + p.ID,
		class:   class,
		Chapter: p.Chapter,
	}
}
