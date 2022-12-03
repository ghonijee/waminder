package whatsapp

import (
	"context"

	"go.mau.fi/whatsmeow"
	waproto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

type WhatsAppService struct {
	Cli *whatsmeow.Client
}

func (wa *WhatsAppService) SendMessage(number string, message string) (string, error) {
	remoteJID := types.NewJID(number, types.DefaultUserServer)
	msgId := whatsmeow.GenerateMessageID()
	msgContent := &waproto.Message{
		Conversation: proto.String(message),
	}

	// Send WhatsApp Message Proto
	_, err := wa.Cli.SendMessage(context.Background(), remoteJID, msgId, msgContent)
	if err != nil {
		return "", err
	}

	return msgId, nil
}
