package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func Create(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}
func GetByID(ctx context.Context, db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("id =?", id).First(&user).Error
	return &user, err
}

func GetAll(ctx context.Context, db *gorm.DB) ([]*User, error) {
	var users []*User
	err := db.WithContext(ctx).Find(&users).Error
	return users, err
}

func DeleteUser(ctx context.Context, db *gorm.DB, id uint) error {
	//
	return db.WithContext(ctx).Delete(&User{}, id).Error
}

func AddUser(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}
