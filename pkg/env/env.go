package env

import (
	"errors"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func ReadEnv(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("environment variable name should not empty")
	}

	value := strings.TrimSpace(os.Getenv(key))
	// if len(value) == 0 {
	// 	message := "Environment variable " + key + " has an empty value"
	// 	return "", errors.New(message)
	// }

	return value, nil
}

func GetEnvString(key string) (string, error) {
	value, err := ReadEnv(key)
	if err != nil {
		return "", err
	}

	return value, nil
}
