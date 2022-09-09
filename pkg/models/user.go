package models

type User struct {
	ID       int    `gorm:"primary_key;auto_inc" json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
