package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"types:varchar(100);unique_index"`
	PasswordHashed string `gorm:"types:varchar(255);not null"`
}

func (User) TableName() string {
	return "user"
}

func Creat(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetUserByUserId(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	return &user, err
}
