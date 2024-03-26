package main

import (
	"go-api-store-boilerplate/src/application/config"
	"go-api-store-boilerplate/src/core/entities"
	"go-api-store-boilerplate/src/infrastructure/database"
)

func main() {
	config := config.Config()

	db := database.Connect(config.Database)

	db.AutoMigrate(&entities.HelloWorldEntity{})
}
