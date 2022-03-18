package commands

import (
	"fmt"
	"strings"

	"github.com/un-versed/whatsapp-cats/utils"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type ShoudDeploy struct {
	Timezone      string `json:"timezone"`
	ShouldIDeploy bool   `json:"shouldideploy"`
	Message       string `json:"message"`
}

func ShoudIDeployTodayHandler(evt interface{}, c *whatsmeow.Client) {
	switch v := evt.(type) {
	case *events.Message:
		msg := strings.ToLower(v.Message.GetConversation())

		if msg == "should i deploy" || msg == "should i deploy today?" || msg == "should i deploy today" {
			err := ShoudIDeployToday(c, v.Info.Chat)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func ShoudIDeployToday(client *whatsmeow.Client, receiver types.JID) error {
	sd := new(ShoudDeploy)
	utils.GetJson("https://shouldideploy.today/api?tz=America/Sao_Paulo", sd)
	err := utils.SendMessage(
		sd.Message,
		client,
		receiver)

	return err
}
