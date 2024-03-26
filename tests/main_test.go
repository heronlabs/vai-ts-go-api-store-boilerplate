package tests

import (
	ConfigContexts "go-api-store-boilerplate/src/application/config/models/contexts"
	"go-api-store-boilerplate/src/infrastructure/database"
	"os"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	localConfig := ConfigContexts.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
		Password: "root",
		Name:     "goApiStoreDB",
	}

	db = database.Connect(localConfig)

	database.CleanUp(db)
}

func teardown() {
	database.CleanUp(db)
}
