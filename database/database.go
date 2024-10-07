package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"os"
)


// Database is the connection handle
var Database = func () (db *gorm.DB) {
	// Load the .env file
	errorVar := godotenv.Load()
	if errorVar != nil {
		panic("Error loading .env file"+ errorVar.Error())
	}


	//Create domain service name for open the connection to the database
	dsn := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the connection to the database
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic("Failed to connect to database!")
	} else {
		println("Connected to database!")
		return db
	}

}()


