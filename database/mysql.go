package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
