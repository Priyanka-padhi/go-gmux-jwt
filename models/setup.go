package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: "host=localhost user=root password=123 dbname=app_db port=5432 sslmode=disable "}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error occured in database")
		//panic("Cannot connect to DB")
		//panic(err)
	}
	println("Running code after ListenAndServe (only happens when server shuts down)")
	DB.AutoMigrate(&User{})

}
