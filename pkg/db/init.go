package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/suryateja1698/ewallet/pkg/models"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%v dbname=%s user=%s password=%s sslmode=disable",
		"host",
		"port",
		"name",
		"user",
		"password",
	))
	if err != nil {
		fmt.Println("DB init fail: ", err)
		os.Exit(1)
	}

	err = db.Debug().AutoMigrate(models.User{}).Error
	if err != nil {
		fmt.Println("auto migration of user table failed")
		os.Exit(1)
	}
	DB = db
}
