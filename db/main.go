package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectDB connection database
func ConnectDB() *gorm.DB {
	DBMS := "mysql"
	HOST := "localhost"
	PORT := "3306"
	DATABASE := "go_orm_tutorial_db"
	USERNAME := "root"
	PASSWORD := "mysql"

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", USERNAME, PASSWORD, HOST, PORT, DATABASE)

	db, err := gorm.Open(DBMS, CONNECT)
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
