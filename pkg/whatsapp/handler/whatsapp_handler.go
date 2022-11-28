package handler

import (
	"fmt"
	messageparse "whatsapp-bot/pkg/message_parse"
	"whatsapp-bot/pkg/schedule/models"
	scheduleService "whatsapp-bot/pkg/schedule/services"
	"whatsapp-bot/pkg/whatsapp/services"

	"go.mau.fi/whatsmeow/types/events"
)

type WhatsAppHandler struct {
	Service services.WhatsAppService
}

func (wh *WhatsAppHandler) ReceivedMessage(value *events.Message) {
	messageModel, err := messageparse.Parse(value.Message.GetConversation())
	if err != nil {
		return
	}
	fmt.Println(value.Info.Sender.User)
	switch messageModel.Type {
	case "RemindMe":
		// Store to DB
		var newJob models.Job
		newJob = messageModel.ToJob()
		newJob.User = value.Info.Sender.User
		scheduleService.AddJob(newJob)
		// Send WhatsApp OK
		wh.Service.SendMessage(newJob.User, "Ok, siap!")

	}
}
