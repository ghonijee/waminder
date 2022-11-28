package env

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	host := os.Getenv("DB_HOST")
	if host != "127.0.0.1" {
		t.Errorf("Error, got %s want 127.0.0.1", host)
	}
}
