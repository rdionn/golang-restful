package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Todo struct {
	Name string
	Todo string
	gorm.Model
}

var sqliteDbCon *gorm.DB

func GetSqliteDbConnection() *gorm.DB {
	if sqliteDbCon != nil {
		return sqliteDbCon
	}

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Todo{})

	sqliteDbCon = db
	return db
}