package models

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

import "gorm.io/driver/mysql"

var DB *gorm.DB
func ConnectToDatabase()  {
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?multiStatements=true&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db , err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil  {
		panic("cannot connect to database " + err.Error())
	}

	DB = db
}
