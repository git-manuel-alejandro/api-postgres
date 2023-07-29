package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS string = "host=localhost user=fazt password=mysecretpassword dbname=gorm port=5432"
var DB *gorm.DB

func DBConection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
