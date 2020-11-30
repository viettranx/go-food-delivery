package foodstorage

import (
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

func (s *storeMysql) Find(id int) (*foodmodel.Food, error) {
	var food foodmodel.Food
	err := s.db.Table(foodmodel.
				Food{}.
				TableName()).First(&food, id).Error

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &food, nil
}
