package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ConnectDB connection database
func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	return db
}

// InitialMigration create database tables
func InitialMigration(model interface{}) {
	db := ConnectDB()
	defer db.Close()
	db.AutoMigrate(model)
}
