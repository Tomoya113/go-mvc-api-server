package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(indb *gorm.DB) (err error) {
	if db != nil {
		db = indb
		return nil
	}

	const dsn = "root:password@tcp(127.0.0.1:3306)/go_mvc_api_server"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err
}

func Get() *gorm.DB {
	return db
}
