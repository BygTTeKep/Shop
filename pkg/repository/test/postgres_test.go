package repository_test

import (
	"os"
	"testing"
)

var (
	databaseURL = "postgres://postgres:Schrlck_hlms1@localhost:5679/postgres?sslmode=disable"
)

func MainTest(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:Schrlck_hlms1@localhost:5679/postgres?sslmode=disable" //переписать в config
	}
	os.Exit(m.Run())
}
