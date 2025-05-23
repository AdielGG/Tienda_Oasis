package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key,auto_increment"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username" gorm:"unique"`
	Age      string `json:"age"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
	//Image    []byte `json:"image"`
}
