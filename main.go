package main

import (
	"log"
	"whatsapp-bot/cmd/server"
	"whatsapp-bot/internal/database"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	setDefaultEnv()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	serverApp := server.Server{
		Port:    viper.GetString("SERVER_PORT"),
		Address: viper.GetString("SERVER_ADDRESS"),
	}
	server.StartRestApp(serverApp, *dbConfigDatabase())
}

func setDefaultEnv() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USERNAME", "user")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "development")
	viper.SetDefault("SERVER_ADDRESS", "localhost")
	viper.SetDefault("SERVER_PORT", "8000")
	viper.SetDefault("APP_NAME", "Waminde-bot")
}

func dbConfigDatabase() *database.Database {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	userName := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	return &database.Database{
		Host:     host,
		Port:     port,
		Name:     dbName,
		Username: userName,
		Password: password,
	}
}
