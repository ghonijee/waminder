package reminder

// import (
// 	"context"
// 	"log"
// 	"time"
// 	"whatsapp-bot/pkg/schedule/data"
// 	"whatsapp-bot/pkg/schedule/models"
// 	waService "whatsapp-bot/pkg/whatsapp/services"
// )

// func Start(ctx context.Context) {
// 	dbConnect := database.InitMySQL(database.NewConnection())
// 	dbConnect.AutoMigrate(&models.Job{})

// 	var whatsAppService waService.WhatsAppService
// 	whatsAppService.Cli = whatsapp.Cli
// 	jobService.CheckJobsInInterval(ctx, time.Minute, whatsAppService)
// }

// func AddJob(job models.Job) error {
// 	data.Store(job)

// 	return nil
// }

// func GetJobs() []models.Job {
// 	jobs, err := data.GetDueJobs()
// 	if err != nil {
// 		log.Panic("Error: " + err.Error())
// 	}
// 	return jobs
// }
