package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(driver *Database) *gorm.DB {
	db, err := gorm.Open(mysql.Open(driver.ToMySQL()), &gorm.Config{})
	if err != nil {
		log.Panicf("could not connect to database: %s", err.Error())
	}

	return db
}
