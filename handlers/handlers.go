package handlers

import (
	"github.com/un-versed/whatsapp-cats/commands"
	"go.mau.fi/whatsmeow"
)

func SetHandlers(c *whatsmeow.Client) {
	// Add Cat Handler
	c.AddEventHandler(func(evt interface{}) {
		commands.CatHandler(evt, c)
	})

	// Add HelloWorld Handler
	c.AddEventHandler(func(evt interface{}) {
		commands.HelloWorldHandler(evt, c)
	})

}
