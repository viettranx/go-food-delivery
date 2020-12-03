package foodstorage

import (
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

func (s *storeMysql) Find(id int, moreInfos ...string) (*foodmodel.Food, error) {
	var food foodmodel.Food

	db := s.db.Table(foodmodel.
		Food{}.
		TableName()).Where("id = ?", id)

	for i := range moreInfos {
		db = db.Preload(moreInfos[i])
	}

	if err := db.First(&food).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &food, nil
}
