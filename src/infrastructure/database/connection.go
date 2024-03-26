package database

import (
	"fmt"
	config "go-api-store-boilerplate/src/application/config/models/contexts"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbConfig config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("DATABASE OK!")

	return db
}
