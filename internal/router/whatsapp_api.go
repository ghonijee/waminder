package router

import (
	"whatsapp-bot/internal/handler"

	"github.com/labstack/echo/v4"
)

func WhatsappApi(e *echo.Group) {
	// whatsapps/login
	e.GET("/whatsapp", handler.Index)
	e.GET("/whatsapp/logout", handler.Logut)
	e.GET("/whatsapp/login", handler.Login)
	// whatsapps/logout
	// whatsapps/status
}
