package whatsapp

import (
	"context"
	"fmt"
	"log"
	"os"
	"whatsapp-bot/pkg/whatsapp/handler"
	"whatsapp-bot/pkg/whatsapp/services"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var WhatsAppDatastore *sqlstore.Container
var Cli *whatsmeow.Client

func eventHandler(evt interface{}) {
	var whatsAppService services.WhatsAppService
	whatsAppService.Cli = Cli
	switch v := evt.(type) {
	case *events.Message:
		handler := handler.WhatsAppHandler{
			Service: whatsAppService,
		}
		handler.ReceivedMessage(v)
	}
}

func Start() {
	datastore, err := sqlstore.New("sqlite3", "file:db/WhatsApp.db?_foreign_keys=on", nil)

	if err != nil {
		log.Printf("Error connect database %v", err)
		return
	}
	WhatsAppDatastore = datastore

	device, err := WhatsAppDatastore.GetFirstDevice()
	if err != nil {
		log.Printf("Failed to get device: %v", err)
		return
	}

	client := whatsmeow.NewClient(device, waLog.Stdout("Client", "INFO", true))

	client.AddEventHandler(eventHandler)
	Cli = client

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("QR code:", evt.Code)

			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}
}
