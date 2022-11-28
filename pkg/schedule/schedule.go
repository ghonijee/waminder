package schedule

import (
	"context"
	"time"
	"whatsapp-bot/pkg/database"
	"whatsapp-bot/pkg/schedule/models"
	jobService "whatsapp-bot/pkg/schedule/services"
	"whatsapp-bot/pkg/whatsapp"
	waService "whatsapp-bot/pkg/whatsapp/services"
)

func Start(ctx context.Context) {
	dbConnect := database.InitMySQL(database.NewConnection())
	dbConnect.AutoMigrate(&models.Job{})

	var whatsAppService waService.WhatsAppService
	whatsAppService.Cli = whatsapp.Cli
	jobService.CheckJobsInInterval(ctx, time.Minute, whatsAppService)
}
