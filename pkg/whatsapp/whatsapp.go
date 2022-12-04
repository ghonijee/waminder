package whatsapp

import (
	"encoding/base64"

	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"

	qrCode "github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func GetClient() (*whatsmeow.Client, error) {
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
