package database

import (
	"fmt"
	"log"

	"github.com/hs1121/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg model.Config) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",cfg.Username,cfg.Password,cfg.DbURL,cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
