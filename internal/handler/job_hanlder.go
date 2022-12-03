package handler

// import (
// 	"fmt"
// 	"time"
// 	"whatsapp-bot/pkg/schedule/data"
// 	"whatsapp-bot/pkg/schedule/models"

// 	"github.com/labstack/echo/v4"
// )

// func Index(c echo.Context) error {
// 	return c.JSON(200, "YES")
// }

// func Store(c echo.Context) error {
// 	content := "Test"
// 	execAtTime, _ := time.Parse("2006-01-02 15:04:05 MST", "2022-11-28 16:51:00 WIB")
// 	fmt.Println(execAtTime)
// 	var jobData models.Job
// 	jobData.Content = content
// 	jobData.Execute_at = execAtTime
// 	jobData.IsActive = true

// 	data.Store(jobData)
// 	return nil
// }
