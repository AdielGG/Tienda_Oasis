package models

type Suggestion struct {
	Id          int    `json: "id" gorm:"primary_key,auto_increment"`
	UserId      int    `json: "user_id" gorm:"not null" gorm:"foreignKey:UserId;references:ID"`
	Type        string `json: "type" gorm:"not null"`
	Title       string `json: "title" gorm:"not null"`
	Description string `json: "description" gorm:"not null"`
	Status      string `json: "status" gorm:"not null"`
}
