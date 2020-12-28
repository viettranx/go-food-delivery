package imgstorage

import "gorm.io/gorm"

type imgSqlStorage struct {
	db *gorm.DB
}

func NewImgSqlStorage(db *gorm.DB) *imgSqlStorage {
	return &imgSqlStorage{
		db: db,
	}
}
