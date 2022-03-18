package commands

import (
	"fmt"
	"strings"

	"github.com/un-versed/whatsapp-cats/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func HelloWorldHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())

		switch msg {
		case "hello world":
			err := utils.SendMessage("Hello world!", c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
