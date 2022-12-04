package messageparse_test

import (
	"fmt"
	"testing"
	"time"
	messageparse "whatsapp-bot/pkg/message_parse"
)

func TestMessageParse(t *testing.T) {
	str := "/remindme \n  31 12 2022 pukul 18:59 \n\n Hei, jangan lupa makan jam 18:30"
	date, _ := messageparse.FindDate(str)
	timeStr, _ := messageparse.FindTime(str)
	messageparse.FindFuncType(str)
	newDate := messageparse.GenerateDate(date, timeStr)
	message, _ := messageparse.Parse(str)
	if newDate != time.Date(2022, 12, 31, 18, 59, 0, 0, time.Local) {
		t.Error("not equals")
	}
	fmt.Println(message)
}
