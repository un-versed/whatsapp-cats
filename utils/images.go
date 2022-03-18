package utils

import (
	"context"
	"io/ioutil"

	waProto "go.mau.fi/whatsmeow/binary/proto"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func GetImageBytes(url string) ([]byte, error) {
	r, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return b, err
}

func UploadImage(url string, caption string, mimetype string, client *whatsmeow.Client, receiver types.JID) error {
	imageBytes, err := GetImageBytes(url)
	if err != nil {
		return err
	}

	resp, err := client.Upload(context.Background(), imageBytes, whatsmeow.MediaImage)
	if err != nil {
		return err
	}

	imageMsg := &waProto.ImageMessage{
		Caption:       proto.String(caption),
		Mimetype:      proto.String(mimetype),
		Url:           &resp.URL,
		DirectPath:    &resp.DirectPath,
		MediaKey:      resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256:    resp.FileSHA256,
		FileLength:    &resp.FileLength,
	}

	_, err = client.SendMessage(receiver, "", &waProto.Message{
		ImageMessage: imageMsg,
	})

	return err
}
