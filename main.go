package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/un-versed/whatsapp-cats/whatsapp"
)

func main() {
	err := whatsapp.Connect()
	if err != nil {
		panic(err)
	}
	defer whatsapp.Disconnect()
}
