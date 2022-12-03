package whatsapp

import (
	// 	"context"
	// 	"fmt"
	// 	"log"
	// 	"os"
	// 	"whatsapp-bot/pkg/whatsapp/handler"
	// 	"whatsapp-bot/pkg/whatsapp/services"

	"encoding/base64"

	_ "github.com/mattn/go-sqlite3"
	// 	"github.com/mdp/qrterminal"
	// 	"go.mau.fi/whatsmeow"
	"errors"
	"log"

	qrCode "github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var WhatsAppDatastore *sqlstore.Container

// var Cli *whatsmeow.Client

func NewClient() (*whatsmeow.Client, error) {
	WhatsAppDatastore, err := sqlstore.New("sqlite3", "file:db/WhatsApp.db?_foreign_keys=on", nil)

	if err != nil {
		log.Printf("Error connect database %v", err)
		return nil, errors.New("error connect database")
	}

	device, err := WhatsAppDatastore.GetFirstDevice()
	if err != nil {
		log.Printf("Failed to get device: %v", err)
		return nil, errors.New("failed to get device")

	}

	return whatsmeow.NewClient(device, waLog.Stdout("Client", "INFO", true)), nil

	// 	client.AddEventHandler(eventHandler)
	// 	Cli = client

	// 	if client.Store.ID == nil {
	// 		// No ID stored, new login
	// 		qrChan, _ := client.GetQRChannel(context.Background())
	// 		err = client.Connect()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		for evt := range qrChan {
	// 			if evt.Event == "code" {
	// 				// Render the QR code here
	// 				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
	// 				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
	// 				fmt.Println("QR code:", evt.Code)

	//			} else {
	//				fmt.Println("Login event:", evt.Event)
	//			}
	//		}
	//	} else {
	//
	//		// Already logged in, just connect
	//		err = client.Connect()
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
}

func WhatsAppGenerateQR(qrChan <-chan whatsmeow.QRChannelItem) (string, int) {
	qrChanCode := make(chan string)
	qrChanTimeout := make(chan int)

	// Get QR Code Data and Timeout
	go func() {
		for evt := range qrChan {
			if evt.Event == "code" {
				qrChanCode <- evt.Code
				qrChanTimeout <- int(evt.Timeout.Seconds())
			}
		}
	}()

	// Generate QR Code Data to PNG Image
	qrTemp := <-qrChanCode
	qrPNG, _ := qrCode.Encode(qrTemp, qrCode.Medium, 256)

	// Return QR Code PNG in Base64 Format and Timeout Information
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrPNG), <-qrChanTimeout
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		print(v.Message.Conversation)
	}
}
