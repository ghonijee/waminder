package messageparse

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"whatsapp-bot/pkg/schedule/models"
)

type MessagePayload struct {
	Type       string
	Message    string
	Execute_at time.Time
}

func Parse(body string) (MessagePayload, error) {
	if !strings.Contains(body, "RemindMe") {
		return MessagePayload{}, errors.New("command not found")
	}
	content := strings.Split(body, "\n")
	typeMessage := content[0]
	message := strings.Split(content[1], ": ")[1]
	exec_at := strings.Split(content[2], ": ")[1]
	execAtTime, _ := time.Parse("2006-01-02 15:04:05 MST", exec_at+" WIB")
	fmt.Println(typeMessage)
	fmt.Println(message)
	fmt.Println(exec_at)

	return MessagePayload{
		Type:       typeMessage,
		Message:    message,
		Execute_at: execAtTime,
	}, nil
}

func (m *MessagePayload) ToJob() models.Job {
	return models.Job{
		Content:    m.Message,
		Execute_at: m.Execute_at,
		IsActive:   true,
	}
}
