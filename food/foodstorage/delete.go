package foodstorage

import (
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

func (s *storeMysql) Delete(id int) error {
	err := s.db.Table(foodmodel.
	Food{}.
		TableName()).
		Where("id =?", id).
		UpdateColumns(map[string]interface{}{"status": 0}).Error

	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
