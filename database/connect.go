package database

import (
	"fmt"
	"manualVuln/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("testDb.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}
