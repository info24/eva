package service

import (
	"github.com/info24/eva/model"
	"github.com/info24/eva/store"
)

func CreateUser(username, password string) error {
	user := model.User{Username: username, Password: password}
	tx := store.GetDB().Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
