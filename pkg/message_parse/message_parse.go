package messageparse

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

type MessagePayload struct {
	Type       string
	Message    string
	Execute_at time.Time
}

func Parse(body string) (MessagePayload, error) {
	funcType := FindFuncType(body)
	if len(funcType) == 0 {
		return MessagePayload{}, errors.New("command not found")
	}
	date, _ := FindDate(body)
	timeStr, _ := FindTime(body)
	newDate := GenerateDate(date, timeStr)
	bodySlice := strings.Split(body, "\n")
	content := bodySlice[len(bodySlice)-1]

	return MessagePayload{
		Type:       funcType,
		Message:    content,
		Execute_at: newDate,
	}, nil
}

func FindDate(str string) (string, error) {
	dateRegex, err := regexp.Compile("(([0-1][0-9]|[3][0-1]).([0-5][0-9]).([0-9]{4}))") // Find date
	if err != nil {
		log.Fatal("Error")
		return "", err
	}
	findDate := strings.Split(dateRegex.FindString(str), "")
	findDate[2] = "-"
	findDate[5] = "-"
	standartDateString := strings.Join(findDate, "")
	return standartDateString, nil
}

func FindTime(str string) (string, error) {
	timeRegex, err := regexp.Compile("([0-1][0-9]|[2][0-3]):([0-5][0-9])") // Find Jam
	if err != nil {
		log.Fatal("Error")
		return "", err
	}
	findTime := timeRegex.FindString(str)
	return findTime + ":00 WIB", nil
}

func FindFuncType(str string) string {
	reg, err := regexp.Compile("\\/\\w*?\\b\n|\\/\\w*?\\b \n")
	if err != nil {
		fmt.Println("Error func type parse")
		return ""
	}
	value := reg.FindString(str)
	return value
}

func GenerateDate(dateStr string, timeStr string) time.Time {
	newStr := fmt.Sprintf("%s %s", dateStr, timeStr)
	dateParse, err := time.Parse("02-01-2006 15:04:05 MST", newStr)
	if err != nil {
		log.Println("Error Parse")
		return time.Now()
	}
	return dateParse
}
