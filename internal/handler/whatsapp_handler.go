package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"whatsapp-bot/pkg/whatsapp"

	"github.com/labstack/echo/v4"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
)

var WhatsAppDatastore *sqlstore.Container
var client *whatsmeow.Client

func init() {
	fmt.Print("Init WhatsApp hanlder")

}

func Index(ctx echo.Context) error {
	var err error
	client, err = whatsapp.NewClient()
	if err != nil {
		return ctx.JSON(500, err.Error())
	}

	if client.Store.ID == nil {
		return ctx.JSON(200, "Belum ada akun yang terhubung")
	}

	return ctx.JSON(200, "WhatsApp sudah terhubung")
}

func Logut(ctx echo.Context) error {
	var err error
	client, err = whatsapp.NewClient()
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	_ = client.SendPresence(types.PresenceUnavailable)
	err = client.Logout()
	if err != nil {
		client.Disconnect()
		err = client.Store.Delete()
		return ctx.JSON(500, err.Error())
	}

	return ctx.String(200, "Logout Success")
}

type RequestLogin struct {
	Output string
}

func Login(ctx echo.Context) error {
	var err error
	// jid := jwtPayload(c).JID

	var reqLogin RequestLogin
	reqLogin.Output = strings.TrimSpace(ctx.FormValue("output"))

	if len(reqLogin.Output) == 0 {
		reqLogin.Output = "html"
	}
	client, err = whatsapp.NewClient()
	if err != nil {
		return ctx.JSON(500, err.Error())
	}

	if client.Store.ID != nil {
		return ctx.JSON(200, "Sudah ada akun yang terhubung")
	}

	qrChanGenerate, _ := client.GetQRChannel(context.Background())
	err = client.Connect()

	if err != nil {
		return ctx.JSON(500, err.Error())
	}

	_ = client.SendPresence(types.PresenceAvailable)

	qrCode, qrTimeOut := whatsapp.WhatsAppGenerateQR(qrChanGenerate)

	htmlContent := `
	<html>
	  <head>
		<title>WhatsApp Multi-Device Login</title>
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
	  </head>
	  <body>
		<img src="` + qrCode + `" />
		<p>
		  <b>QR Code Scan</b>
		  <br/>
		  Timeout in ` + strconv.Itoa(qrTimeOut) + ` Second(s)
		</p>
	  </body>
	</html>`

	return ctx.HTML(200, htmlContent)

}
