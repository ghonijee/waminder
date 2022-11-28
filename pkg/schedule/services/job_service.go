package services

import (
	"context"
	"log"
	"time"
	"whatsapp-bot/pkg/schedule/data"
	"whatsapp-bot/pkg/schedule/models"
	waService "whatsapp-bot/pkg/whatsapp/services"
)

func AddJob(job models.Job) error {
	data.Store(job)

	return nil
}

func GetJobs() []models.Job {
	jobs, err := data.GetDueJobs()
	if err != nil {
		log.Panic("Error: " + err.Error())
	}
	return jobs
}

func CheckJobsInInterval(ctx context.Context, duration time.Duration, whatsAppService waService.WhatsAppService) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				log.Println("Ticks Received...")
				jobs := GetJobs()
				for _, e := range jobs {
					// Send Content
					e.Content = e.Content + "\n\nFrom Waminder"
					whatsAppService.SendMessage(e.User, e.Content)
					// Disable jobs after called
					e.IsActive = false
					data.DisableJob(e)
				}
			}
		}
	}()
}
