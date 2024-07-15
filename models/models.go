package models

import "go_api_mux_gorm/database"

type Category struct {
	Id   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);" json:"name"`
}

// Estructure for define the type of the categories
// It will be the table name
type Categories []Category

func Migrations() {
	database.Database.AutoMigrate(&Category{})
}
