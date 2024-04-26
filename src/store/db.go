package store

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func init() {
	host := os.Getenv("DB_HOST")
	var dial gorm.Dialector
	if host != "" {
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")

		port := os.Getenv("DB_PORT")
		dbname := os.Getenv("DB_NAME")
		dial = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname))
	} else {
		dial = sqlite.Open("eva.db")
	}

	_db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = _db
}

func GetDB() *gorm.DB {
	return db
}
