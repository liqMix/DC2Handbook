package data

import (
	"encoding/json"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const LOCAL_STORAGE_KEY = "userdata"

type UserData struct {
	context    *app.Context    `json:"-"`
	Photos     map[string]bool `json:"photos"`
	Inventions map[string]bool `json:"inventions"`
	Chapter    uint8           `json:"chapter"`
}

/* Photos */
func (ud *UserData) HasPhoto(photoID string) bool {
	var current, ok = ud.Photos[photoID]
	if !ok {
		return false
	} else {
		return current
	}
}
func (ud *UserData) TogglePhoto(photoID string) {
	var current, ok = ud.Photos[photoID]
	if !ok {
		ud.Photos[photoID] = true
	} else {
		ud.Photos[photoID] = !current
	}
	ud.save()
}

/* Inventions */
func (ud *UserData) HasInvention(inventionID string) bool {
	var current, ok = ud.Inventions[inventionID]
	if !ok {
		return false
	} else {
		return current
	}
}
func (ud *UserData) ToggleInvention(inventionID string) {
	var current, ok = ud.Inventions[inventionID]
	if !ok {
		ud.Inventions[inventionID] = true
	} else {
		ud.Inventions[inventionID] = !current
	}
	ud.save()
}

/* Local Storage Handlers */
func (ud *UserData) save() {
	var data, err = json.Marshal(*ud)
	if err != nil {
		app.Log(err)
	}
	var localStorage = (*ud.context).LocalStorage()
	err = localStorage.Set(LOCAL_STORAGE_KEY, string(data))
	if err != nil {
		app.Log(err)
	}
}

func FetchUserData(ctx *app.Context) *UserData {
	var userData = &UserData{
		Photos:     make(map[string]bool),
		Inventions: make(map[string]bool),
	}
	var data string
	(*ctx).LocalStorage().Get(LOCAL_STORAGE_KEY, &data)
	json.Unmarshal([]byte(data), userData)
	userData.context = ctx
	return userData
}
