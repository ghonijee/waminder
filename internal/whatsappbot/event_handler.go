package whatsappbot

import (
	"fmt"
	"log"
	"whatsapp-bot/internal/types"
	messageparse "whatsapp-bot/pkg/message_parse"

	"go.mau.fi/whatsmeow/types/events"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		var messagePayload messageparse.MessagePayload
		messagePayload, err := messageparse.Parse(v.Message.GetConversation())
		if err != nil {
			fmt.Println("Error parse message to struct : " + err.Error())
		}
		BotFunctionHandler(v, messagePayload)
	}
}

func BotFunctionHandler(message *events.Message, payload messageparse.MessagePayload) {
	switch payload.Type {
	case Remindme:
		// Store to DB
		var reminder types.Reminder
		reminder.User = message.Info.Sender.User
		reminder.Content = payload.Message
		reminder.Execute_at = payload.Execute_at
		reminder.IsActive = true
		reminder.OriginalMessage = message.Message.GetConversation()
		if reminder.Store() != nil {
			log.Println("Failed to store data")
		} else {
			whatsAppService.SendMessage(reminder.User, "Okey, i will remind you at "+reminder.Execute_at.Format("02/01/2006 15:04:05 MST"))
		}
		// Reply Oke, i will remind you

	default:
		whatsAppService.SendMessage(message.Info.Sender.User, "Maaf fungsi yang kamu masukan tidak tersedia \n\n Silakan kirim /help untuk lihat format dan bantuan")

		// Send, Maaf saya tidak faham perintahnya,
		// silahkan kirim /help untuk lihat format dan bantuan
	}
}

const (
	Remindme string = "/remindme"
	Help     string = "/help"
)
