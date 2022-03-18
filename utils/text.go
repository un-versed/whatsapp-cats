package utils

import (
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func SendMessage(message string, client *whatsmeow.Client, receiver types.JID) error {
	_, err := client.SendMessage(receiver, "", &waProto.Message{
		Conversation: proto.String(message),
	})

	return err
}
