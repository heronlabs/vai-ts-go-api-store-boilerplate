package entities

import "gorm.io/gorm"

type BaseEntity struct {
	gorm.Model
	UUID string `json:"uuid"`
}
