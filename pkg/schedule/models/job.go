package models

import "time"

type Job struct {
	Id         uint `gorm:"primaryKey"`
	User       string
	Content    string
	Execute_at time.Time
	IsActive   bool
}
