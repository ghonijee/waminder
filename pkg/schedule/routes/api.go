package routes

import (
	job "whatsapp-bot/pkg/schedule/handler"

	"github.com/labstack/echo/v4"
)

func RouteAPI(e *echo.Echo) {
	api := e.Group("/api/v1")

	// Job
	api.GET("/job", job.Index)
	api.POST("/job", job.Store)
}
