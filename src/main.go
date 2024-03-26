package main

import (
	"go-api-store-boilerplate/src/application/api"
	"go-api-store-boilerplate/src/application/config"
	"go-api-store-boilerplate/src/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Config()

	db := database.Connect(config.Database)

	ginApp := gin.Default()
	api.Api(ginApp, db)
	api.Server(config.Api.Port, ginApp)
}
