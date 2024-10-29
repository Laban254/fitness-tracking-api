package database

import (
    "fitness-tracking-api/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func InitializeDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("fitness.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    err = DB.AutoMigrate(
        &models.User{},
        &models.UserProfile{},
        &models.Goal{},       
        &models.Activity{},  
        &models.Workout{},
        &models.Nutrition{},
        &models.Progress{},
    )
    if err != nil {
        log.Fatalf("Error migrating database schema: %v", err)
    }
}
