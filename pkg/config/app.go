package app

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// global var db of type *gorm.DB
var (
	db *gorm.DB
)

// func to connect to database
func Connect() {
	d, err := gorm.Open("mysql", "shiva:yourpassword@/gocrudapisql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("Connected to database successfully")
	log.Println("Connected to database successfully")
}

// func to get db instance
func GetDB() *gorm.DB {
	return db
}
