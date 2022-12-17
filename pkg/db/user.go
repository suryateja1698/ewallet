package db

import (
	"log"

	"github.com/suryateja1698/ewallet/graph/model"
)

func AddUser(user model.User) error {
	tx := DB.Create(&user).Model(model.User{})
	if tx.Error != nil {
		log.Printf("err in creating user: %v", tx.Error)
	}
	return nil
}

func GetUser(id string) (*model.User, error) {
	var user model.User
	tx := DB.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		log.Printf("couldn't get user %v", tx.Error)
	}
	return &user, nil
}

func GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	tx := DB.Find(&users)
	if tx.Error != nil {
		log.Printf("couldn't get users %v", tx.Error)
	}
	return users, nil
}
