package commands

import (
	"fmt"
	"strings"

	"github.com/un-versed/whatsapp-cats/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type Cat struct {
	Breeds []interface{} `json:"breeds"`
	ID     string        `json:"id"`
	URL    string        `json:"url"`
	Width  int           `json:"width"`
	Height int           `json:"height"`
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
	utils.GetJson("https://api.thecatapi.com/v1/images/search?limit=1&size=full&mime_types=jpg", &c)
	err := utils.UploadImage(
		c[0].URL,
		"um gatinho fofo pra vc ☺️",
		"image/png",
		client,
		receiver)

	return err
}
