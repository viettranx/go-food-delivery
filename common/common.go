package common

import "gorm.io/gorm"

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type Hasher interface {
	Hash() string
	GetSalt() string
}
