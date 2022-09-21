package config

import (
	"alterra-agmc/day-7/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Connect() *gorm.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbhost, dbport, dbname)

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Cannot connect database")
	}

	MigrateDB()

	return db
}

func MigrateDB() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
}
