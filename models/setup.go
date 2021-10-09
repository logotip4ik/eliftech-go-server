package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "postgres://postgres:aboba@localhost:5432/homework-1"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}

	db.AutoMigrate(&Post{})
	db.AutoMigrate(&User{})

	DB = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}