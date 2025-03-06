package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	UserName string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//Image    []byte `json:"image"`
}
