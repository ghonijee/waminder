package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"whatsapp-bot/pkg/schedule"
	scheduleRoute "whatsapp-bot/pkg/schedule/routes"
	"whatsapp-bot/pkg/whatsapp"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

func main() {
	echo := echo.New()
	scheduleRoute.RouteAPI(echo)

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	ctx := context.Background()

	whatsapp.Start()

	schedule.Start(ctx)

	if err := echo.StartH2CServer(":8000", s); err != http.ErrServerClosed {
		log.Fatal(err)
		<-ctx.Done()
	}
}
