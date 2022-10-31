package types

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PhotoStatus string

const (
	AVAILABLE PhotoStatus = "AVAILABLE"
	MISSED                = "MISSED"
	NEW                   = "NEW"
	TAKEN                 = "TAKEN"
	UPCOMING              = "UPCOMING"
)

func (ps PhotoStatus) ToUI(userChapter string, photoChapter string) app.UI {
	switch ps {
	case NEW:
		return app.H5().Body(
			app.Text("NEW"),
		).Class("warn")

	case TAKEN:
		return app.H5().Body(
			app.Text("TAKEN"),
		).Class("success")

	case MISSED:
		return app.H5().Body(
			app.Text("MISSED"),
		).Class("error")

	case UPCOMING:
		ucInt, err := strconv.Atoi(userChapter)
		if err != nil {
			ucInt = 0
		}
		pcInt, err := strconv.Atoi(photoChapter)
		if err != nil {
			pcInt = 0
		}
		return app.H5().Body(
			app.Text(fmt.Sprintf("IN %v CHAPTER(S)", pcInt-ucInt)),
		).Class("deemphasize")

	case AVAILABLE:
	default:
		// Fall through to return below
	}
	return app.H5().Body(
		app.Text("AVAILABLE"),
	).Class("moderate-warn")
}

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

func (p *Photo) ToLinkItem() *LinkItem {
	name := p.Name
	if p.IsScoop {
		name += " (S)"
	}
	return &LinkItem{
		label: name,
		href:  "/photos#" + p.ID,
	}
}
