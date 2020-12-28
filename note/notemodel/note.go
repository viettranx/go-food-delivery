package notemodel

import (
	"fooddlv/common"
)

const EntityName = "Note"

type Note struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title;"`
	Content         string `json:"content" gorm:"column:content;"`
	UserId          int    `json:"user_id" gorm:"column:user_id;"`
	//User            *SimpleUser `json:"user" gorm:"preload:false;"`
	// Computed fields
	HasLiked   bool `json:"has_liked" gorm:"-"`
	LikedCount int  `json:"liked_count" gorm:"-"`
}

func (Note) TableName() string {
	return "notes"
}

type SimpleUser struct {
	common.SQLModel `json:",inline"`
	LastName        string          `json:"last_name" gorm:"column:last_name;"`
	FirstName       string          `json:"first_name" gorm:"column:first_name;"`
	Roles           common.RoleEnum `json:"-" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Avatar          *common.JSON    `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}
