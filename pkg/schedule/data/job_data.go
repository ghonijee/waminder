package data

import (
	"fmt"
	"log"
	"time"
	"whatsapp-bot/pkg/database"
	"whatsapp-bot/pkg/schedule/models"
)

var dbConnect = database.InitMySQL(database.NewConnection())

func Store(data models.Job) error {
	dbConnect.Create(&data)
	return nil
}

func GetDueJobs() ([]models.Job, error) {
	var jobs []models.Job
	dbConnect.Where("execute_at <= ? AND is_active = 1", time.Now()).Find(&jobs)
	return jobs, nil
}

func DisableJob(data models.Job) error {
	fmt.Println(data.Id)
	err := dbConnect.Model(&data).Update("is_active", 0).Error
	if err != nil {
		log.Panic("Error: " + err.Error())
	}
	return nil
}
