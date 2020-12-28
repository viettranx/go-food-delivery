package repo

import (
	"context"
	"fooddlv/common"
	"fooddlv/food_likes/model"
)

type CreateFoodLikesStorage interface {
	CreateFoodLikes(ctx context.Context, data *model.FoodLikes) error
}
type createFoodLikesRepo struct {
	storage CreateFoodLikesStorage
}

func NewCreateFoodLikesRepo(storage CreateFoodLikesStorage) *createFoodLikesRepo {
	return &createFoodLikesRepo{
		storage: storage,
	}
}

func (repo *createFoodLikesRepo) CreateFoodLikes(ctx context.Context, data *model.FoodLikes) error {
	if err := repo.storage.CreateFoodLikes(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.FoodLikes{}.TableName(), err)
	}
	return nil
}
