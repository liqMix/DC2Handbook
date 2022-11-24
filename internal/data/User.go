package data

import (
	"encoding/json"

	"github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const LOCAL_STORAGE_KEY = "userdata"

type UserData struct {
	context    *app.Context    `json:"-"`
	Photos     map[string]bool `json:"photos"`
	Inventions map[string]bool `json:"inventions"`
	Chapter    string          `json:"chapter"`
	update     bool
}

var userData *UserData = &UserData{}

func GetUserData() *UserData {
	return userData
}

func InitUserData(ctx *app.Context) {
	userData = fetchUserData(ctx)
}

func ResetUserData(ctx *app.Context) {
	userData = &UserData{
		context: ctx,
	}
	userData.save()
}

/* Chapter */
func (ud *UserData) SetChapter(chapter string) {
	ud.Chapter = chapter
	ud.save()
}

/* Items */
func (ud *UserData) GetItemStatus(item *types.Item) types.Status {
	// TODO: give an inkling of a microthought to what this does
	return types.NA
}

/* Photos */
func (ud *UserData) GetPhotoStatus(photo *types.Photo) types.Status {
	if ud.HasPhoto(photo.ID) {
		return types.TAKEN
	}
	if ud.Chapter > photo.Chapter && photo.Missable {
		return types.MISSED
	}
	if ud.Chapter < photo.Chapter {
		return types.UPCOMING
	}
	if ud.Chapter == photo.Chapter {
		return types.NEW
	}
	return types.AVAILABLE
}

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
func (ud *UserData) GetInventionStatus(i *types.Invention) types.Status {
	if ud.HasInvention(i.ID) {
		return types.INVENTED
	}

	// If any of the required photos are missed
	for _, v := range i.Photos {
		p, err := GetPhoto(v.ID)
		if err == nil && ud.GetPhotoStatus(&p) == types.MISSED {
			return types.MISSED
		}
	}

	if ud.Chapter < i.Chapter {
		return types.UPCOMING
	}
	if ud.Chapter == i.Chapter {
		return types.NEW
	}
	return types.AVAILABLE
}

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

	ud.update = !ud.update
	(*ud.context).SetState(types.E_USERUPDATE, ud.update)
}

func fetchUserData(ctx *app.Context) *UserData {
	var userData = &UserData{}
	var data string
	(*ctx).LocalStorage().Get(LOCAL_STORAGE_KEY, &data)
	json.Unmarshal([]byte(data), userData)
	userData.context = ctx
	if userData.Photos == nil {
		userData.Photos = make(map[string]bool)
	}
	if userData.Inventions == nil {
		userData.Inventions = make(map[string]bool)
	}
	return userData
}
