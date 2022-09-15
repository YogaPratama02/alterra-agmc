package config

import (
	"alterra-agmc/day-2/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	fmt.Println("port:", dbport)
	fmt.Println("username:", username)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbhost, dbport, dbname)

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Cannot connect database")
	}

	// sqlDB, _ := db.DB()
	// err = sqlDB.Ping()
	// if err != nil {
	// 	log.Fatalf("Database not responding with err: %s", err)
	// }

	MigrateDB()

	return db
}

func MigrateDB() {
	db.AutoMigrate(&models.User{})
}
