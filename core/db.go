package core

import (
	"fmt"
	"gotest/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:2iuixigi@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	// 自動建立資料表
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database connected.")
}
