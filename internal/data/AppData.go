package data

import (
	"encoding/json"
	"io"
	"net/http"

	. "github.com/liqMix/DC2Photobook/internal/types"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const DEFAULT_DATA_DIR = "/web/json/"
const ITEMS_PATH = DEFAULT_DATA_DIR + "items.json"
const PHOTOS_PATH = DEFAULT_DATA_DIR + "photos.json"
const INVENTIONS_PATH = DEFAULT_DATA_DIR + "inventions.json"

var applicationData *AppData = &AppData{}

type AppData struct {
	UserData   *UserData
	Photos     []Photo
	Inventions []Invention
	Items      []Item

	host string
}

func ApplicationData() *AppData {
	return applicationData
}

func InitAppData(ctx *app.Context) {
	applicationData.host = "http://" + app.Window().URL().Host
	applicationData.initItems()
	applicationData.initPhotos()
	applicationData.initInventions()
	applicationData.UserData = FetchUserData(ctx)
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
