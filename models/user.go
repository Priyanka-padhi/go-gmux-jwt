package models

import "gorm.io/gorm"

type User struct {
	gorm.Model //it will make the struct into the model
	//Name       string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
