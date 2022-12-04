package types

import (
	"log"
	"time"
	"whatsapp-bot/pkg/database"
)

func init() {
	dbConnect := database.InitMySQL(database.NewConnection())
	dbConnect.AutoMigrate(&Reminder{})
}

type Reminder struct {
	Id              uint `gorm:"primaryKey"`
	User            string
	Content         string
	Execute_at      time.Time
	IsActive        bool
	OriginalMessage string
}

var dbConnect = database.InitMySQL(database.NewConnection())

func (data *Reminder) Store() error {
	dbConnect.Create(&data)
	return nil
}

func GetDueReminders() ([]Reminder, error) {
	var Reminders []Reminder
	dbConnect.Where("execute_at <= ? AND is_active = 1", time.Now()).Find(&Reminders)
	return Reminders, nil
}

func (data *Reminder) DisableReminder() error {
	err := dbConnect.Model(&data).Update("is_active", 0).Error
	if err != nil {
		log.Panic("Error: " + err.Error())
	}
	return nil
}
