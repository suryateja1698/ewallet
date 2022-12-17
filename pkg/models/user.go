package models

type User struct {
	ID       string `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
