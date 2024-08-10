package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	ID    uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name  *string `json:"name"`
	Title *string `json:"title"`
}

func MigrateEmployee(db *gorm.DB) error {
	err := db.AutoMigrate(&Employee{})
	return err
}
