package common

import (
	"fooddlv/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetDBConnection() *gorm.DB
	GetPubsub() pubsub.Pubsub
}

type Hasher interface {
	Hash() string
	GetSalt() string
}
