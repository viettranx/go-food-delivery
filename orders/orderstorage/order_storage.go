package orderstorage

import (
	"gorm.io/gorm"
)

type orderStorage struct {
	db *gorm.DB
}

func NewOrderSQLStore(db *gorm.DB) *orderStorage {
	return &orderStorage{db: db}
}
