package model

import (
	"github.com/info24/eva/store"
	"time"
)

func init() {
	db := store.GetDB()
	db.AutoMigrate(&Device{}, &User{})

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
