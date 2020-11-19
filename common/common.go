package common

import "gorm.io/gorm"

type AppContext interface {
	GetDBConnection() *gorm.DB
}
