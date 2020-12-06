package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnViaGorm() (db *gorm.DB) {
	dbUser := "root"
	dbPass := "dqk09031998"
	dbName := "go_training"
	dsn := dbUser + ":" + dbPass + "@/" + dbName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return
}
