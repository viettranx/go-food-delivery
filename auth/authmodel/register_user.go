package authmodel

import (
	"errors"
	"fooddlv/common"
	"fooddlv/user/usermodel"
)

const EntityName = "User"

type CreateUser struct {
	common.SQLModel `json:",inline"`
	Email           string          `json:"email" gorm:"column:email;"`
	Password        string          `json:"password" gorm:"column:password;"`
	LastName        string          `json:"last_name" gorm:"column:last_name;"`
	FirstName       string          `json:"first_name" gorm:"column:first_name;"`
	Phone           *string         `json:"phone" gorm:"column:phone;"`
	Roles           common.RoleEnum `json:"role" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Avatar          *common.JSON    `json:"avatar" gorm:"column:avatar;type:json"`
}

func (cu *CreateUser) Validate() error {
	if len(cu.Email) == 0 {
		return errors.New("email can't not empty")
	}
	if len(cu.Password) == 0 {
		return errors.New("password can't not empty")
	}
	return nil
}

func (CreateUser) TableName() string {
	return usermodel.User{}.TableName()
}

func (cu *CreateUser) ToCreateUser(hasher common.Hasher) *usermodel.CreateUser {
	return &usermodel.CreateUser{
		Email:     cu.Email,
		Password:  hasher.Hash(),
		LastName:  cu.LastName,
		FirstName: cu.FirstName,
		Phone:     cu.Phone,
		Roles:     cu.Roles,
		Salt:      hasher.GetSalt(),
		Avatar:    cu.Avatar,
	}
}
