package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/suryateja1698/ewallet/graph/model"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%v dbname=%s user=%s password=%s sslmode=disable",
		"localhost",
		"5432",
		"gql-transaction",
		"gql",
		"teja",
	))
	if err != nil {
		fmt.Println("DB init fail: ", err)
		os.Exit(1)
	}

	err = db.Debug().AutoMigrate(model.User{}).Error
	if err != nil {
		fmt.Println("auto migration of user table failed")
		os.Exit(1)
	}
	DB = db
}
