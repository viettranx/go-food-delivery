package common

import "gorm.io/gorm"

const KeyCurrentUser = "CurrentUser"

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type Hasher interface {
	Hash() string
	GetSalt() string
}

type Requester interface {
	GetUserId() int
}
