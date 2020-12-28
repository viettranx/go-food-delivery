package storage

import (
	"context"
	"fooddlv/common"
	"fooddlv/food_likes/model"
)

func (store *storeMysql) CreateFoodLikes(ctx context.Context, data *model.FoodLikes) error {
	if err := store.db.Table(model.FoodLikes{}.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
