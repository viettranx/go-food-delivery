package imgstorage

import (
	"context"
	"fooddlv/common"
)

func (s *imgSqlStorage) Create(ctx context.Context, data []common.Image) error {
	db := s.db.Begin()

	if err := db.Table(common.Image{}.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
