package userstorage

import "gorm.io/gorm"

type userMySql struct {
	db *gorm.DB
}

func NewUserMysql(db *gorm.DB) *userMySql {
	return &userMySql{db: db}
}
