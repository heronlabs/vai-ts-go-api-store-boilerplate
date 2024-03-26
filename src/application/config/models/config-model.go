package config

import contexts "go-api-store-boilerplate/src/application/config/models/contexts"

type ConfigModel struct {
	Api      contexts.ApiConfig      `json:"apiConfig"`
	Database contexts.DatabaseConfig `json:"databaseConfig"`
}
