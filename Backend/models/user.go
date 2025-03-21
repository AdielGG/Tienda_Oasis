package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username" gorm:"unique"`
	Age      string `json:"age"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	//Image    []byte `json:"image"`
}
