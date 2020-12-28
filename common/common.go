package common

import (
	"fooddlv/pubsub"
	"gorm.io/gorm"
)

const KeyCurrentUser = "CurrentUser"

type AppContext interface {
	GetDBConnection() *gorm.DB
	GetPubsub() pubsub.Pubsub
	RealtimeEngine() RealtimeEngine
}

type Hasher interface {
	Hash() string
	GetSalt() string
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

type SimpleUser struct {
	SQLModel `json:",inline"`
	Email    string   `json:"email"`
	Roles    RoleEnum `json:"roles"`
}

func (u *SimpleUser) GetUserId() int {
	return u.ID
}
func (u *SimpleUser) GetEmail() string {
	return u.Email
}

func (u *SimpleUser) GetRole() RoleEnum {
	return u.Roles
}

type RealtimeEngine interface {
	EmitToUser(userId int, evt string, data interface{}) error
	BroadCastTo(userIds []int, evt string, data interface{}) error
	CheckOnline(userIds []int) map[int]bool
}
