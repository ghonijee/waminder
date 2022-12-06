package database

import (
	"fmt"
	"whatsapp-bot/pkg/env"
)

type Database struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func NewConnection() *Database {
	host, _ := env.GetEnvString("DB_HOST")
	port, _ := env.GetEnvString("DB_PORT")
	userName, _ := env.GetEnvString("DB_USERNAME")
	password, _ := env.GetEnvString("DB_PASSWORD")
	dbName, _ := env.GetEnvString("DB_NAME")

	return &Database{
		Host:     host,
		Port:     port,
		Name:     dbName,
		Username: userName,
		Password: password,
	}
}

func (db *Database) ToMySQL() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Port, db.Name)
	return dsn
}
