package types

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Status string

const (
	AVAILABLE Status = "AVAILABLE"
	MISSED    Status = "MISSED"
	NEW       Status = "NEW"
	TAKEN     Status = "TAKEN"
	INVENTED  Status = "INVENTED"
	UPCOMING  Status = "UPCOMING"
	NA        Status = "NA"
)

func StatusFromString(str string) Status {
	stringMap := map[string]Status{
		"AVAILABLE": AVAILABLE,
		"MISSED":    MISSED,
		"NEW":       NEW,
		"TAKEN":     TAKEN,
		"INVENTED":  INVENTED,
		"UPCOMING":  UPCOMING,
	}

	status, ok := stringMap[str]
	if ok {
		return status
	}
	return NA
}

func (s Status) ToClass() string {
	class := ""
	switch s {
	case AVAILABLE:
		class = "moderate-warn"
	case NEW:
		class = "warn"

	case TAKEN:
		class = "success"
	case INVENTED:
		class = "success"

	case MISSED:
		class = "error"

	case UPCOMING:
		class = "deemphasize"
	}
	return class
}

func (s Status) ToUI(userChapter string, photoChapter string, icon string) app.HTMLH5 {
	text := string(s)
	class := s.ToClass()

	if s == UPCOMING {
		ucInt, err := strconv.Atoi(userChapter)
		if err != nil {
			ucInt = 0
		}
		pcInt, err := strconv.Atoi(photoChapter)
		if err != nil {
			pcInt = 0
		}
		text = fmt.Sprintf("IN %v CHAPTER(S)", pcInt-ucInt)
	}

	return app.H5().Body(
		app.Text(icon),
		app.Text(text),
	).Class("status " + class)
}

func (s Status) toSelectItem() SelectItem {
	si := NewSelectItem(
		string(s),
		s,
		s.ToClass(),
	)
	return si
}

func CreateStatusSelect(selected Status) app.HTMLSelect {
	return CreateSelectInput(
		[]SelectItem{
			NewSelectItem("ALL", string(NA), ""),
			NEW.toSelectItem(),
			AVAILABLE.toSelectItem(),
			UPCOMING.toSelectItem(),
			MISSED.toSelectItem(),
			TAKEN.toSelectItem(),
			INVENTED.toSelectItem(),
		},
		selected,
	)
}
