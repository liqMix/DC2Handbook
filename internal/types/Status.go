package types

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Status string

const (
	AVAILABLE Status = "AVAILABLE"
	MISSED           = "MISSED"
	NEW              = "NEW"
	TAKEN            = "TAKEN"
	INVENTED         = "CREATED"
	UPCOMING         = "UPCOMING"
)

func (s Status) ToUI(userChapter string, photoChapter string, icon string) app.HTMLH5 {
	text := string(s)
	class := "moderate-warn"

	switch s {
	case NEW:
		class = "warn"

	case TAKEN:
	case INVENTED:
		class = "success"

	case MISSED:
		class = "error"

	case UPCOMING:
		ucInt, err := strconv.Atoi(userChapter)
		if err != nil {
			ucInt = 0
		}
		pcInt, err := strconv.Atoi(photoChapter)
		if err != nil {
			pcInt = 0
		}
		text = fmt.Sprintf("IN %v CHAPTER(S)", pcInt-ucInt)
		class = "deemphasize"

	case AVAILABLE:
	default:
		// Fall through to return below
	}
	return app.H5().Body(
		app.Text(icon),
		app.Text(text),
	).Class("status " + class)
}
