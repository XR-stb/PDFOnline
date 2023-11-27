package models

import "backend/pkg/user/role"

type User struct {
	Id       string    `json:"id" gorm:"column:id;type:char(36);primaryKey"`
	Username string    `json:"username" gorm:"column:username;type:varchar(32);unique;not null"`
	Password string    `json:"-" gorm:"column:password;type:char(32);not null"`
	Email    string    `json:"email" gorm:"column:email;type:varchar(64);not null"`
	Role     role.Role `json:"-" gorm:"column:role;type:tinyint(1);not null"`
}

func (User) TableName() string {
	return "users"
}
