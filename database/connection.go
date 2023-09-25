package database

import (
	"auth/database/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DBInstance *gorm.DB

func InitDB() {
	var err error
	port := os.Getenv("DB_PORT")
	url := os.Getenv("DB_URL")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := "host=" + url + " user=" + username + " password=" + password + " dbname=users port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"

	DBInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err, " Cannot connect to the DB!")
	}
	if DBInstance == nil {
		fmt.Println("INSTANCE IS FUCKING NULL")
	}
	fmt.Println("Connected to the DB")

	err = DBInstance.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal(err, " Cannot Migrate!")
	}
}
