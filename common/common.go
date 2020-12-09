package common

import (
	"fooddlv/pubsub"
	"gorm.io/gorm"
)

const KeyCurrentUser = "CurrentUser"

type AppContext interface {
	GetDBConnection() *gorm.DB
	GetPubsub() pubsub.Pubsub
}

type Hasher interface {
	Hash() string
	GetSalt() string
}

type Requester interface {
	GetUserId() int
}
