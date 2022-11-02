package sqlstore

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DB_URL")

	res := m.Run()

	os.Exit(res)
}
