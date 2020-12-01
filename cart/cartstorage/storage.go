package cartstorage

import "gorm.io/gorm"

type cartMysql struct {
	db *gorm.DB
}

func NewCartMysql(db *gorm.DB) *cartMysql {
	return &cartMysql{db: db}
}
