package foodstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

func (s *storeMysql) List(
	ctx context.Context,
	paging *common.Paging,
	filter *foodmodel.ListFilter,
) ([]foodmodel.Food, error) {
	var rs []foodmodel.Food
	rs = make([] foodmodel.Food, 0)

	db := s.db.Table(foodmodel.Food{}.TableName()).Where("status = 1")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if paging.Total == 0 {
		return rs, nil
	}

	if err := s.db.
		Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).
		Find(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return rs, nil
}

