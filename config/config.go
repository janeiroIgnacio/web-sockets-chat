package config

import (
	"LefkasChat/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
func GetDBInstance() *gorm.DB {
	if db != nil {
		return db
	}
	db, _ := startDatabase();
	return db
}

//StartDatabase init db
func startDatabase() (*gorm.DB, error) {

	connectionString := fmt.Sprintf("root:root@tcp(localhost)/lefkas?charset=utf8&parseTime=true")
	log.Println("connection string: %s", connectionString)
	var err error
	DB, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Println("connection string: %s", connectionString)
		return nil, err
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
	return DB, nil
}
