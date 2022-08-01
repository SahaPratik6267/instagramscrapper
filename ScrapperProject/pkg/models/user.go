package models

import (
	"github.com/SahaPratik6267/instagramscrapper/ScrapperProject/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	var user User
	db := db.Where("email=?", email).Find(&user)
	return &user, db
}
