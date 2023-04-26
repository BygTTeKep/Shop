package repository_test

import (
	"os"
	"testing"
)

var (
	databaseURL = "postgres://postgres:123456@localhost:5679/postgres?sslmode=disable"
)

func MainTest(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:12345@localhost:5679/postgres?sslmode=disable"
	}
	os.Exit(m.Run())
}
