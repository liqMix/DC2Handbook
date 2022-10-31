package data

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const DEFAULT_DATA_DIR = "/web/json/"
const CHAPTERS_PATH = DEFAULT_DATA_DIR + "chapters.json"
const ITEMS_PATH = DEFAULT_DATA_DIR + "items.json"
const PHOTOS_PATH = DEFAULT_DATA_DIR + "photos.json"
const INVENTIONS_PATH = DEFAULT_DATA_DIR + "inventions.json"

var appData *AppData = nil

type AppData struct {
	Chapters   []Chapter
	Photos     []Photo
	Inventions []Invention
	Items      []Item

	host string
}

/* Init */
func InitAppData(ctx *app.Context) {
	appData = &AppData{}
	appData.host = "http://" + app.Window().URL().Host
	appData.initChapters()
	appData.initItems()
	appData.initPhotos()
	appData.initInventions()
	(*ctx).SetState("loaded", true)
}

func (ad *AppData) initChapters() {
	resp, err := http.Get(ad.host + CHAPTERS_PATH)
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Chapters)
}

func (ad *AppData) initItems() {
	resp, err := http.Get(ad.host + ITEMS_PATH)
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Items)
}

func (ad *AppData) initPhotos() {
	resp, err := http.Get(ad.host + PHOTOS_PATH)
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Photos)
}

func (ad *AppData) initInventions() {
	resp, err := http.Get(ad.host + INVENTIONS_PATH)
	if err != nil {
		app.Log(err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ad.Inventions)
}

/* Get */
func GetAppData() *AppData {
	return appData
}

func GetChapter(chapterID string) (*Chapter, error) {
	for _, v := range appData.Chapters {
		app.Log(v)
		if v.ID == chapterID {
			return &v, nil
		}
	}
	return &Chapter{}, errors.New("Chapter not found")
}
