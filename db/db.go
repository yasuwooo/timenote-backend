package db

import (
	"log"
	"timenote/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("timenote.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = DB.AutoMigrate(&model.ActivityLog{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
