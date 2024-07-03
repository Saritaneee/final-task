package database

import (
	"log"
	"pbi-task/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbError error

func Connect(connectionString string) {
	DB, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect db")
	}
	log.Println("Connected to Database")
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Photo{})
	log.Println("Database migration completed")
}
