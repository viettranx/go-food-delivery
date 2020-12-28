package storage

import (
	"context"
	"fooddlv/common"
	"fooddlv/food_likes/model"
	"gorm.io/gorm"
)

func (store *storeMysql) FindByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	relations ...string,
) (*model.FoodLikes, error) {

	var foodLike model.FoodLikes
	db := store.db.Table(model.FoodLikes{}.TableName())

	if err := db.Where(conditions).First(&foodLike).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.ErrDB(err)
		}

		return nil, err
	}

	return &foodLike, nil
}
