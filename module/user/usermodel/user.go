package usermodel

import (
	"fooddlv/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string          `json:"email" gorm:"column:email;"`
	Password        string          `json:"-" gorm:"column:password;"`
	LastName        string          `json:"last_name" gorm:"column:last_name;"`
	FirstName       string          `json:"first_name" gorm:"column:first_name;"`
	Phone           string          `json:"phone" gorm:"column:phone;"`
	Roles           common.RoleEnum `json:"-" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Salt            string          `json:"-" gorm:"column:salt;"`
	Avatar          *common.JSON    `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) ComparePassword(hasher common.Hasher) bool {
	hashedPassword := hasher.Hash()
	return u.Password == hashedPassword
}

func (u *User) IsActive() bool {
	if u == nil {
		return false
	}
	return u.Status == 1
}

func (u *User) ToSimpleUser() *common.SimpleUser {
	var simpleUser common.SimpleUser
	simpleUser.ID = u.ID
	simpleUser.Email = u.Email
	simpleUser.Roles = u.Roles
	return &simpleUser
}

