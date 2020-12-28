package usermodel

import (
	"fooddlv/common"
)

type CreateUser struct {
	common.SQLModel `json:",inline"`
	Email           string          `json:"email" gorm:"column:email;"`
	Password        string          `json:"password" gorm:"column:password;"`
	LastName        string          `json:"last_name" gorm:"column:last_name;"`
	FirstName       string          `json:"first_name" gorm:"column:first_name;"`
	Phone           *string         `json:"phone" gorm:"column:phone;"`
	Roles           common.RoleEnum `json:"role" gorm:"column:roles;type:ENUM('user','admin'); default:user"`
	Salt            string          `json:"salt" gorm:"column:salt;"`
	Avatar          *common.JSON    `json:"avatar" gorm:"column:avatar;type:json"`
}

func (CreateUser) TableName() string {
	return User{}.TableName()
}
