package foodstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

func (s *storeMysql) Create(ctx context.Context, data *foodmodel.Food) error {
	if err := s.db.Table(foodmodel.Food{}.
		TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
