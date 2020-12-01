package detailstorage

import "gorm.io/gorm"

type orderDetailStorage struct {
	db *gorm.DB
}

func NewOrderDetailSQLStorage(db *gorm.DB) *orderDetailStorage {
	return &orderDetailStorage{db: db}
}
