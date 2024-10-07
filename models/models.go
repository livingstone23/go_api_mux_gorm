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


type ProductPicture struct {
	Id 	   		uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name 		string `gorm:"type:varchar(100);" json:"name"`
	ProductID 	uint   `gorm:"type:int;" json:"product_id"`
	Product 	Product `json:"product"`
}

type ProductPictures []ProductPicture


//Create the structures for users
type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PerfilID uint   `gorm:"type:int;" json:"perfil_id"`
	Perfil   Perfil `json:"perfil"`
	Name 	string 	`gorm:"type:varchar(100);" json:"Name"`
	Email 	string 	`gorm:"type:varchar(100);" json:"email"`
	Phone 	string 	`gorm:"type:varchar(100);" json:"phone"`
	Password string `gorm:"type:varchar(100);" json:"password"`
	DateRegister time.Time `gorm:"type:datetime;" json:"date_register"`
}

type Users []User


type Perfil struct {
	Id   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);" json:"name"`
}

type Perfils []Perfil



//Method to create the tables in the database
func Migrations() {
	//database.Database.AutoMigrate(&ProductPicture{})
	//database.Database.AutoMigrate(&Product{})
	//database.Database.AutoMigrate(&Category{})

	database.Database.AutoMigrate(&User{}, &Perfil{} )
	//Una unica manera de habilitar el total de migraciones
	//database.Database.AutoMigrate(&ProductPicture{}, &Product{}, &Category{} )

}
