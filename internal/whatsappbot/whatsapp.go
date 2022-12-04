package whatsappbot

import (
	"context"
	"time"
	"whatsapp-bot/internal/types"
	"whatsapp-bot/pkg/whatsapp"

	"log"

	_ "github.com/mattn/go-sqlite3"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

var WhatsAppDatastore *sqlstore.Container

var Cli *whatsmeow.Client

var errClient error

var whatsAppService whatsapp.WhatsAppService

func Setup(ctx context.Context) {
	Cli, errClient = whatsapp.GetClient()
	if errClient != nil {
		log.Println("Error Get WhatsApp Client: " + errClient.Error())
	}
	Cli.AddEventHandler(eventHandler)
	if Cli.Store.ID != nil {
		Cli.Connect()
		whatsAppService = whatsapp.WhatsAppService{
			Cli: Cli,
		}
	}

	checkJobsInInterval(ctx, time.Minute, whatsAppService)
}

func checkJobsInInterval(ctx context.Context, duration time.Duration, whatsAppService whatsapp.WhatsAppService) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				log.Println("Ticks Received...")
				jobs, _ := types.GetDueReminders()
				for _, e := range jobs {
					// Send Content
					whatsAppService.SendMessage(e.User, e.Content)
					// Disable jobs after called
					e.DisableReminder()
					whatsAppService.SendMessage(e.User, "Thanks for using our bot \nWaminderBot")
				}
			}
		}
	}()
}
