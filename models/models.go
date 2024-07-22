package models

import (
	"go_api_mux_gorm/database"
	"time"
)

type Category struct {
	Id   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);" json:"name"`
}

// Estructure for define the type of the categories
// It will be the table name
type Categories []Category

type Product struct {
	Id           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(100);" json:"name"`
	Price        int       `gorm:"type:int;" json:"price"`
	Stock        int       `gorm:"type:int;" json:"stock"`
	Description  string    `gorm:"type:varchar(500);" json:"description"`
	DateRegister time.Time `gorm:"type:datetime;" json:"date_register"`
	CategoryID   uint      `gorm:"type:int;" json:"category_id"`
	Category     Category  `json:"category"`
}

type Products []Product

func Migrations() {
	database.Database.AutoMigrate(&Product{})
	//database.Database.AutoMigrate(&Category{})
}
