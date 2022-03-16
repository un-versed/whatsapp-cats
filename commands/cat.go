package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/un-versed/whatsapp-cats/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type Cat struct {
	Breeds []interface{} `json:"breeds"`
	ID     string        `json:"id"`
	URL    string        `json:"url"`
	Width  int           `json:"width"`
	Height int           `json:"height"`
}

func GetJson(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func CatHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())

		switch msg {
		case "gato":
			err := SendCatPicture(c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}

		case "cat":
			err := SendCatPicture(c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func SendCatPicture(client *whatsmeow.Client, receiver types.JID) error {
	c := []Cat{}
	GetJson("https://api.thecatapi.com/v1/images/search?limit=1&size=full&mime_types=jpg", &c)
	err := utils.UploadImage(
		c[0].URL,
		"um gatinho fofo pra vc ☺️",
		"image/png",
		client,
		receiver)

	return err
}
