package data

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const IMAGE_DIR = "/web/img/"
const PLACEHOLDER_IMAGE = IMAGE_DIR + "placeholder.png"
const JSON_DIR = "/web/json/"

var appData *AppData = nil

type AppData struct {
	Chapters   []Chapter
	Photos     []Photo
	Inventions []Invention
	Items      []Item

	byId map[string]interface{}
	host string
}

/* Init */
func InitAppData(ctx *app.Context) {
	appData = &AppData{}
	appData.host = "http://" + app.Window().URL().Host
	appData.byId = make(map[string]interface{})
	appData.initInventions()
	appData.initPhotos()
	appData.initChapters()
	appData.initItems()
	(*ctx).SetState(string(E_LOADED), true)
}

func (ad *AppData) initChapters() {
	resp, err := http.Get(ad.host + JSON_DIR + "chapters.json")
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Chapters)
	for _, v := range ad.Chapters {
		ad.byId[v.ID] = v
	}
}

func (ad *AppData) initItems() {
	resp, err := http.Get(ad.host + JSON_DIR + "items.json")
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Items)
	for _, v := range ad.Items {
		ad.byId[v.ID] = v
	}
}

func (ad *AppData) initPhotos() {
	resp, err := http.Get(ad.host + JSON_DIR + "photos.json")
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Photos)
	for _, v := range ad.Photos {
		ad.byId[v.ID] = v
	}
}

func (ad *AppData) initInventions() {
	resp, err := http.Get(ad.host + JSON_DIR + "inventions.json")
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Inventions)
	for _, v := range ad.Inventions {
		ad.byId[v.ID] = v
	}
}

/* Get */
func GetAppData() *AppData {
	return appData
}

func GetChapter(id string) (Chapter, error) {
	c, ok := appData.byId[id]
	if !ok {
		return Chapter{}, errors.New("Chapter not found")
	}
	return c.(Chapter), nil
}

func GetPhoto(id string) (Photo, error) {
	p, ok := appData.byId[id]
	if !ok {
		return Photo{}, errors.New("Photo not found")
	}
	return p.(Photo), nil
}

func GetInvention(id string) (Invention, error) {
	i, ok := appData.byId[id]
	if !ok {
		return Invention{}, errors.New("Invention not found")
	}
	return i.(Invention), nil
}

func GetItem(id string) (Item, error) {
	i, ok := appData.byId[id]
	if !ok {
		return Item{}, errors.New("Item not found")
	}
	return i.(Item), nil
}
