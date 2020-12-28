package notestorage

import "gorm.io/gorm"

type storeMysql struct {
	db *gorm.DB
}

func NewMysqlStore(db *gorm.DB) *storeMysql {
	return &storeMysql{db: db}
}
