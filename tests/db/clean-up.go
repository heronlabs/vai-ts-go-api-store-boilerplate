package tests

import (
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.Exec("DELETE FROM hello_world_entities")
}
