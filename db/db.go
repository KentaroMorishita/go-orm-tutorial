package db

import (
	"fmt"
	"strconv"

	"go-orm-tutorial/env"

	"github.com/jinzhu/gorm"
	// use mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectDB connection database
func ConnectDB() *gorm.DB {
	DBMS := env.Get("DBMS", "mysql")
	HOST := env.Get("HOST", "localhost")
	PORT := env.Get("PORT", "3306")
	DATABASE := env.Get("DATABASE", "go_orm_tutorial_db")
	USERNAME := env.Get("USERNAME", "")
	PASSWORD := env.Get("PASSWORD", "")
	DEBUG, _ := strconv.ParseBool(env.Get("DEBUG", "false"))

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", USERNAME, PASSWORD, HOST, PORT, DATABASE)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	if DEBUG {
		return db.Debug()
	}
	return db
}

// InitialMigration create database tables
func InitialMigration(model interface{}) {
	db := ConnectDB()
	defer db.Close()
	db.AutoMigrate(model)
}
