package config

import (
	"fmt"
	config "go-api-store-boilerplate/src/application/config/models"
	"os"
)

func Config() config.ConfigModel {
	var configModel config.ConfigModel

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		configModel.Api.Port = "1337"
	} else {
		configModel.Api.Port = apiPort
	}

	databaseHost := os.Getenv("DATABASE_HOST")
	configModel.Database.Host = databaseHost

	databasePort := os.Getenv("DATABASE_PORT")
	configModel.Database.Port = databasePort

	databaseUsername := os.Getenv("DATABASE_USERNAME")
	configModel.Database.Username = databaseUsername

	databasePassword := os.Getenv("DATABASE_PASSWORD")
	configModel.Database.Password = databasePassword

	databaseName := os.Getenv("DATABASE_NAME")
	configModel.Database.Name = databaseName

	fmt.Println("CONFIG OK!")

	return configModel
}
