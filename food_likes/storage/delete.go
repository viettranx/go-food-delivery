package storage

import (
	"context"
	"fooddlv/common"
	"fooddlv/food_likes/model"
)

func (s storeMysql) Delete(
	ctx context.Context,
	foodLike *model.FoodLikes) error {

	if err := s.db.Table(model.FoodLikes{}.
		TableName()).
		Delete(&foodLike).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
