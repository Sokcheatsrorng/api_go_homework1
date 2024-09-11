package database

import (
	"homework1/internal/model"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func InitDB(dsn string) *gorm.DB {

    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Perform automatic migration
	err = db.AutoMigrate(&model.User{}, &model.Product{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    return db;
}