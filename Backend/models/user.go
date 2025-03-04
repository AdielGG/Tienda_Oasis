package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
