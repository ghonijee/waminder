package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"whatsapp-bot/internal/router"
	"whatsapp-bot/pkg/env"

	"github.com/labstack/echo/v4"
	"go.mau.fi/whatsmeow"
)

type WhatsAppCLI *whatsmeow.Client
type Server struct {
	Address string
	Port    string
}

func main() {
	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	var err error
	echo := echo.New()
	router := routeApiConfig(echo)
	routeInit(router)
	// // waClient, err := whatsapp.NewClient()
	// if err != nil {
	// 	log.Panic(err)
	// 	return
	// }
	// Get Server Configuration
	var serverConfig Server

	serverConfig.Address, err = env.GetEnvString("SERVER_ADDRESS")
	if err != nil {
		serverConfig.Address = "127.0.0.1"
	}

	serverConfig.Port, err = env.GetEnvString("SERVER_PORT")
	if err != nil {
		serverConfig.Port = "3000"
	}

	// Start Server
	go func() {
		err := echo.Start(serverConfig.Address + ":" + serverConfig.Port)
		if err != nil && err != http.ErrServerClosed {
			log.Panicln(err.Error())
		}
	}()

	// Watch for Shutdown Signal
	sigShutdown := make(chan os.Signal, 1)
	signal.Notify(sigShutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigShutdown

	// Wait 5 Seconds Before Graceful Shutdown
	defer cancelShutdown()

	// Try To Shutdown Server
	<-ctxShutdown.Done()
	err = echo.Shutdown(ctxShutdown)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func routeApiConfig(echo *echo.Echo) *echo.Group {
	prefix := "/api/v1"
	api := echo.Group(prefix)
	return api
}

func routeInit(api *echo.Group) {
	router.ReminderApi(api)
	router.WhatsappApi(api)
}
