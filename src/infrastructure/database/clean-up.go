package database

import (
	"gorm.io/gorm"
)

func CleanUp(db *gorm.DB) {
	db.Exec("DELETE FROM hello_world_entities")
}
