package imgstorage

import (
	"context"
	"fooddlv/common"
)

func (s *imgSqlStorage) GetImages(ctx context.Context, cond map[string]interface{}, ids []int) ([]common.Image, error) {

	var imgs []common.Image
	if err := s.db.Table(common.Image{}.TableName()).Find(&imgs).Where("id IN (?)", ids).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return imgs, nil
}
